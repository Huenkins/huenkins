package jobstack

import (
	"sync"
)

// IJob is interface for Job
type IJob interface {
	Dump() ([]byte, error)
	Restore([]byte) error
	Run() error
	Status() (*Status, error)
	Cancel() (*Status, error)
}

// Job is object for manager a job
type Job struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`
}

// NewJob is a simple constructor
func NewJob() *Job {
	return &Job{}
}

// Dump returns dumps for one job
func (job Job) Dump() ([]byte, error) {
	job.mu.Lock()
	defer job.mu.Unlock()

	return []byte{}, nil
}

// Restore recreates a job by dump
func (job Job) Restore([]byte) error {
	job.mu.Lock()
	defer job.mu.Unlock()

	return nil
}

// Run starts the job
func (job Job) Run() error {
	job.mu.Lock()
	defer job.mu.Unlock()

	return nil
}

// Status returns status for the job
func (job Job) Status() (*Status, error) {
	job.mu.Lock()
	defer job.mu.Unlock()

	return &Status{}, nil
}

// Cancel finishes the job
func (job Job) Cancel() (*Status, error) {
	job.mu.Lock()
	defer job.mu.Unlock()

	return &Status{}, nil
}
