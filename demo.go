package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// types
	var x1 string = "Hello World from go."
	y1 := uint(math.Pow(4, 2))
	yy1 := fmt.Sprint((y1))
	yyy1, error1 := strconv.ParseInt(yy1, 10, 32)
	z1 := x1 + fmt.Sprint(y1)
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(yyy1, error1)
	fmt.Println(z1)

	// conditions
	x2 := uint(8)
	y2 := 10
	z2 := int(x2) < y2
	fmt.Println(z2)
	if x2 > 10 {
		fmt.Println("run")
	}

	// loops
}
