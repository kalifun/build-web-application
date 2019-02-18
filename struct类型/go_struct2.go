package main
import (
	"fmt"
)
type Skills []string
type Human struct{
	name string
	age int
	weight int
}
type Student struct {
	Human
	Skills
	int
	speciality string
}
func main() {
	jane := Student{Human:Human{"jane",35,100},speciality:"biology"}
	fmt.Println("his name is ",jane.name)
	fmt.Println("his age is ",jane.age)
	fmt.Println("his weight is ",jane.weight)
	fmt.Println("his speciality is ",jane.speciality)
	jane.Skills = []string{"anatomy"}
	fmt.Println("his skills are",jane.Skills)
	fmt.Println("jane skills add")
	jane.Skills = append(jane.Skills,"physics","golang")
	fmt.Println("his all skills ",jane.Skills)
	jane.int = 3
	fmt.Println("his number is ",jane.int)
}