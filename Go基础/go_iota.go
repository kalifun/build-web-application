package main

import (
	"fmt"
)
const (
	x = iota   //x = 0
	y = iota   //y = 1
	z = iota   //z =2
	w     //常量声明省略值时，默认和之前一个值得字面相同。因此w = 3
)

const v = iota    //每遇到一个const关键字，iota就会重置，v=0
const (
	h,i,j = iota,iota,iota   //h=0,i=0,j=0,iota在同一行值相同
)

const (
	a = iota //a = 0
	b = "B"   
	c = iota   //c = 2
	d,e,f = iota,iota,iota  //d=3,e=3,f=3
	g = iota //g=4
)

func main() {
	fmt.Printf("a = %d , b = %s , c = %d ,d = %d,e = %d,f = %d,g = %d,h = %d,i = %d,j = %d,x = %d,y = %d,z = %d,w = %d,v = %d",a,b,c,d,e,f,g,h,i,j,x,y,z,w,v)
}