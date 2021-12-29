package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	id          int
	firstName   string
	lastName    string
	contactInfo //or contact contactInfo
}

func main() {
	emre := person{id: 1, firstName: "Emre", lastName: "Bilal"} //person{1, "Emre", "Bilal"}
	fmt.Println(emre)                                           //output-> {1 Emre Bilal { 0}}
	fmt.Println(emre.firstName)                                 //output-> Emre
	fmt.Println(emre.lastName)                                  //output-> Bilal

	//zero value
	var emre2 person
	fmt.Println(emre2)       //output-> {0   { 0}}
	fmt.Printf("%+v", emre2) //output-> {id:0 firstName: lastName: contactInfo:{email: zipCode:0}}

	emre2.firstName = "Emre2"
	emre2.lastName = "Bilal"
	fmt.Printf("%+v", emre2) //output-> {id:0 firstName:Emre2 lastName:Bilal contactInfo:{email: zipCode:0}}

	//embedding structs
	alex := person{
		id:        2,
		firstName: "Alex",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "alex.party@gmail.com",
			zipCode: 94000,
		},
	}
	//fmt.Printf("%+v", alex)
	alex.print() //output-> {id:2 firstName:Alex lastName:Party contactInfo:{email:alex.party@gmail.com zipCode:94000}}

	//with receiver functions
	alex.updateName("Jim")
	fmt.Println(alex.firstName) //output-> Alex (It appears that it is not affected by the update as structs are value type)

	//Go is pass by value lang
	//Value Types: structs, int, float, string, bool (use pointers to change these things in a func)
	//Reference Types: slices, maps, channels, pointers, functions (don't worry about pointers with these)

	//pointers
	alexPointer := &alex //(&) gives the memory address of the value this variable is pointing at
	alexPointer.updateName2("Jim")
	fmt.Println(alex.firstName)

	alex.updateName2("Jimmy") //clean use
	fmt.Println(alex.firstName)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateName2(newFirstName string) { //(*) gives the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}
