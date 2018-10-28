package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	var err error

	if err != nil {
		fmt.Printf("Error! Crazy copy-paste leads to injury", err, debug.Stack())
	}
}