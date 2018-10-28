package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)

func main() {
	var mode string
	var err error
	flag.StringVar(&mode, "mode", "z", "Application mode")

	switch mode {
	case "z":
		err = makeSzip()
	}

	if err != nil {
		fmt.Printf("Error! Crazy copy-paste leads to injury", err, debug.Stack())
	}
}

func makeSzip() (err error)  {
	fmt.Println("It's makeZip ")
	return
}