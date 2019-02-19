# Go的http包详解
## Conn的goroutine
### GO为了实现高并发和高性能，使用了goroutine来处理Conn的读写事件，这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。
```go
Go在等待客户端请求
c, err := srv.newConn(rw)
if err != nil {
    continue
}
go c.server()
```
### 每一次请求都会创建一个Conn，这个Conn里面保存了该次请求的信息，然后再传递到对应的handler，该handler中便可以读取到相应的header信息，保证了每个请求的独立性。
## ServerMux的自定义
### http默认路由
```go
type ServerMux struct {
    mu sync.RWMutex
    m map[string]muxEntry
    hosts bool
}
```
### 下面看一下muxEntry
```go
type muxEntry struct{
    explicit bool
    h Handler
    pattern string
}
```
### 接着看一下Handler的定义
```go
type Handler interface{
    ServerHTTP(ResponseWriter, *Request)
}
```
### Handler是一个接口，还需要定义一个类型HandlerFunc来实现ServerHTTP。
```go
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServerHTTP(w ResponesWriter,r *Request){
    f(w,r)
}
```
# 后续补充