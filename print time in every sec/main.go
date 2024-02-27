package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)

		time.Sleep(1 * time.Second)
		if i >= 10 {
			break
		} else {
			if i != 10 {
				continue
			}
		}
	}

}
