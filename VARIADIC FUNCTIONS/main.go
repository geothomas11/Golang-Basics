package main

import "fmt"

func funcName(arg1 int, arg2 string, arg ...float64) {
	fmt.Println(arg1)
	fmt.Println(arg2)
	fmt.Println(arg)

}

func main() {
	funcName(1, "Helo", 2.5, 3.7, 5.1)
	nums := []int{1, 2, 3, 4, 5}
	total := sum(nums...)
	fmt.Println(total)

}

func sum(values ...int) int {
	result := 0
	for _, v := range values {
		fmt.Println(v)
		result += v
	}
	return result

}
