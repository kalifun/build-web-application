# 流程和函数
## 流程控制
> ### Go中流程控制分三大类：条件判断，循环控制和无条件跳转
### if
```go
package main
import (
	"fmt"
)
func main() {
	x := 5
	if x > 10 {
		fmt.Println("x is greater than 10")
	}else{
		fmt.Println("x is less than 10")
	}
}
```
#### if有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方不起作用。
### goto
> #### Go有goto语句--请明智的使用它。用goto跳转到必须在当前函数内定义的标签。标签是大小写敏感的。
```go
package main
import(
	"fmt"
)
func main() {
	i := 0
Here:
	fmt.Println(i)
	i++
	goto Here
}
```
### for
> #### 既可以用来循环读取数据，而且当做while来控制逻辑，还能迭代操作。
```go
package main

import (
	"fmt"
)
func main() {
	sum := 0;
	for index:=0;index<10;index++{
		sum += index
	}
	fmt.Println("sum is equal to ",sum)
}
```
#### 在循环里面有两个关键操作break和continue。break操作是跳出当前循环，continue是跳过本次循环。当嵌套过深的时候，break可以配合标签使用，跳转至标签指定位置。
```go
package main

import (
	"fmt"
)
func main() {
	sum := 0;
	for index:=0;index<10;index++{
		if index == 5 {
			//break
			continue
		}
	}
	fmt.Println(index)
}
```
```
break 打印：10，9，8，7，6
continue打印出：10,9,8,7,6,4,3,2,1
```
#### for配合range可以用于读取slice和map。
```go
for k,v := range map {
	fmt.Println("map's key:",k)
	fmt.Println("map's value :",v)
}
```
> #### Go支持多返回值，对于“声明而未被调用”的变量，编译器会报错，在这种情况下，可以使用_来丢弃不需要的返回值。
### switch
> #### 有的时候需要些很多的if-else来实现一些逻辑，这时候代码看上去就会很丑很冗余，而且不易于后期维护，这时候switch就很好的解决这个问题。
```go
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
```
> #### case后面接的值类型必须一致。Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项。
## 函数
> ### 函数是Go里面的核心设计，它通过关键字func来声明。
```go
func funcname(input1 type,input2 type) (output1 type,output2 type){
	return value1,value2
}
```
- #### 关键字func用来声明一个函数funcname
- #### 函数可以有一个或者多个参数，每个参数后面带有类型，用过,分隔。
- #### 函数可以返回多个值
- #### 上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型。
- #### 如果只有一个返回值且不声明返回值变量，那么你可以省略包括返回值的括号。
- #### 如果没有返回值，那么就直接省略到最后的返回信息。
- #### 如果有返回值，那么必须在函数的外层添加return语句。
## 多个返回值
```go
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
```
> ### 最好命名返回值，因为不命名返回值，虽然是的代码更加简洁，但是会造成生成的文档可读性差。
## 变参
> #### Go函数支持变参。接受变参的函数是有着不定数量的变参。首先需要定义函数接受变参。
```go
func myfunc(arg ...int){}
```
#### arg ...int告诉这个函数接受不定数量的变参。在函数体重，变量arg是一个int的slice。
```go
for _,n := range arg {
	fmt.Printf("Add the number is : %d\n",n)
}
```
## 传参与传指针
> ###  当我们传一个参数值到被调用函数里时，实际上是传了这个值得copy，当在被调用的函数里改变参数值时，相应的实参是不会改变的，数值变化只作用在copy上。
```go
package main

import (
	"fmt"
)

func add(a int) int {
	a = a + 1
	return a
}
func main() {
	x := 3
	fmt.Println("x = ", x)
	x1 := add(x)
	fmt.Println("add(x) = ", x1)
	fmt.Println("x = ", x)
}
```
```
x =  3
add(x) =  4
x = 3
```
#### 当我们需要对实体进行改变怎么办，那就需要指针了。
```go
package main
import (
	"fmt"
)

func add1(a *int) int{
	*a = *a+1
	return *a
}
func main() {
	x := 3
	fmt.Println("x = ",x)
	x1 := add1(&x)
	fmt.Println("add1(&x) = ",x1)
	fmt.Println("x = ",x)
}
```
```
x =  3
add1(&x) =  4
x =  4
```
- #### 传指针是的很多个函数能操作同一个对象。
- #### 传指针比较轻量级，只是传内存地址，我们可以用指针传递体积大的结构体。
- #### Go语言中channel，slice，map这三种类型的实现机制类指针，所以可以直接传递，而不用取地址后传递指针。
## defer
> ### 延迟语句(defer)，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后函数返回。
```go
func main() bool {
	file.Open("file")
	defer file.Close()
	if failurex {
		return false
	}
	if failurey {
		return false
	}
	return true
}
```
#### 如果有很多调用defer，那么defer是采用后进先出的模式，下面输出4，3，2，1，0
```go
for i := 0;i < 5 ;i++ {
	defer fmt.Printf("%d",i)
}
```
## 函数作为值，类型
> ### 在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数。
```go
package main
import (
	"fmt"
)
type testInt func(int) bool
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int,f testInt) []int {
	var result []int
	for _,value := range slice {
		if f(value) {
			result = append(result,value)
		}
	}
	return result
}
func main() {
	slice := []int {1,2,3,4,5,7}
	fmt.Println("slice = ",slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd elements of slice are: ",odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ",even)
}
```
```
slice =  [1 2 3 4 5 7]
odd elements of slice are:  [1 3 5 7]
Even elements of slice are:  [2 4]
```
#### 函数当做值和类型在我们写一些通用接口时实用。testInt这个类型是宇哥函数类型，然后两个filter函数的参数和返回值与testInt类型是一样的，但是我们可以实现很多种的逻辑。
## Panic和Recover
> ### Go没像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制。
### Panic
> #### 是一个内建函数，可以中断原有的控制流程，进入一个panic状态中。当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。panic可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。
### Recover
> #### 是一个内建的函数，可以让进入panic状态的goroutine恢复过来。recover仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。如果当前的goroutine陷入panic状态，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
##### 使用panic
```go
var user = os.Getenv("USER")
func init() {
	if user == "" {
		panic("No Value for $USER")
	}
}
```
```go
func throwPanic(f func()) (b bool) {
	defer func(){
		if x := recover();x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果出现了panic，那么就可以恢复回来
	return  
}
```
