package task

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

/*

	How to use:

	print.Printf(
		`String is: {{.StringData}}, Int is: {{.IntData}}, Bool is: {{.BoolData}}`,
		print.Add("StringData", "stringData", "IntData", 9999, "BoolData", true)
	)

*/

const (
	TaskFuncPoint = "TaskFunc"
)

// ITask is interface for Task
type ITask interface {
	Run() error
	Status() (*Status, error)
	Cancel() (*Status, error)
}

// Task is object for manager all tasks
type OnePoint struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Data []byte `json:"data"`
}

// Task is object for manager all tasks
type AllPoint struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`

	Data []*OnePoint `json:"data"`
}

// Run starts the task
func (all *AllPoint) Get(name string) (*OnePoint, bool) {
	all.mu.Lock()
	defer all.mu.Unlock()

	for _, v := range all.Data {
		if v.Name == name {
			return v, true
		}
	}

	return nil, false
}

type TaskFunc func(context.Context, ITask)

// Task is object for manager all tasks
type Task struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`

	ID     string `json:"id"`
	Status string `json:"status"`

	InData string    `json:"in_data"`
	Data   *AllPoint `json:"data"`

	_cancel context.CancelFunc

	CreatedTime time.Time `json:"created_time"`
	StartTime   time.Time `json:"start_time"`
}

// NewTask is a simple constructor
func checkInData(data []*OnePoint) error {
	return nil
}

// NewTask is a simple constructor
func New(globalContext *GlobalContext, inData []byte) (*Task, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	data := AllPoint{}
	if err := json.Unmarshal(inData, &data); err != nil {
		return nil, err
	}

	if err := checkInData(data); err != nil {
		return nil, err
	}

	task := &Task{
		ID:     id.String(),
		Status: "created",

		InData: inData,
		Data:   data,

		// _cancel context.CancelFunc
		CreatedTime: time.Now(),
	}

	if err := task.run(globalContext); err != nil {
		return nil, err
	}

	return task, nil
}

// Run starts the task
func (task *Task) run(globalContext *GlobalContext) error {
	task.mu.Lock()
	defer task.mu.Unlock()

	runFuncName, err := task.Data.Get(TaskFuncPoint)
	if err != nil {
		return err
	}

	runFunc, err := globalContext.GetPluginFunction(runFuncName)
	if err != nil {
		return err
	}

	go func(ctx context.Context, task *Task, fn TaskFunc) {
		defer task.finish(lastID)
		fn(ctx, task)
	}(ctx, task, runFunc)

	return nil
}

// Status returns status for one task or for all tasks
func (stack *Task) Status() (*Status, error) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	return *Status{}, nil
}

// Cancel finishes one task
func (stack *Task) Cancel() (*Status, error) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	return &Status{}, nil
}
