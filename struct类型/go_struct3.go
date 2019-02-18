package main
import (
	"fmt"
)
type Human struct{
	name string
	age int
	phone string
}

type Employea struct{
	Human
	speciality string
	phone string
}
func main() {
	Bob := Employea{Human{"Bob",34,"123-456"},"Designer","456-789"}
	fmt.Println("Bob's work phone is :",Bob.phone)
	fmt.Println("Bob's personal phone is :",Bob.Human.phone)
}