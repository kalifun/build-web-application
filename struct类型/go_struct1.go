	package main
	import (
		"fmt"
	)
	type Human struct{
		name string
		age int
		weight int
	}
	type Student struct {
		Human
		speciality string
	}
	func main() {
		mark := Student{Human{"Mark",25,120},"computer"}
		fmt.Println("his name is ",mark.name)
		fmt.Println("his age is ",mark.age)
		fmt.Println("his weight is ",mark.weight)
		fmt.Println("his speciality is ",mark.speciality)
		mark.speciality = "AI"
		fmt.Println("Mark changed his speciality")
		fmt.Println("his new speciallity",mark.speciality)
		fmt.Println("mark become old")
		mark.age = 46
		fmt.Println("his age is ",mark.age)
		mark.weight += 60
		fmt.Println("his weight is ",mark.weight)
	}