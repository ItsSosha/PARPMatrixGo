package main

import (
	"fmt"
	"sync"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArray(wg *sync.WaitGroup, array []int) {
	defer wg.Done()
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printArray(&wg, arr)
	}

	wg.Wait()
}
