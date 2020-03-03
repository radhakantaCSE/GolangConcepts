package main

import "fmt"

func main() {

	/******************* ARRAY *************************/
	// Array declaration in Go
	var nums [5]int
	fmt.Println("Array : ", nums)

	// Array initialization
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array Declare:", b)

	// 2D array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	/******************* ARRAY *************************/

}
