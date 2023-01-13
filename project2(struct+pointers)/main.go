package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 515700,
		},
	}

	//varianta bruta
	//jimPointer := &jim //jim este o referinta la strct - &jim =>accesez adresa unde ii stocat
	//jimPointer.updateName("Jimmy")

	//varinata scurta
	jim.updateName("Jimmy")
	jim.print()

}

//*person - stelulta in fata unui tip inseamnca ca functia poate fi apelata doar cu un variabila de tip pointer
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
