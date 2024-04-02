package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func SaveFile(filename, content string) {
	if err := os.WriteFile(filename, []byte(content), os.ModeAppend); err != nil {
		panic(err)
	}
}

func ReadFile(filename string) *string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	stringdata := string(data)
	return &stringdata
}

func CreateTempFile() {
	if _, err := os.Stat(TempFileLocation); err == nil {
		SaveFile(TempFileLocation, "")
	} else if errors.Is(err, os.ErrNotExist) {
		os.WriteFile(TempFileLocation, []byte(""), 0644)
	} else {
		panic(err)
	}
}

func OpenTextEditor() *string {
	CreateTempFile()
	command := os.Getenv("EDITOR")
	if runtime.GOOS == "windows" {
		command = "notepad"
	} else if command == "" {
		command = "nano"
	}
	cmd := exec.Command(command, TempFileLocation)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Waiting for the editor to be closed...")
	cmd.Wait()
	content := ReadFile(TempFileLocation)
	os.Remove(TempFileLocation)
	return content
}
