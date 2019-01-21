package main

import (
	"fmt"
)

func main() {
	//full fledged for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	i := 0
	isLessThanFive := true
	//for loop with just a condition
	for isLessThanFive {
		if i == 10 {
			isLessThanFive = false
		}
		fmt.Println(i)
		i++
	}
	//for loop with nothing
	for {
		if i == 15 {
			break
		}
		fmt.Println(i)
		i++
	}
}
