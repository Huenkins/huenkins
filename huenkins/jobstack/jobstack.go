package jobstack

import (
	"sync"
)

/*

	How to use:

	print.Printf(
		`String is: {{.StringData}}, Int is: {{.IntData}}, Bool is: {{.BoolData}}`,
		print.Add("StringData", "stringData", "IntData", 9999, "BoolData", true)
	)

*/

// IJobStack is interface for JobStack
type IJobStack interface {
	Dump(string) ([]byte, error)
	DumpAll() ([][]byte, error)
	Restore([]byte) (string, error)
	Run(string) error
	Status(...string) ([]*Status, error)
	Cancel(string) (*Status, error)
	CancelAll() ([]*Status, error)
	IDs() []string
}

// JobStack is object for manager all jobs
type JobStack struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`
}

// New
func New() *JobStack {
	return &JobStack{}
}

// Dump returns dumps for one job
func (stack JobStack) Dump(string) ([]byte, error) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	return []byte{}, nil
}

// Restore recreates a job by it's dump
func (stack JobStack) Restore([]byte) (string, error) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	return "", nil
}

// Run starts one job executing
func (stack JobStack) Run(string) error {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	return nil
}

// Status returns status for one job or for all jobs
func (stack JobStack) Status(...string) ([]*Status, error) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	return []*Status{}, nil
}

// Cancel finishes one job
func (stack JobStack) Cancel(string) (*Status, error) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	return &Status{}, nil
}

// IDs return list of id for all jobs
func (stack JobStack) IDs() []string {
	stack.mu.RLock()
	defer stack.mu.RUnlock()

	return []string{}
}
