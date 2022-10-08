package main

import "fmt"

func factorial(num int) int {
		if num <= 0 {
		return 1
	} else {
		return factorial(num-1) * num
	}
}

func main() {
	fmt.Print(factorial(8))
}
