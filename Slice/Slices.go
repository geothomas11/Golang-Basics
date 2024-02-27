//package main

//import "fmt"

//func main() {

// //Slice
// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// slice_1 := arr[1:5]
// fmt.Println(arr)
// fmt.Println(slice_1)

// //Sub_slice
// arr := [9]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
// slice := arr[0:6]
// fmt.Println(slice)
// fmt.Println(arr)

// sub_slice := slice[0:5]
// fmt.Println(sub_slice)

// //MAKE FUNC
// slice := make([]int, 10, 20)
// fmt.Println(slice)
// fmt.Println(len(slice))
// fmt.Println(cap(slice))

// //TEST
// arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
// slice := arr[2:5]
// fmt.Println(cap(arr))
// fmt.Println(len(slice))

// //Test-2
// arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 1}
// slice := arr[:3]
// fmt.Println(arr)
// fmt.Println(slice)

// slice[1] = 900
// fmt.Println("After modification")
// fmt.Println(arr)
// fmt.Println(slice)

//FUNC APPEND

// arr := [6]int{10, 20, 30, 40, 50, 60}
// slice := arr[1:3]
// fmt.Println(slice)
// fmt.Println(len(slice))
// fmt.Println(cap(slice))

// //AFTER APPEND
// fmt.Println("After Append")
// slice = arr[1:3]
// slice_1 := append(slice, -50, 100)

// fmt.Println(slice_1)

//TEST_1 slices
// arr := [5]int{10, 20, 30, 40, 50}
// slice := arr[:2]
// arr_2 := [5]int{5, 10, 15, 20, 25}
// slice2 := arr_2[:2]
// newSlice := append(slice, slice2...) //append between slice1 & slice2
// fmt.Println(slice)
// fmt.Println(slice2)
// fmt.Println(newSlice)

//DELETING FROM Slice

// arr := [5]int{10, 20, 30, 40, 50}
// i := 2
// fmt.Println(arr)

// slice1 := arr[:i]
// slice2 := arr[i+1:]
// newSlice := append(slice1, slice2...)
// fmt.Println(newSlice)

// 	//COPYING FROM SLICE

// 	srcSlice := []int{10, 20, 30, 40, 50, 60}
// 	destSlice := make([]int, 5)
// 	num := copy(destSlice, srcSlice)
// 	fmt.Println(destSlice)
// 	fmt.Println("Number of elements copies:", num)

//}

// package main

// import "fmt"

// func main() {
// 	// Creating a slice of maps
// 	data := []map[string]interface{}{
// 		{"name": "John", "age": 25, "city": "New York"},
// 		{"name": "Alice", "age": 30, "city": "San Francisco"},
// 		{"name": "Bob", "age": 28, "city": "Los Angeles"},
// 	}

// 	// Accessing and printing elements from the slice of maps
// 	for _, item := range data {
// 		fmt.Printf("Name: %s, Age: %d, City: %s\n", item["name"], item["age"], item["city"])
// 	}
// }

// //PRINTING A MAP INSIDE AN ARRAY

// package main

// import "fmt"

// func main() {
// 	userR := map[string][]string{
// 		"alice":   {"admin", "editor"},
// 		"Bob":     {"Viewer"},
// 		"Charlie": {"editor"},
// 	}

// 	for username, values := range userR {
// 		fmt.Printf("User:%s,Roles:%v \n", username, values)
// 	}
// }
