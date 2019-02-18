# 面向对象
## method
> ### 现在假设有这个一个场景，你定义了一个struct叫做长方形，你现在想要计算它的面积。
```go
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
```
```
➜  面向对象 git:(master) ✗ go run go_area.go
Area of r1 is : 24
Area of r2 is : 36
```
#### 这段代码计算出长方形面积，但是area()不是作为oblong得方法实现的。
#### method的语法如下：
```go
func (r ReceiverType) funcname(parameters) (results)
```
#### 下面利用method实现的代码。
```go
package main

import (
	"fmt"
	"math"
)
type Rectangle struct{
	width,height float64
}
type Circle struct{
	radius float64
}
func (r Rectangle) area() float64{
	return r.width*r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi 
}
func main() {
	r1 := Rectangle{2,2}
	r2 := Rectangle{4,5}
	r3 := Circle{10}
	r4 := Circle{2}

	fmt.Println("Area of r1 is :",r1.area())
	fmt.Println("Area of r2 is :",r2.area())
	fmt.Println("Area of r3 is :",r3.area())
	fmt.Println("Area of r4 is :",r4.area())
}
```
```
➜  面向对象 git:(master) ✗ go run go_method.go 
Area of r1 is : 4
Area of r2 is : 20
Area of r3 is : 314.1592653589793
Area of r4 is : 12.5663706143591
```
> #### 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样。
> #### method里面可以访问接收者的字段
> #### 调用method通过.访问，就像struct里面访问字段一样。
#### 下面例子展示在任何的自定义类型中定义任意多的method
```go
package main

import (
	"fmt"
)
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)
type Color byte

type Box struct{
	width,height,depth float64
	color Color
}

type BoxList []Box

func (b Box) volume() float64 {
	return b.width*b.height*b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (b1 BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _,b := range b1 {
		if bv := b.volume();bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}
func (b1 BoxList) PaintBlack() {
	for i := range b1 {
		b1[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE","BLACK","BLUE","RED","YELLOW"}
	return strings[c]
}
func main() {
	boxes := BoxList {
		Box{4,4,4,RED},
		Box{10,10,1,YELLOW},
		Box{1,1,20,BLACK},
		Box{10,10,1,BLUE},
		Box{10,30,1,WHITE},
		Box{20,20,20,YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n",len(boxes))
	fmt.Println("The volume of the first one is ",boxes[0].volume(),"cm^3")
	fmt.Println("The color of the last one is ",boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is ",boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	fmt.Println("The color of the second one is ",boxes[1].color.String())

	fmt.Println("obviously,now, the biggest one is",boxes.BiggestColor().String())
	
}
```
- #### Color作为byte的别名。
- #### 定义了一个struct:Box,含有是哪个长宽高字段和一个颜色属性。
- #### 定义了一个slice:BoxList,含Box
#### 然后以上面的自定义类型为接收者定义了一些method
- #### volume()定义了接收者为Box，返回Box的容量。
- #### SetColor(c Color),把Box的颜色改为c
- #### BiggestColor()定义在BoxList上面，返回list里面容量最大的颜色
- #### PaintBlack()把boxlist里面所有的颜色设置为黑色
- #### String() 在Color上面，返回color的具体颜色。
## method继承
```go
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
```
```
➜  面向对象 git:(master) ✗ go run go_method3.go 
Hi,I am Mark you can call me on 222222222
Hi,I am Sam you can call me on 111111111
```
## method重写
> #### 如果Emplyee想要实现自己的sayhi怎么办？简单，和匿名字段冲突一样的道理，我们可以重写method来实现。
```go
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
func (e *Employee) SayHi() {
	fmt.Printf("Hi I am %s,I work at %s.Call me on %s\n",e.name,e.company,e.phone)
}
func main() {
	mark := Student{Human{"Mark",25,"222222222"},"MIT"}
	sam := Employee{Human{"Sam",45,"111111111"},"Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}
```
```
➜  面向对象 git:(master) ✗ go run go_method4.go
Hi,I am Mark you can call me on 222222222
Hi I am Sam,I work at Golang Inc.Call me on 111111111
```