package main
import (
	"fmt"
)
func main() {
	i := 10
	switch i {
	case 1 :
		fmt.Println("i is equal to 1")
	case 2 :
		fmt.Println("i is equal to 2")
	case 3,4,5 :
		fmt.Println("i is equal to 3,4,or 5")
	default:
		fmt.Println("All I know is that i is an integer.")
	}
}