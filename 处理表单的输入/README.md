# 处理表单的输入
> ## 命名文件login.gtpl,放入当前新建项目的目录里
```html
<html>
<head>
<title>TEST</title>
</head>
<body>
<form action="/login" method="post">
用户名:<input type="text" name="usename">
密码:<input type="password" name="password">
<input type="submit" value="登录">
</form>
</body>
</html>
```
## 处理login页面的form数据。
```go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)
func sayhelloName(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello world!")
}

func login(w http.ResponseWriter,r *http.Request) {
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	}else{
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
	}
}
func main() {
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServer:",err)
	}
}
```
### 获取请求方法是通过r.Method来完成的，这是一个字符串类型的变量，返回GET,POST,PUT等method信息。
### 我们输入用户名和密码后发现在服务器端是不会打印出来任何输出的，默认情况下，handler里面是不会自动解析form的，必须显示的调用r.ParseForm()后，你才能对这个表单数据进行操作。
```go
r.ParseForm()
fmt.Println("username:",r.Form["username"])
```
### 重新跑程序，终端就可以看到信息啦。r.Form里面包含了所有请求的参数，比如URL中query-string，post的数据，put的数据，所以当你在URL中的query-string字段和post冲突时，会保存在一个slice，里面储存了多个值。
### 现在我们将logi.gtpl修改一下，查看服务器输出的是不是一个slice。
```html
<form action="/login?username=kali" method="post">
```
```
method: POST
username: [kali]
password: [123456]
```
### request.Form是一个url.Values类型，里面存储的是对应的类似key=value的信息
```go
v := url.Values{}
v.Set("name","Ava")
v.Add("friend","Jess")
v.Add("friend","Sarah")
v.Add("friend","Zoe")
fmt.Println(v.Get("name"))
fmt.Println(v.Get("friend"))
fmt.Println(v["friend"])
```