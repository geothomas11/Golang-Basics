// package main

// import "fmt"

// func main() {
// 	var stud map[string]string
// 	stud["en"] = ("English")

// 	fmt.Println(stud)

// }

package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[:6]
	d := a[3:]
	e := a[3:6]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(a)
}
