package main
import (
	"fmt"
)

func sum(a,b int) (int,int) {
	return a+b,a*b
}
func main() {
	x := 3
	y := 4
	xplus,xtimes := sum(x,y)

	fmt.Printf("%d + %d = %d\n",x,y,xplus)
	fmt.Printf("%d * %d = %d",x,y,xtimes)
}