package main

import "fmt"

func main() {

	// // Loooping through the array

	// var grades = [5]int{90, 80, 70, 60}

	// for i := 0; i < len(grades); i++ {
	// 	fmt.Println(grades[i])
	// }

	// //RANEGE KEYWORD USEgf
	// var grades = [5]int{90, 80, 70, 60}

	// for index, element := range grades {
	// 	fmt.Println(index, "=>", element)

	// }

	// 	//Multidimentional array

	// 	arr := [3][2]int{{2, 3}, {4, 16}, {8, 64}}
	// 	fmt.Println(arr[2][1])
	arr := [6]int{1, 2, 3, 4, 5, 6}

	for _, value := range arr {
		if value%2 == 0 {
			println("even num=", value)
		} else {
			fmt.Println("odd num=", value)
		}
	}
}
