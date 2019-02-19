# 并发
## goroutine
> ### goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine只需要极少的栈内存(大概4~5KB)，当然会根据相应的数据伸缩。
### goroutine是通过runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。
```go
go hello(a,b,c)
```
### 通过关键字go就启动了goroutine。
```go
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i:=0;i <5;i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}
```
```
hello
world
hello
world
hello
world
hello
world
hello
```
> #### runtime.Gosched()表示让CPU把时间片让给别人，下次某个时候继续恢复执行该goroutine。
> #### 默认情况下，在Go1.5将表示并发系统线程个数的runtime.GOMAXPROCS的初始值由1改为了运行环境的CPU核数。
## channels
> ### goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据通信呢？go提供了一个很好的通信机制channel。
```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```
#### channel通过操作符<-来接收和发送数据。
```go
ch <- v //发送v到channel ch
v := <- ch // 从ch中接收数据，并赋值给v
```
```go
package main
import (
	"fmt"
)

func sum(a []int,c chan int) {
	total := 0
	for _,v := range a {
		total += v
		fmt.Println("total",total)
	}
	c <- total
}
func main() {
	a := []int{7,2,8,-9,4,0}
	c := make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	x,y := <-c,<-c
	fmt.Println(x,y,x+y)
}
```
```
total -9
total -5
total -5
total 7
total 9
total 17
-5 17 12
```
## Buffered Channels
> ### 上面我们介绍了默认的非缓存类型的channel，不过Go也允许指定channel的缓冲大小，很简单，就是channel可以存储多少元素。ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel。在这个channel 中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。

```go
ch := make(chan type,value)
```
#### 当value=0时，channel时无缓冲阻塞读写的，当value>0时，channel有缓冲，是非阻塞的，直到写满value个元素才阻塞写入。
```go
package main
import "fmt"
func main(){
    c := make(chan int,2)
    c <- 1
    c <- 2
    fmt.Println(<-c)
    fmt.Println(<-c)
}
        //修改为1报如下的错误:
        //fatal error: all goroutines are asleep - deadlock!
```
## Range和Close
```go
package main
import (
	"fmt"
)
func fibonacci(n int,c chan int) {
	x,y := 1,1
	for i := 0;i < n;i++{
		c <- x
		x,y = y,x+y
	}
}
func main() {
	c := make(chan int,10)
	go fibonacci(cap(c),c)
	for i := range c {
		fmt.Println(i)
	}

}
```
```
➜  并发 git:(master) ✗ go run go_channel2.go 
1
1
2
3
5
8
13
21
34
55
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/fun/go/src/build-web-application/并发/go_channel2.go:15 +0xec
exit status 2
```
> #### 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic。
> #### 另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确定没有任何发送数据了，或者你想结束range循环之类的。
##  select
> ### 当出现多个channel的时候，通过select可以监听channel上的数据流动。select默认是阻塞的，只有当监听的channel中有发送或者接收才会运行，当多个channel都准备好的时候，select是随机选择一个执行的。
```go
package main
import (
	"fmt"
)
func fibonacci(c, quit chan int) {
	x,y := 1,1
	for {
		select {
		case c <- x:
			x,y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0;i < 10;i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c,quit)
}
```
> #### select还有default，类似于switch。
## 超时
> ### 我们可以利用select来设置超时
```go
func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
				case v := <- c :
					println(v)
				case <- time.After(S * time.Second):
					println("timeout")
					o <- true
					break
			}
		}
	}
}
```
## runtime goroutine
> ### runtime包中有几个处理goroutine的函数：
-  #### Goexit
##### 退出当前执行的goroutine，但是defer函数还会继续调用。
- #### Gosched
##### 让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
- #### NumCPU
##### 返回CPU核心数
- #### NumGoroutine
##### 返回正在执行和排队的任务总数
- #### GOMAXPROCS
#####用来设置可以并行计算的CPU核数的最大值，并返回之前的值。