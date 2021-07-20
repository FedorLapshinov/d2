package main

import "fmt"

func main() {
	fmt.Println(SolutionSquareGenerator(67, 5))
}

func SolutionSquareGenerator(start, n int) []int {
	var i int
	var q int
	qwer := []int{}
	for i = 0; i < n; i++ {
		q = (start + i) * (start + i)
		qwer = append(qwer, q)
	}
	return qwer
}
