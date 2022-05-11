package main

import "fmt"

func main() {
	fmt.Println("Even Odd Assignment")
	rng := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, r := range rng {
		if r%2 == 0 {
			fmt.Println(r, " is even")
		} else {
			fmt.Println(r, " is odd")
		}
	}
}
