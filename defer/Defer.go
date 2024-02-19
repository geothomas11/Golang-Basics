// package main

// import "fmt"

// func main() {
// 	a := 4
// 	b := 5
// 	mul := a * b
// 	add := mul + b
// 	avg := (a * b) / 2
// 	fmt.Println("Sum is", mul)
// 	fmt.Println("After Adding:", add)
// 	fmt.Println("average is:", avg)
// 	fmt.Println("After Defering")
// 	defer fmt.Println("Multiplied value  is: ", mul)

// }

// package main

// import "fmt"

// func main() {
// 	for i, cards := range cards {
// 		fmt.Println(i, cards)
// 		fmt.Println(cards)

// 	}
// }

package main

import "fmt"

func add(a, b int) int {
	sum := a + b
	return sum

}

func main() {
	fmt.Println(add(10, 5))
}
