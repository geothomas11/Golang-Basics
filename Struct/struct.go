package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func mass() {
	person1 := Person{
		Name: "Kamal",
		Age:  20,
	}

	fmt.Println("1st name:", person1.Name)
	fmt.Println("Age", person1.Age)

	person2 := Person{
		Name: "Alice",
		Age:  18,
	}
	fmt.Println("2nd name :", person2.Name)
	fmt.Println("age:", person2.Age)

	fmt.Println(person1)
	fmt.Println(person2)
	// fmt.Println(mass(*person1.Age,16))

}
