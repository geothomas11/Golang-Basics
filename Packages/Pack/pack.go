package main

import (
	"fmt"
	mypackage "gose/Packages/myPackage"
)

func main() {
	number1 := 20
	number2 := 10

	sum := mypackage.MyFunction(number1, number2)
	fmt.Println(mypackage.U)

	fmt.Println("Sum is:", sum)

}
