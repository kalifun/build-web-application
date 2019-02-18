# interface
## 什么是interface
> ### 简单的说，interface是一组method签名的组合，我们通过interface来定义对象的一组行动。
## interface类型
>  ### interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就可以实现了此接口。
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
	loan float32
}
type Employee struct{
	Human
	company string
	money float32
}

func (h *Human) SayHi() {
	fmt.Printf("Hi ,I am %s you can call me on %s\n",h.name,h.phone)
}

func (h *Human) Sing(lyrics string) {
	fmt.Println("La la,la la la....",lyrics)
}

func (h *Human) Guzzle(beerstein string) {
	fmt.Println("Guzzle Guzzle Guzzle...",beerstein)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s,I work at %s.call me on %s\n",e.name,e.company,e.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerstein string)
}

type YongChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
func main() {

}
```
## interface值
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
	loan float32
}
type Employee struct{
	Human
	company string
	money float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi ,I am %s you can call me on %s\n",h.name,h.phone)
}
func (h Human) Sing(lyrics string) {
	fmt.Println("La la,la la la....",lyrics)
}
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s,I work at %s.call me on %s\n",e.name,e.company,e.phone)
}
type Men interface{
	SayHi()
	Sing(lyrics string)
}
func main() {
	mike := Student{Human{"Mike",25,"1111111"},"Mit",0.00}
	pual := Student{Human{"Pual",25,"11112222"},"Harvard",0.00}
	sam := Employee{Human{"Sam",30,"222222"},"Golang Inc.",1000}
	tom := Employee{Human{"Tom",30,"222222"},"Java Inc.",2000}

	var i Men
	i = mike
	fmt.Println("This is Mike,a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is tom,an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men,3)
	x[0],x[1],x[2] = pual,sam,mike

	for _,value := range x{
		value.SayHi()
	}
}
```
```
➜  interface git:(master) ✗ go run go_interface2.go
This is Mike,a Student:
Hi ,I am Mike you can call me on 1111111
La la,la la la.... November rain
This is tom,an Employee:
Hi, I am Tom,I work at Java Inc..call me on 222222
La la,la la la.... Born to be wild
Let's use a slice of Men and see what happens
Hi ,I am Pual you can call me on 11112222
Hi, I am Sam,I work at Golang Inc..call me on 222222
Hi ,I am Mike you can call me on 1111111
```
#### 通过上面的diamante，你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，Go通过interface实现了duck-typing。
## 空interface
> ### 空interface不包含任何的method，正因如此，所有的类型都实现了空interface。空interface对于描述起不打任何的作用，但是空interface在我们需要储存任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。
```go
var a interface{}
var i int = 5
s := "Hello world"
a = i 
a = s
```
## interface函数参数
> ### interface的变量可以持有任意实现该interface类型的对象，我们是不是可以通过定义interface参数，让函数接受各种类型的参数。
## interface 变量储存的类型
- ### Comma-ok断言
#### Go语言里面有一个语法，可以直接判断是否是该类型的变量：value,ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言类型。
```go
package main
import (
	"fmt"
	"strconv"
)

type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return "(name:" + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
func main() {
	list := make(List,3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis",70}

	for index,elemnet := range list {
		if value,ok := elemnet.(int);ok {
			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
		}else if value,ok := elemnet.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n",index,value)
		}else if value,ok := elemnet.(Person);ok{
			fmt.Printf("list[%d] is a Person and its value is %s\n",index,value)
		}else{
			fmt.Printf("list[%d] is of a different type\n",index)
		}
	}
}
```
```
➜  interface git:(master) ✗ go run go_interface4.go 
list[0] is an int and its value is 1
list[1] is a string and its value is Hello
list[2] is a Person and its value is (name:Dennis - age: 70 years)
```
- ### switch测试
```go
package main
import (
	"fmt"
	"strconv"
)

type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return "(name:" + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
func main() {
	list := make(List,3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis",70}

	for index,element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n",index,value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n",index,value)
		default:
			fmt.Printf("list[%d] is of a different type\n",index)
		}
	}
}
```
## 嵌入interface
```go
type Interface interface {
    sort.Interface
    Push(x interface{})
    Pop() interface{}
}
```
#### 我们看到sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了。下面三个方法：
```go
type Interface interface {
    Len() int
    Less(i,j int) bool
    Swap(i,j int)
}
```
## 反射
> ### 所谓反射就是能检查程序在运行的状态。我们一般用到的包是reflect包。
### 要去反射是一个类型的值(这些值都实现了空interface)，首先需要把它转化成reflect对象(reflect.Type或者reflect.Value,根据不同的情况调用不同的函数)
```go
t := reflect.TypeOf(i)
v := reflect.ValueOf(i)
```
### 转化为reflect对象后我们就可以进行一些操作了，也就是将reflect对象转化成相应的值。
```go
tag := t.Elem().Field(0).Tag
name := v.Elem().Field(0).String()
```
### 获取反射值返回相应的类型和数值
```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.println("Type:",v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
```