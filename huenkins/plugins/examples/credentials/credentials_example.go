package main

import (
	"fmt"

	"github.com/iostrovok/huenkins/huenkins/globalcontext"
)

var V *globalcontext.GlobalContext

func Init() {
	fmt.Printf("Hello, Init!!!!!\n")
}

func F() {
	fmt.Printf("Hello, number %T\n", V)
}

func T() int {
	return 10
}
