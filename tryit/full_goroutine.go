package tryit

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func FullGoroutine() {
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	for i:=range channel{
		fmt.Println(i)
	}
}

type Test1 struct {
	x int
}
func PtrReferGoroutine()  {
	//t := Test1{x: 1}
	a := 1

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(a)
	}()

	a = 3

	time.Sleep(5*time.Second)
}
