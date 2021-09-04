package tryit

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

var XXX string

type Task func(ctx context.Context) error

type TaskRecord struct {
	FuncName string
	StartAt  time.Time
	FinishAt time.Time
	Status   int
}

type TaskRecordClient struct {
	DB string
}

func (t *TaskRecordClient) print() {
	fmt.Println("this is a task record client")
}

func (t *TaskRecordClient) Insert(r *TaskRecord) error {
	fmt.Println(r)
	fmt.Println(t.DB)
	return nil
}

func Wrapper(func()) func() {
	return func() {
		fmt.Println("Wrapper", GetTaskName(task))
	}
}

func TimeElapsedWrapper(func()) func() {
	return func() {
		fmt.Println("TimeElapsedWrapper",GetTaskName(task))
	}
}

func Run() {
	XXX = "adfsgdhfgrewrstdhf"

	Wrapper(TimeElapsedWrapper(task1))()

	fmt.Println("end")
}

// GetTaskName 获取任务函数的名字
func GetTaskName(task Task) string {
	return runtime.FuncForPC(reflect.ValueOf(task).Pointer()).Name()
}

func task(ctx context.Context) error {
	fmt.Println("fuck ?")
	time.Sleep(2 * time.Second)
	return errors.New("fuck iur")
}
func task1() {
	fmt.Println("task1")
}
