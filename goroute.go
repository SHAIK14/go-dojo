package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func task1() {
	fmt.Println("task1")
	wg.Done()
}
func task2() {
	fmt.Println("task2")
	wg.Done()
}

func goroute() {
	wg.Add(2)
	go task1()
	go task2()
	wg.Wait()
}
