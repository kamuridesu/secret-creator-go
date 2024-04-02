package main

import (
	"fmt"
	"os"
	// "path/filepath"
	"runtime"
)

var (
	TempFileLocation string
)

func init() {
	path := "/tmp"
	if runtime.GOOS == "windows" {
		path = `\AppData\Local\Temp\`
	}
	file, err := os.CreateTemp(path, "")
	if err != nil {
		panic(err)
	}
	TempFileLocation = file.Name()
}

func main() {
	content := OpenTextEditor()
	fmt.Println(content)
	result := ConvertValuesToBase64(content)
	PrintResult(result)
}
