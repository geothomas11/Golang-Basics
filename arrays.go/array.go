package main

// import "fmt"

// func main() {
// 	var grades [5]int = [5]int{10, 20, 30, 1}

// 	var fruits [3]string = [3]string{"apple", "lemon", "banana"}
// 	names := [4]string{"anil", "ravi", "krishna", "vish"}

// 	students := [...]string{"Chacko", "ananya", "rahul", "arun", "dinil"}

// 	fmt.Println(len(grades))
// 	fmt.Println(len(fruits))
// 	fmt.Println(students)
// 	fmt.Println(names)

// // Loooping through the array

// var grades = [5]int{90, 80, 70, 60}

// for i := 0; i < len(grades); i++ {
// 	fmt.Println(grades[i])
// }

// // //RANEGE KEYWORD USE
// var grades = [5]int{90, 80, 70, 60}

// for index, element := range grades {
// 	fmt.Println(index, "=>", element)

// }

// //Multidimentional array

// arr := [3][2]int{{2, 3}, {4, 16}, {8, 64}}
// fmt.Println(arr[2][1])
// package main

import "fmt"

func filterOdd(input []int) []int {
	var result []int
	for _, value := range input {
		if value%2 != 0 {
			result = append(result, value)
		}
	}
	return result
}

func filterEven(input []int) []int {
	var result []int
	for _, value := range input {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	// First array with 5 values
	firstArray := []int{1, 2, 3, 4, 5}

	// Use different functions to filter odd and even numbers
	oddNumbers := filterOdd(firstArray)
	evenNumbers := filterEven(firstArray)

	// Print the resulting arrays
	fmt.Println("Resulting array with odd numbers:", oddNumbers)
	fmt.Println("Resulting array with even numbers:", evenNumbers)
}
