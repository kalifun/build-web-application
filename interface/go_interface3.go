package main

import (
	"fmt"
	"strconv"
)

type Human struct{
	name string 
	age int
	phone string
}

func (h Human) String() string{
	return "<" +h.name+" - "+ strconv.Itoa(h.age) + " years - "+h.phone+">"
}
func main() {
	Bob := Human{"Bod",39,"111111"}
	fmt.Println("This Human is :",Bob)
}