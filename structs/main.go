package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	person := Person{"Vitor", 26}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}