# struct类型
## struct
> ### Go语言值中，也和c或者其他语言一样，我们可以声明新的类型，作为其他类型的属性或字段的容器。
```go
type person struct{
    name string
    age int
}
```
- #### 一个string类型的字段name，用来保存用户名称的属性。
- #### 一个int类型的字段age，用来保存用户年龄这个属性。
```go
type person struct{
    name string
    age int
}
var p person
P.name = "fun"
P.age = 23
fmt.Printf("The person's name is %s",P.name)
```
#### 除了上面这种P的声明使用之外，还有另外几种声明使用方法：
- ##### 按照顺序提初始值。P := person{"Tom",25}
- ##### 通过filed:value的方式初始化，这样可以任意顺序。P := person{age:24,name:"TIM"}
- ##### 当然也可以通过new函数分配一个指针，此处P的类型为*person P := new(person)
```go
package main
import (
	"fmt"
)
type person struct{
	name string
	age int
}
func older(p1,p2 person) (person,int) {
	if p1.age > p2.age {
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}

func main() {
	var tom person
	tom.name,tom.age = "TOM",18
	bob := person{age:25,name:"BOB"}
	pual := person{"PUAL",43}

	tb_older,tb_diff := older(tom,bob)
	tp_older,tp_diff := older(tom,pual)
	bp_older,bp_diff := older(bob,pual)


	fmt.Printf("of %s and %s,%s is older by %d years\n",tom.name,bob.name,tb_older.name,tb_diff)
	fmt.Printf("of %s and %s,%s is older by %d years\n",tom.name,pual.name,tp_older.name,tp_diff)
	fmt.Printf("of %s and %s,%s is older by %d years\n",bob.name,pual.name,bp_older.name,bp_diff)
}
```
```bash
➜  struct类型 git:(master) ✗ go run go_struct.go 
of TOM and BOB,BOB is older by 7 years
of TOM and PUAL,PUAL is older by 25 years
of BOB and PUAL,PUAL is older by 18 years
```
## struct的匿名字段
> ### 不写字段名的方式，也就是匿名字段，也称为嵌入式字段。当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。
```go
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
```
```bash
➜  struct类型 git:(master) ✗ go run go_struct1.go 
his name is  Mark
his age is  25
his weight is  120
his speciality is  computer
Mark changed his speciality
his new speciallity AI
mark become old
his age is  46
his weight is  180
```
#### 是不是觉得可以继承超酷，接下来是不是更酷。
```go
mark.Human = Human{"Marcus",55,220}
mark.Human.age -= 1
```
#### 通过匿名访问和修改字段相当有用，但是不仅仅是struct字段，所有的内置类型和自定义类型都可以作为匿名字段。
```go
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
```
```
➜  struct类型 git:(master) ✗ go run go_struct2.go 
his name is  jane
his age is  35
his weight is  100
his speciality is  biology
his skills are [anatomy]
jane skills add
his all skills  [anatomy physics golang]
his number is  3
```
> #### 如果human里面有一个字段叫phone，而student也有一个字段叫做phone，那么怎么办?
##### 最外层的优先访问，也就是当你通过student.phone访问的时候，是访问student里面的字段。
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

type Employea struct{
	Human
	speciality string
	phone string
}
func main() {
	Bob := Employea{Human{"Bob",34,"123-456"},"Designer","456-789"}
	fmt.Println("Bob's work phone is :",Bob.phone)
	fmt.Println("Bob's personal phone is :",Bob.Human.phone)
}
```
```
➜  struct类型 git:(master) ✗  go run go_struct3.go 
Bob's work phone is : 456-789
Bob's personal phone is : 123-456
```