package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email  string
	mobile string
	addressInfo
}

type addressInfo struct {
	doorNo     string
	flatName   string
	streetName string
	city       string
	zip        int
}

func main() {

	//initializeStruct()
	workingWithSlice()
}

func workingWithSlice() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	UpdateSlice(mySlice)
	fmt.Println(mySlice)
}

func UpdateSlice(s []string) {
	s[0] = "Bye"
}

func initializeStruct() {
	alice := person{
		firstName: "Alice",
		lastName:  "Blob",
		contactInfo: contactInfo{
			email:  "alice@gmail.com",
			mobile: "1234567890",
			addressInfo: addressInfo{
				doorNo:     "401",
				flatName:   "Casa Grand",
				streetName: "ECR",
				city:       "Chennai",
				zip:        123456,
			},
		},
	}

	alice.updateName("Jimmy")
	alice.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (poineterToPerson *person) updateName(newFirstName string) {
	(*poineterToPerson).firstName = newFirstName
}
