package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//name := person{"go", "lang"}
	//fmt.Println(name)

	var name person

	name.firstName = "go"
	name.lastName = "lang"

	fmt.Println(name)
	fmt.Printf("%+v", name)
}
