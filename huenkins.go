package main

import (
	"fmt"
	"runtime"

	"github.com/iostrovok/huenkins/huenkins/jobstack"
	"github.com/iostrovok/huenkins/huenkins/print"
)

const (
	// One thread with one CPU, but at least 4.
	defulatTotalGoProcess int = 4
)

var (
	totalGoProcess int = 4
)

// One thread with one CPU.
func CheckNumCPU() int {
	if runtime.NumCPU() > defulatTotalGoProcess {
		return runtime.NumCPU()
	}
	return defulatTotalGoProcess
}

// One thread with one CPU.
func init() {
	totalGoProcess = CheckNumCPU()
}

// Main function
func main() {
	fmt.Printf("totalGoProcess: %d\n", totalGoProcess)

	print.Printf(`
	My String is: {{.StringData}}
	My Int is:    {{.IntData}}
	My Bool is :  {{.BoolData}}

`, print.Add("StringData", "stringData", "IntData", 9999, "BoolData", true))

	stack := jobstack.New()
	dump, err := stack.Dump("")
	print.Printf("dump: {{.dump}}, err: {{.err}}\n", print.Add("dump", dump, "err", err))
}
