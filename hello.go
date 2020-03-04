package main

import (
	"fmt"
	"strconv"
)

func sum(){
	var x int = 5
	var y int = 7
	var sum int = x+y
	fmt.Printf(strconv.Itoa(sum))
}

func main() {
	fmt.Printf("hello, world\n")
	sum()
}