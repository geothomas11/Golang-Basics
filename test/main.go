// package main

// import "fmt"

// func add(x float64, y float64) float64 {
// 	sum := x + y
// 	return float64(sum)

// }
// func main() {

// 	fmt.Println(add(12.58, 7.56))

// }

//Func for add and find avg///

// package main

// import "fmt"

// func did(x, y int) int {
// 	sum := x + y
// 	return sum

// }
// func avg(a, b int) int {
// 	avg := (a + b) / 2 //make sure if the brackets are on on the arguments
// 	return avg

// }

// func main() {

// 	fmt.Println(did(5, 9))
// 	fmt.Println(avg(10, 8))
// }

// // SWAP the words//
// package main

// import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x

// }
// func main() {
// 	a, b := swap("Hello", "World")
// 	fmt.Println(a, b)

// }

// // Naked Return//
// package main

// import "fmt"

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	return

// }

// func main() {
// 	fmt.Println(split(30))

// }

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

}

// package main

// import "fmt"

//	func main() {
//		v := "ABS" //change me
//		fmt.Printf("V is the type %T\n", v)
//	}

//FINDING SMALL,Large,highest
// package main

// import "fmt"

// const (
// 	//CREATE A HUGE NUMBER BY SHIFTING A 1 BIT LEFT 100 places

// 	//IN OTHERWORDS THE BINARY NUMBER THATS IS 1 FOLLOWED BY ZEROS

// 	Big = 1 << 150
// 	//Shift it Right agian 99 places,so we end up with 1<<1, or 2.

// 	Small = Big >> 149
// )

// func needInt(x int) int {
// 	return x*10 + 1
// }
// func NeedFloat(x float64) float64 {
// 	return x * 0.1

// }

// func main() {
// 	fmt.Println(needInt(Small))
// 	fmt.Println(NeedFloat(Small))
// 	fmt.Println(NeedFloat(Big))

// }
