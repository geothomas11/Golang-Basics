// Panic
package main

import "fmt"

func Div(a, b float64) float64 {
	var result float64
	defer func() { //using of defer function
		if r := recover(); r != nil { ///use recover to correct the error
			fmt.Println("recovered from a panic")
			result = 0
		}
	}() //anonymus function to execute it firstly

	result = (a / b)
	return result
}

func main() {

	func() {
		fmt.Println("The division is:")
	}()
	fmt.Println(Div(4, 5))
}
