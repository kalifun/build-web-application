# Go基础
## 定义变量
> ### Go语言里定义变量有多种方式。
### 使用var关键词是Go基础的定义变量方式。
```go
var variable type
//定义一个名称为variable，类型为type的变量
```
### 定义多个变量
```go
var name1,name2,name3 type
// 定义三个都为type的变量
```
### 定义变量并初始化值
```go
var variable type = value
// 初始化variable的变量为value值，类型为type
```
### 同时初始化多个变量
```go
var name1,name2,name3 type = v1,v2,v3
// 定义三个变量，他们分别初始化为相应的值name1为v1。。。然后go会依次初始化它们。
```
### 是否觉得还是稍微繁琐了点，go人性化帮我们解决了无需定义类型。
```go
var name1,name2,name3 = v1,v2,v3
// go会根据相应值得类型帮你初始化它们
```
### 看到上面你是不是觉得还是没差，依然繁琐。那就再来种更简约的。
```go
name1,name2,name3 := v1,v2,v3
// 编译器会根据初始化的值自动推导相应的类型
```
### := 符号取代了var和type，这种称为简短声明。不过它有一个限制，那就是它只能用在函数内部；函数外部无法编译通过，所以一般使用var方式来定义全局变量。
### _(下划线)是特殊的变量名，任何赋予它的值都会被丢弃；我们将35赋予b，并同时丢弃34.
```go
_,b := 34,35
```
## 常量
> ### 所谓常量，也就是在程序编译阶段就是确定下来的值，而程序在运行时无法改变该值。在go程序中，常量可定义为数值，布尔值，字符串等类型。
```go
const constname = value
// 如果需要，也可以明确指定常量类型
const Pi float32 = 3.1415926
```
### 下面是一些常量声明的例子：
```go
const Pi = 3.1415926
const i = 10000
const MaxThread = 10
const prefix = "xxfefce_"
```
### GO常量和一般程序语言不同的是，可以指定相当多的小数位数。
## 内置基础类型
### Boolean
> #### 在Go中，布尔值得类型为bool，值是true或false，默认为false。
```go
var isActive bool   //全局变量声明
var enable,disable = true,false  //忽略类型声明
func test() {
    var available bool    //一般声明
    valid := false      //简短声明
    available = true    // 赋值操作
}
```
### 数值类型
> #### 整数类型有无符号和带符号两种。go同时支持int和uint，但具体长度取决于不同编译器的实现。Go有已经定义好位数的类型：rune，int8，int16，int32，int64和byte，uint8，uint16，uint32，uint64.其中rune是int32的别称，byte是uint8的别称。
> #### 这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。
#### 浮点数的类型有float32和float64两种，默认是float64。
#### Go还支持复数，它默认类型是complex128（64位实数+64位虚数）。复数形式为RE+IMi，其中RE是实数部分，IM是虚数部分，而最后i是虚数单位。
```go
var c complex64 = 5+5i
fmt.Printf("Value is : %v",c)
```
### 字符串
> #### Go中的字符串都是采用UTF-8字符集编码。字符串是用一对双引号("")或者反引号('')括起来定义，它的类型是string。
```go
var frenchHello string
var emptyString string
func test() {
    no,yes,maybe := "no","yes","maybe"
    japanesHeelo := "konichiwa"
    frenchHello  = "Bonjour"
}
```
#### 在go字符串是不可变的，例如下面的代码编译会报错：cannot assign to s[0]
```go
var s string = "Hello"
s[0] = 'S'
```
#### 但是如果真的想要修改怎么办呢？
```go
s := "Hello"
c := []byte(s)  //将字符串s转换为[]byte类型
c[0] = 'C'
s2 := string(c)   //再转换回string类型
fmt.Printf("%s\n",s2)
```
#### Go中可以使用+操作连接两个字符串
```go
s := "Hello"
m := "world"
a := s+m
fmt.Printf("%s\n",a)
```
#### 修改字符串也可以
```go
s := "Hello"
s = "c" + s[1:]   //字符串虽然不能更改，但是可进行切片操作
fmt.Print("%s\n",s)
```
#### 如果要声明一个多行的字符串怎么办？可以通过‘来声明。
```go
m := 'hello
        world'
```
#### 括起来的字符为Raw字符串，没有字符转义，换换也将原样输出。
### 错误类型
> ### Go内置了一个error类型，专门用来处理错误信息，go的package里面还有专门有一个包errors来处理错误
```go
err := errors.New("emit macho dwarf:elf header corrupted")
if err != nil {
    fmt.Print(err)
}
```
### ~~Go数据底层的储存~~
## 分组声明
> ### 同时声明多个常量，变量，或者导入多个包时，可以采用分组的方式进行声明。
```go
import "fmt"
import "os"
const i = 100
const Pi = 3.1415
const prefix = "GO_"
var i int
var Pi float32
var prefix string
```
### 可以分组写成
```go
import (
    "fmt"
    "os"
)
const(
    i = 100
    Pi = 3.1415
    prefix = "Go_"
)
var(
    i int
    Pi float32
    prefix string
)
```
## iota枚举
> ### Go里面有一个关键词iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1
```go
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
```
## Go程序设计的一些规则
- ### 大写字母开头的变量是可以导出的，也就是其他包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
- ### 大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。
## array，slice，map
### array
> #### array就是数组
```go
var arr [n]type
```
#### 在[n]type中，n表示数组的长度，type表示储存元素的类型。对数组的操作和其他语言类似，都是通过[]来进行对于读取或者赋值。
```go
var arr [10]int
arr[0] = 42
arr[1] = 13
fmt.Printf("The first element is %d\n",arr[0])
fmt.Printf("The last element is %d\n",arr[10])
```
#### 由于长度也是数组类型的一部分，因此[3]int和[4]int是不同类型，数组也就不能改变长度。数组之间的赋值是值得赋值，当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，那么久需要用到后面介绍的slice类型。
#### 数组也可使用:=来声明
```go
a := [3]int{1,2,3}  
b := [10]int{1,2,3,4}
c := [...]int{4,5,6}   //长度省略而采用[...]的方式，go会自动根据元素个数来计算长度
```
#### Go支持嵌套数组，即多维数组。
```go
doublearray := [2][4]int{[4]int{1,2,3,4},[4]int{5,6,7,8}}
easyarray := [2][4]int{{1,2,3,4},{5,6,7,8}}
```
### slice
> #### 在众多场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”，在go里面这种数据结构叫slice。
#### slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
```go
var fslice []int
```
#### 接下来我们可以声明一个slice，并初始化数据。
```go
slice := []byte{'a','b','c','d'}
```
#### slice可以从一个数组或者一个已经创在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组开始位置，j是结束位置，但不包含array[j]，它的长度是j-1.
```go
var ar = [10]byte{'a','b','c','d','e','f','g','h','j'}
var a,b []byte   
a = ar[2:5]   //现在a含有的元素ar[2],ar[3],ar[4]
b = ar[3:5]   //b元素是ar[3],ar[4]
```
> #### 注意slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长大衣或者使用...自动计算长度，而声明slice时，方括号内没有任何字符。
#### slice需要注意的点
- ##### slice的默认开始位置0，ar[:n]等价于ar[0:n]
- ##### slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
- ##### 如果从一个数组里直接获取slice，可以这样ar[:],因为默认第一个序列是0.第二个是数组长度，即等价于ar[0:len(ar)]
> ##### slice是引用类型，当引用改变其中元素的值时，其他的所有引用都会改变该值。
- ###### 一个指针，指向数组中的slice指定的开始位置
- ###### 长度，即slice的长度
- ###### 最大长度，也就是slice开始位置到数组的最后位置的长度。
#####  slice内置函数
- ###### len获取slice的长度
- ###### cap获取slice的最大容量
- ###### append向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice。
- ###### copy 函数copy从源slice的src中赋值元素到目标dst，并返回复制的元素的个数。
### map
> #### map也就是Python中字典的概念，它的格式为map[keytype]valuetype
```go
var numbers map[string]int
numbers = make(map[string]int)
numbers["one"] = 1
numbers["ten"] = 10
numbers["three"] = 3
fmt.Println("第三个数字是：",numbers["three"])
```
- ##### map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取。
- ##### map的长度是不固定的，也就是和slice一样，也是一种引用类型
- ##### 内置的len函数同样适用于map，返回map拥有的key的数量。
- ##### map的值可以很方便的修改，通过number["one"] = 11可以很容易修改值。
- ##### map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须适用mutex lock机制。 
##### 通过delete删除map的元素
```go
rating := map[string]float32{"c":5,"go":4.5,"python":4.5,"c++":2}
csharpRating, ok := rating["c#"]
if ok {
    fmt.Println("c# is in the map",csharpRating)
}else{
    fmt.Pintln("we have no rating")
}
delete(rating,"c")
```
##### 上面说过了，map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应改变。
```go
m := make(map[string]sting)
m["hello] = "bobo"
m1 := m
m1["hello"] = "sla"
```
### make,new操作
> #### make用于内奸类型(map,slice和channel)的内存分配。new用于各种类型的内存分配。
#### 内建函数new本质上来说和其他语言的同名函数一样；new(T)分配了零值填充的T类型的内存空间，并且返回其他地址。
#### 内建函数make(T,args)与new(T)有着不同的功能，make只能创建slice，map和channel，并且返回一个有初始值的T类型。
### 零值
> #### 关于零值，所指并非是空值，而是一种”变量未填充前“的默认值，通常为0.