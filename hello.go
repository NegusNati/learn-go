package main

import "fmt"


func main(){

	// hello()
	// value()


	
}

func hello(){
	fmt.Println("Hello world");
}

func value(){

	fmt.Println("Go" + " " + "String concatination")

	fmt.Print("no new line?")
	fmt.Print(" Yeah, this will be just next to the 'no new line' string!")
	fmt.Println("1+1 = " , 1+1)

	fmt.Println(7.0/3.0)

	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

