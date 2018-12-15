package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	mobile  string
	address addressInfo
}

type addressInfo struct {
	doorNo     string
	flatName   string
	streetName string
	city       string
	zip        int
}

func main() {

	prsn := person{
		firstName: "Packiaseelan",
		lastName:  "Socretes",
		contact: contactInfo{
			email:  "packiaseelan14@gmail.com",
			mobile: "1234567890",
			address: addressInfo{
				doorNo:     "C4",
				flatName:   "Vinayaga flat",
				streetName: "Pandian Street",
				city:       "Chennai",
				zip:        123456,
			},
		},
	}

	fmt.Println(prsn)
	fmt.Printf("%+v", prsn)
}
