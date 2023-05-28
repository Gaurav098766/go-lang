package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

// defining struct
type person struct{
	firstName string
	lastName string
	contact contactInfo
}

func main()  {
	jim := person{
		firstName:"Jim",
		lastName: "Party",
		contact:contactInfo{
			email: "jim@gmail.com",
			zipCode:98526, 
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v",p)
}