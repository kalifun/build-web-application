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


