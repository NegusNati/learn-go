package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	// hello()
	// value()
	// variable()
	// constant()
	// forloop()
	// branchingWithIf()
	// switches()
	// arrays()
	sliceExample()

}

func hello() {
	fmt.Println("Hello world")
}

func value() {

	fmt.Println("Go" + " " + "String concatination")

	fmt.Print("no new line?")
	fmt.Print(" Yeah, this will be just next to the 'no new line' string!")
	fmt.Println("1+1 = ", 1+1)

	fmt.Println(7.0 / 3.0)

	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variable() {

	var name string
	var num int
	// name = "Negus"
	// num = 1 by default it is set 0 surprisingly not null
	fmt.Println(name, num)
	var yes, no string = "yeah", "nope"
	fmt.Println(yes, no)
	var truth = true
	fmt.Println(truth)
	var y bool
	fmt.Println(y) // default is false

	app := "Apple" //for both initalization and declaring variables ' := '
	fmt.Println(app)

}

func constant() {

	const s string = "Negus"

	fmt.Println(s)
	const n = 5000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func forloop() {

	i := 1
	for i < 3 {
		fmt.Println(i) //1,2
		i += 1
	}

	for j := 1; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("======")

	for i := range 5 {
		fmt.Println("range", i)
	}

	fmt.Println("======")

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	for i, v := range arr {
		fmt.Println("range", i, v) // range in 'i' index with 'v' value
	}

	fmt.Println("======")
	for {
		fmt.Println(" just loop")
		break //until break just loop again and again
	}

	fmt.Println("======")
	for i := range 6 {
		if i%2 == 0 {
			fmt.Println("it is even")
			continue
		}
		fmt.Println(i)
	}

}

func branchingWithIf() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is not even")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisable by 4")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either or")
	}

}

func switches() {

	y := 2

	switch y {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("it is nither 1 or 2")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend , LFGGGG")
	default:
		fmt.Println("it's week day")

	}

	t := time.Now()

	switch {
	case t.Hour() >= 12:
		fmt.Println("after noon ")
	default:
		fmt.Println(" before noon", t.Hour())
	}
	fmt.Print(" Go is running on")

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Definetly not mine", os)
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type %T\n", t)

		}
	}

	whatAmI(true)
	whatAmI(30)
	whatAmI("string data")

}

func arrays() {

	var a [5]int
	fmt.Println("array a:", a) // [0 0 0 0 0] it is by default 0 if we don't initalize it

	a[4] = 100
	fmt.Println("array a after [4] is initalzed:", a) //
	fmt.Println("array a [4]:", a[4])                 //

	fmt.Println("array a length :", len(a)) //

	b := []int{1, 2, 3, 4, 5}
	fmt.Println("array b after initalization:", b) //

	// var c [] int
	// c ={ 1,2,3,4,5} we can't assign like this
	// fmt.Println("array c after initalization:",c)

	c := [...]int{1, 3, 6, 9} // if we don't want to assgin the array size
	fmt.Println("dcl :", c)

	d := [...]int{100, 3: 10, 11}
	fmt.Println("dcl :", d)

	var twoDimensional [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensional[i][j] = i + j
		}
	}

	fmt.Println("twoDimensional :", twoDimensional)

	twoD := [2][4]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	fmt.Println("twoD :", twoD)

}

func sliceExample() {

	// var s []string
	// fmt.Println("uninit: ", s, s == nil, len(s))

	// s = make([]string, 3)
	// fmt.Println("emp: ", s, s == nil, len(s), "Cap: ", cap(s))



	// An Array of integers
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v, Type: %T\n", arr, arr)
	//Array: [1 2 3 4 5], Type: [5]int

	// A Slice of integers
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice with values: %v, Type: %T\n", slice, slice)
	// Slice with values: [1 2 3 4 5], Type: []int

	//we can Create another slice with different length
	anotherSlice := []int{10, 20, 30}
	fmt.Printf("Another slice: %v, Type: %T\n", anotherSlice, anotherSlice)
	// Another slice: [10 20 30], Type: []int



}
