package main

import (
	"os"
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
	result := ConvertValuesToBase64(content)
	PrintResult(result)
}
