package main

import (
	"fmt"
)

func RunParallelTasks(tasks []func() error, maxTasksCount int, maxErrorsCount int) {
	chanel := make(chan func() error, maxTasksCount)
	errChanel := make(chan error, maxErrorsCount)

	for _, task := range tasks {
		chanel <- task
		go func(c chan func() error, errC chan error) {
			f := <-c
			err := f()
			errC <- err
		}(chanel, errChanel)
	}

	close(chanel)

	for err := range errChanel {
		fmt.Println(err)
	}
}

func main() {
	tasks := []func() error{
		func() error {
			fmt.Println("task 1")
			return nil
		},
		func() error {
			fmt.Println("task 2")
			return nil
		},
		func() error {
			fmt.Println("task 3")
			return nil
		},
		func() error {
			fmt.Println("task 4")
			return nil
		},
		func() error {
			fmt.Println("task 5")
			return nil
		},
		func() error {
			fmt.Println("task 6")
			return nil
		},
		func() error {
			fmt.Println("task 7")
			return nil
		},
	}

	RunParallelTasks(tasks, 2, 1)
}
