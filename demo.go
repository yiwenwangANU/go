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
	x3 :=0
	for x3 < 3 {
		fmt.Println(x3)
		x3++
	}
	str := "hello world你好"
	for idx:=0; idx < len(str); idx++{
		fmt.Printf("%c", str[idx])
	}
	for _, char :=range str{
		fmt.Printf("%c", char)
	}
	// array
	fmt.Println()
	fmt.Println("++++++++++++Array+++++++++++++")
	var arr1 = [2]int{1, 2} // cannot change size later on
	arr1[0] = 10
	fmt.Println(arr1)

	var arr2 = [...][2]int{{1, 1}, {2, 2}}
	arr2[0] = [2]int{10, 11}
	fmt.Printf("%T", arr2)

	for _, nested := range arr2{
		for _, value := range nested{
			fmt.Println(value)
		}
		
	}

	// slice
	fmt.Println("++++++++++++Slice+++++++++++++")
	arr3 := [5]int{1, 2, 3, 4, 5}
	sl := arr3[1:3] // pointer -> arr[1], length -> 2, capacity -> 4
	fmt.Println(sl)
	sl = sl[:4] // slice the slice
	// sl[0] = 22 // update slice update the origin array
	fmt.Println(sl)
	sl2 := []string{"hello", "world"}
	for i:=0; i<10;i++{
		sl2 = append(sl2, "tim") //how append in go work, create new underline array if reach capacity
		fmt.Println(sl2, len(sl2), cap(sl2))
	}
	sl3 := make([]int, 2, 20) // initialize with 0
	fmt.Println(sl3)
	updateSlice(sl3) // function muate slice
	fmt.Println(sl3)

	// struct
	fmt.Println("++++++++++++Struct+++++++++++++")
	p1 := Person{name:"foo", age: 1, favSport: []Sport{{"soccer", "D"}}}
	fmt.Println(p1)
	fmt.Println("name is "+ p1.getName())
	// function
	fmt.Println("++++++++++++function+++++++++++++")
	f1 := func(x int) int{
		return x+1
	}
	fmt.Println(callFunc(f1))
	// pointer
	fmt.Println("++++++++++++Pointer+++++++++++++")
	x4 :=0
	y4 := &x4
	fmt.Println(x4, y4, *y4)
	a4 := 5
	change(&a4)
	fmt.Println(a4)
}
type Sport struct {
	name string
	position string
}
type Person struct{
	name string
	age uint
	favSport []Sport
}

func (p Person) getName() string {
	return p.name
}

func updateSlice(sl []int) {
	sl[0] = -22
}

func change(x *int){
	*x = 10
}

func callFunc(callback func(int) int) int{
	return callback(10)
}