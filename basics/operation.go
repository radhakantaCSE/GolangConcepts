package main

import "fmt"

func main() {

	// ********************** For loop **********************
	for i := 1; i < 10; i++ {
		println("Data ", i)
	}

	for {
		println("Infinite loop")
		break
	}

	// While loop in go
	sum := 10
	for sum > 100 {
		sum = sum + 10
		println("Sum of value ", sum)
	}

	// **************** If/else condition *******************
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}

	if x := 10; x <= 10 {
		println("Hello - True")
	} else {
		println("Hello ... False")
	}

	// ********************** Switch cases **********************
	countryCode := 91
	switch countryCode {
	case 91:
		println("Country code - India")
	case 92:
		println("Country code - Canada")
	default:
		println("Invalid country code")
	}

	i := 9
	switch {
	case i < 10:
		println("Its a one digit number.")
	case i > 10 && i < 100:
		println("Two digit number")
	default:
		println("Unable to count")
	}

	// Switch case return a value from a function
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
