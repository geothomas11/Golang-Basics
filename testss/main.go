// package main

// import "fmt"

// func main() {
// 	originalMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

// 	// Define the keys you want to include in the slice
// 	selectedKeys := []string{"a", "c", "e"}

// 	// Create a new map to hold the selected key-value pairs
// 	slicedMap := make(map[string]int)

// 	// Copy the selected key-value pairs into the new map
// 	for _, key := range selectedKeys {
// 		if value, ok := originalMap[key]; ok {
// 			slicedMap[key] = value
// 		}
// 	}

// 	// Print the result
// 	fmt.Println(slicedMap)
// }

package main

import "fmt"

func getsum(a, b float64) float64 {
	return (a + b) / 2

}

func main() {
	num1 := 10
	num2 := 5
	result := getsum(float64(num1), float64(num2))
	fmt.Printf("Value is :%v\n", num1, num2, result)

}

func calculateAverage(a, b float64) float64 {
	return (a + b) / 2
}

// func main() {
// 	// Example usage
// 	num1 := 10.5
// 	num2 := 20.5

// 	result := calculateAverage(num1, num2)

// 	fmt.Printf("The average of %.2f\n", result)
// }
