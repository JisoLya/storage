package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type taskType uint8
type taskStatus uint8

const (
	DONE        taskStatus = 0
	UNDO        taskStatus = 1
	IN_PROGRESS taskStatus = 2
)

const (
	MAP    taskType = 0
	REDUCE taskType = 1
	PSEUDO taskType = 2
)

type Task struct {
	FileName   string
	TaskType   taskType
	TaskStatus taskStatus
}

func main() {
	file, err := ioutil.TempFile("", "temp")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	data := []byte("data to write")
	if _, err := file.Write(data); err != nil {
		panic(err)
	}

	readData, err := ioutil.ReadAll(file)
	fmt.Println(string(readData))
}
