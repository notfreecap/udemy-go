package main

import "fmt"

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range values {
		var nType string
		if val%2 == 0 {
			nType = "even"
		} else {
			nType = "odd "
		}
		fmt.Println(nType, val)
	}
}
