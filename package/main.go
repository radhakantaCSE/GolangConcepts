package main

import (
	utils "GolangConcepts/utils"
	"fmt"
)

func main() {

	// If any variable will be start with "capital" char then, it will be exposed to use outside of package
	fmt.Println("Printing variable from variable.go -> ", utils.Expose)

	// Compile time error because, we can't use unExposed variable
	//fmt.Println("Printing variable from variable.go -> ", utils.unExposed)

	// Exposed function in utils tried to call
	fmt.Println("UnExposed function in util method -> ", utils.SendMessage())

	// UnExposed function in utils tried to call -----> Compile time error
	//fmt.Println("UnExposed function in util method -> ", utils.greetMessage())
}
