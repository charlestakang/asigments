package main

import (
	"fmt"
	"strconv"
)

func main() {
	for n := 2000; n <= 3200; n++ {
		if (n%7 == 0) && (n%5 != 0) {
			fmt.Print(strconv.FormatInt(int64(n), 10) + ", ")
		}
	}
}
