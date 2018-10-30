package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
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
	if err = walkFiles("."); err != nil {
		return
	}

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

// FileMeta
type FileMeta struct {
	Name string `yaml:=filename`
}
func fileMeta(info os.FileInfo) (meta *FileMeta,err error) {
	meta = &FileMeta{
		Name: info.Name(),
	}
	return
}

func packFiles() (err error)  {
	buf := new(bytes.Buffer)
	zipwriter := zip.NewWriter(buf)

	// --
	var fileWriter io.Writer
	var filename string
	if fileWriter, err = zipwriter.Create(filename); err != nil{
		return
	}

	var fileReader io.Reader
	if _,err = io.Copy(fileWriter, fileReader); err != nil {
		return
	}
	// --

	if err = zipwriter.Close(); err != nil{
		return
	}

	// zipData = buf.Bytes()
	return
}