package task

import (
	"fmt"
)

type Task struct {
	at     int
	task   func() error
	active bool
}

func (t Task) String() string {
	return fmt.Sprintf("Task: %d", t.at)
}

func (t Task) At() int {
	return t.at
}

func (t *Task) Active() bool {
	return t.active
}

func (t Task) ShouldExecute(frameCount int) bool {
	return t.at <= frameCount
}

func (t Task) Execute() error {
	return t.task()
}

func NewTask(at int, task func() error) Task {
	return Task{
		at:     at,
		task:   task,
		active: true,
	}
}
