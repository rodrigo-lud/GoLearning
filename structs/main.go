package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email  string
	zipCod int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerPerson *person) updateName(newName string) {
	(*pointerPerson).firstName = newName
}

func main() {
	// rod := person{"Rodrigo", "Gomes"}
	//rod := person{firstName: "Rodrigo", lastName: "Gomes"}
	//fmt.Println(rod)
	// var rod person
	// fmt.Printf("%+v\n", rod)

	// rod.firstName = "Rodrigo"
	// rod.lastName = "Gomes"
	// rod.contactInfo.email = "rodrigo.lud@gmail.com"

	// fmt.Println(rod)
	// fmt.Printf("%+v\n", rod)

	jQuest := person{
		firstName: "J",
		lastName:  "Quest",
		contactInfo: contactInfo{
			email:  "jquest@gmail.com",
			zipCod: 92071,
		},
	}
	jQuest.updateName("Jota")
	jQuest.print()

	// jQuestPonter := &jQuest
	// jQuestPonter.updateName("Jota")
	// jQuest.print()

}
