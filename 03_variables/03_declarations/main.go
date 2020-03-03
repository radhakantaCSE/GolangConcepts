package main

import (
	"fmt"
	"reflect"
)

var Exposer = "I am exposed myself to out side of package"

func main() {

	// Print hello world
	fmt.Println("Hello World!!!")

	// Concatenated value
	fmt.Println("Sum of two numbers(10,20) is : ", 10+30)

	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	// *********************** Variable declaration ***********************************
	var name = "Nagendra" // Without strict type (Infer type to variable)
	fmt.Println("My Name is ", name)

	var message string
	var a, b, c int // Multiple variable declaration
	a = 1
	message = "My Mom"
	fmt.Println(message, a, b, c)

	var msg = "Hello World!" // Multiple initialization
	var aa, bb, cc int = 1, 2, 3
	fmt.Println(msg, aa, bb, cc)

	var mixInfMsg = "Hello World!" // Mixed type infer to variables
	var mA, mB, mC = 1, false, 3
	fmt.Println(mixInfMsg, mA, mB, mC)

	salary, errMsg := 2000, "Error occur" // Most short hand (Only possible inside function)
	println("My salary is ", salary, errMsg)

	fmt.Println("Exposed to outside-> ", Exposer)

	var age int8 = 28 // With strict type
	fmt.Println("My Age is ", age)

	// *********************** Constants ***********************************
	const MESSAGE = "Invalid order id"
	println("Validation message -> ", MESSAGE)

	// *********************** Check type of variable ***********************************
	fmt.Println("Variable type is ", reflect.TypeOf(age))

}
