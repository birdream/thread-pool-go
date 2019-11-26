package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Task sss
type Task struct {
	Id uuid.UUID
}

// Exec gei me a func as task
func (t *Task) Exec(cb interface{}, args ...interface{}) interface{} {
	fmt.Printf("Task Id: %s \n", t.Id)

	if len(args) == 1 {
		cb.(func(interface{}))(args[0])
	} else if len(args) == 0 {
		cb.(func())()
	} else {
		cb.(func(...interface{}))(args...)
	}

	return nil
}

// NewTask get a task
func NewTask() *Task {
	uid := uuid.New()

	return &Task{Id: uid}
}
