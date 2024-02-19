package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("About struct")

	Geo := User{
		Name:   "Geo",
		Email:  "Gt@gmail.com",
		Status: true,
		Age:    16,
	}
	fmt.Println(Geo)
	fmt.Printf("Users details are: %+v\n", Geo)
	fmt.Printf("Users Name is :%v\n", Geo.Name)
	fmt.Printf("Users Email is:%v\n", Geo.Email)
	fmt.Printf("The adress is:%v\n", Geo.Age)
	fmt.Printf("The address of email is:%v", &Geo.Email)
}
