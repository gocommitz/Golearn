package main

import (
	"data"
	"fmt"
)

func main() {
	//imported package 'data'
	fmt.Println(data.PrintData("Joy"))
	//Pointers essential
	g := "Hello World"
	fmt.Println(g)
	b := &g
	fmt.Println(*b)
	*b = "End Of World"
	fmt.Println(g)
}
