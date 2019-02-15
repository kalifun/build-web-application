# 你好,GO
## 我只想学Hello，world
> ### 学习语言的开始都是从Hello，World
```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello,World!")
}
```
```
Hello,World!
```
## 详解
-  package pkhname 这是告诉我们这是属于哪个包，当我们用main则代表可以单独运行程序。当编译后最后都会成为*.a的文件。
- import  这就是可以理解Python中的import。因为我们需要使用printf，需要使用fmt这个包。但是要注意的是，go和Python不一样。当你不需要使用的包时请不需要加入，不然无法编译过。