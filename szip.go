package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func walkFiles(path string) (err error){
	var files []os.FileInfo
	if files, err = ioutil.ReadDir(path); err != nil {
		return
	}
	for i := range files{
		full := filepath.Join(path, files[i].Name())

		if files[i].IsDir(){
			if err = walkFiles(full); err != nil{
				return
			}
		}
		// Metadata packing
	}
	return
}