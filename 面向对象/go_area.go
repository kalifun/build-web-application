package main
import (
	"fmt"
)
type oblong struct{
	width,height float64
}
func area(r oblong) float64{
	return r.width*r.height
}
func main() {
	r1 := oblong{12,2}
	r2 := oblong{9,4}
	fmt.Println("Area of r1 is :",area(r1))
	fmt.Println("Area of r2 is :",area(r2))
}