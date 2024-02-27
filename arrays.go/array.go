// package main

// import "fmt"

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
// import "fmt"

// func filterOdd(input []int) []int {
// 	var result []int
// 	for _, value := range input {
// 		if value%2 != 0 {
// 			result = append(result, value)
// 		}
// 	}
// 	return result
// }

// func filterEven(input []int) []int {
// 	var result []int
// 	for _, value := range input {
// 		if value%2 == 0 {
// 			result = append(result, value)
// 		}
// 	}
// 	return result
// }

// func main() {
// 	// First array with 5 values
// 	firstArray := []int{1, 2, 3, 4, 5}

// 	// Use different functions to filter odd and even numbers
// 	oddNumbers := filterOdd(firstArray)
// 	evenNumbers := filterEven(firstArray)

// 	// Print the resulting arrays
// 	fmt.Println("Resulting array with odd numbers:", oddNumbers)
// 	fmt.Println("Resulting array with even numbers:", evenNumbers)
// }

// //CREARTING AN ARRAY WITH THE EXISTING ARRAY

// package main

// import "fmt"

// func main() {
// 	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	arr2 := arr1

// 	fmt.Println("before Arr1 elements:", arr1)
// 	fmt.Println("Before Arr2 elements:", arr2)

// 	arr2[0] = 9

// 	fmt.Println("After Arr1 elements:", arr1)
// 	fmt.Println("After Arr2 elements:", arr2)
// }

// //Find the largest element from an array

// package main

// import "fmt"

// func main() {
// 	var large int = 0
// 	var arr [5]int

// 	arr = [...]int{64, 34, 13, 43, 73}

// 	large = arr[0]
// 	for i := 1; i <= 4; i++ {
// 		if large < arr[i] {
// 			large = arr[i]
// 		}
// 	}
// 	fmt.Println("largest element is:", large)

// }

// // FIND SECOND LARGEST ELEMENT IN AN ARRAY
// package main

// import "fmt"

// func main() {
// 	var large int = 0
// 	var arr []int
// 	var large2 int = 0

// 	arr = []int{64, 34, 13, 43, 73}

// 	large = arr[0]
// 	for i := 1; i <= 4; i++ {
// 		if large < arr[i] {
// 			large = arr[i]
// 		} else if large2 < arr[i] {
// 			large2 = arr[i]
// 		}
// 	}
// 	fmt.Println("largest element is:", large)
// 	fmt.Println("Second largest element is:", large2)

// }

// PRINT EVEN AND ODD NUMBS
// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// for i := 0; i <= 5; i++ {
// 	if arr[i]%2 == 0 {
// 		fmt.Println("EVen numbers are:", arr[i])
// 	} else {
// 		fmt.Println("Odd numbers are:", arr[i])
// 	}
// }

// ///use Range keyword
// for _, value := range arr {
// 	if value%2 == 0 {
// 		fmt.Println("even numbers are:", value)
// 	}

// }
//}

// package main

// import "fmt"

// func main() {
// 	data := []map[string]interface{}{
// 		"arun", "anu", "manu", "sanu",
// 	}

// 	for _, item := range data {
// 		fmt.Printf("String is%s\n", item)
// 	}
// }

