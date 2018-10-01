package jobstack

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/iostrovok/huenkins/huenkins/tasks"
	uuid "github.com/satori/go.uuid"
)

const (
	StatusInProgress = "in_progress"
	StatusWating     = "wating"
)

var ErrorInProgress error

func init() {
	ErrorInProgress = errors.New("in_progress")
}

// IJob is interface for Job
type IJob interface {
	Dump() ([]byte, error)
	Restore([]byte) error
	Run() error
	Status() (*Status, error)
	Cancel() (*Status, error)
	Data() (map[string]interface{}, error)
}

type JobFunc func(context.Context, IJob, string)

// Job is object for manager a job
type Job struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`
	in sync.RWMutex `json:"-"`

	JobID       string `json:"job_id"`
	JobStatus   string `json:"job_status"`
	HistoryType string `json:"history_type"`
	JobType     string `json:"job_type"`

	_cancel context.CancelFunc

	CreatedTime   time.Time      `json:"created_time"`
	LastStartTime time.Time      `json:"last_start_time"`
	NextStartTime time.Time      `json:"next_start_time"`
	Interval      time.Duration  `json:"interval"`
	History       []HistoryPoint `json:"history"`
	Tasks         []*tasks.Task  `json:"tasks"`

	runFunc JobFunc `json:"-"`

	// Update
	JobData map[string]interface{} `json:"job_data"`
}

// Data returns copy of Data
func (job *Job) Data() (map[string]interface{}, error) {
	job.mu.RLock()
	defer job.mu.RUnlock()

	data, err := json.Marshal(job.JobData)
	if err != nil {
		return nil, err
	}

	out := map[string]interface{}{}
	err = json.Unmarshal(data, &out)

	return out, err
}

// NewJob is a simple constructor
func NewJob() (*Job, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &Job{
		JobID:   id.String(),
		History: []HistoryPoint{},

		// Update
		JobData: map[string]interface{}{},
	}, nil
}

// Dump returns dumps for one job
func (job *Job) Dump() ([]byte, error) {
	job.mu.Lock()
	defer job.mu.Unlock()

	return []byte{}, nil
}

// Restore recreates a job by dump
func (job *Job) Restore([]byte) error {
	job.mu.Lock()
	defer job.mu.Unlock()

	return nil
}

// checkMustRun starts the job
func (job *Job) checkMustRun() bool {
	job.mu.Lock()
	defer job.mu.Unlock()

	if job.NextStartTime.After(time.Now()) || job.NextStartTime.Equal(time.Now()) {
		return true
	}

	return false
}

// finish updates the history
func (job *Job) finish(lastID string) {

}

// Run starts the job
func (job *Job) RunByTimer() error {
	job.in.Lock()
	defer job.in.Unlock()

	if job.JobType != "interval" {
		return nil
	}

	if job.NextStartTime.After(time.Now()) {
		return nil
	}

	job.LastStartTime = time.Now()
	job.NextStartTime = job.LastStartTime.Add(job.Interval)

	return job.AddTask()
}

// Run starts the job
func (job *Job) AddTask() error {
	job.mu.Lock()
	defer job.mu.Unlock()

	task, err := NewTask()
	if err != nil {
		return err
	}
	job.Tasks = append(job.Tasks, task)

	return nil
}

// Status returns status for the job
func (job *Job) Status() (*Status, error) {
	job.mu.Lock()
	defer job.mu.Unlock()

	return &Status{}, nil
}

// Cancel finishes the job
func (job *Job) Cancel() ([]*Status, error) {
	job.mu.Lock()
	defer job.mu.Unlock()

	out := []*Status{}

	for _, t := range job.Tasks {
		st, err := t.Cancel()
		out = append(out, err)
	}

	return out, nil
}
