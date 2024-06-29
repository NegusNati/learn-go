package main

import (
	"fmt"
	"math"
)

func main() {

	// hello()
	// value()
	// variable()
	// constant()
	forloop()
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
	for i, v := range arr {
		fmt.Println("range", i, v) // range in 'i' index with 'v' value
	}

	fmt.Println("======")
	for {
		fmt.Println(" just loop")
		break //until break just loop again and again
	}

	fmt.Println("======")
	for i := range 5 {
		if i%2 == 0 {
			fmt.Println("it is even")
			continue
		}
		fmt.Println(i)
	}

}
