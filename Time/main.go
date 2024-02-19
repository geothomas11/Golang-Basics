package main

import (
	"fmt"
	"time"
)

func main() {
	FormattedTime := time.Now().Format("2006-05-01 03:04 AM")
	fmt.Println("Formatted Time:", FormattedTime)

	parsedTime, _ := time.Parse("2006-05-06 11:04:00", "2024-01-31 13:04:00")
	fmt.Println("Parsed Time:", parsedTime)
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// startTime := time.Now()
// t := startTime.Format("03:04:05 PM")
// fmt.Println("Start Time:", t)
// //Perform Some Work or sleep
// time.Sleep(2 * time.Second)

// endTime := time.Now()
// t = endTime.Format("03:04:06 PM")
// fmt.Println("End time:", t)

// duration := endTime.Sub(startTime)
// fmt.Println("Duration:", duration)

// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	for {
// 		fmt.Println("Hi")
// 		time.Sleep(2 * time.Second)
// 		endTime := time.Now()
// 		fmt.Println(endTime)
// 		fmt.Println(time.Nanosecond)
// 		time.Date("2001-05-04 10:30:00")
// 	}

// currentTime := time.Now()
// fmt.Println("Current Time:", currentTime)

// //Formatted time
// FormattedTime := currentTime.Format("2006-01-02 01:05:25")
// fmt.Println("Formatted Time:", FormattedTime)

// //Parsing Time
// //Can parse a time string into a time.Time value using the 'time.Parse' fuunction:
// timeStr := "2023-09-21 11:35:45"
// parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
// if err != nil {
// 	fmt.Println("Error Parse Time:", err)
// } else {
// 	fmt.Println("Parsed Time:", parsedTime)
// }
// //Adding 2 Hours tothe current time
// newTime := currentTime.Add(2 * time.Hour)
// fmt.Println("New time is :", newTime)

// //Comparing two time values
// if currentTime.Before(newTime) {
// 	fmt.Println("Current time before new time")

// 	//Time Duration
// 	delay := 5 * time.Second
// 	fmt.Println("Wating for", delay)
// 	time.Sleep(delay) //Sleep is the specified Duration
// 	fmt.Println("Done wating!")
// 	fmt.Println("Current time and Date is :", FormattedTime)
//}
//}
