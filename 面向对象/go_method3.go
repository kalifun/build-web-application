package main
import (
	"fmt"
)
type Human struct{
	name string
	age int
	phone string
}
type Student struct{
	Human
	school string
}
type Employee struct{
	Human
	company string
}
func (h *Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}
func main() {
	mark := Student{Human{"Mark",25,"222222222"},"MIT"}
	sam := Employee{Human{"Sam",45,"111111111"},"Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}