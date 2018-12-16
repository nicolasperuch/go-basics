package main

import "fmt"

type address struct {
	adress string
	number string
	city string
}

type person struct {
	firstName string
	lastName string
	address
}

func main(){
	
	address := address{"Baker Street", "221B", "London"}
	person := person { 
		firstName: "Xeroque", 
		lastName: "Homes", 
		address: address,
	}
	
	address.updateAddressNumber("221")
	person.print()
	address.print()

	person.address.number = "221A"
	person.print()
	address.print()
	fmt.Println()
}

func (p person) print() {
	fmt.Printf("\n %v", p)
}

func (a address) print() {
	fmt.Printf("\n %v", a)
}

func (a *address) updateAddressNumber(number string) {
	(*a).number = number
} 
