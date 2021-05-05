package main

import (
	"fmt"

	something "github.com/learngo/Something"
)

func main() {
	fmt.Println("Hello world")
	something.SayHello()
	//Declaring Constant
	const name string = "jaedo1"
	//Declaring Variable
	var name1 string = "jaedo2"
	//Assigning value to variable
	name1 = "jambo2"
	//Print Variable
	fmt.Println(name)
	fmt.Println(name1)
	// This short version can only work inside the func main() {}
	lee := "jaedo3"
	fmt.Println(lee)
	//Go will determine the type of the variable
	lee1 := false
	fmt.Println(lee1)
}
