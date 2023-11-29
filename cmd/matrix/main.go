package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func initializeMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(10) // Random integers between 0 and 9
		}
	}
	return matrix
}

func multiplyMatricesSeq(a, b [][]int) [][]int {
	result := make([][]int, len(a))
	for i := range result {
		result[i] = make([]int, len(b[0]))
		for j := range b[0] {
			for k := range a[0] {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func multiplyMatricesParallel(a, b [][]int) [][]int {

	result := make([][]int, len(a))
	var wg sync.WaitGroup

	for i := range result {
		result[i] = make([]int, len(b[0]))
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := range b[0] {
				for k := range a[0] {
					result[i][j] += a[i][k] * b[k][j]
				}
			}
		}(i)
	}

	wg.Wait()
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rows, cols := 1024, 1024

	a := initializeMatrix(rows, cols)
	b := initializeMatrix(cols, rows)

	start := time.Now()
	multiplyMatricesSeq(a, b)
	fmt.Println("Sequential duration:", time.Since(start))

	start = time.Now()
	multiplyMatricesParallel(a, b)
	fmt.Println("Parallel duration:", time.Since(start))
}
