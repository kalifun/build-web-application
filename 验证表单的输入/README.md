# 验证表单的输入
> ### 开发web的一个原则就是，不能信任用户输入的任何信息，所以验证和过滤用户的输入信息变得非常重要。
## 必填字段
> ### 你想确保从一个表单元素中得到一个值。
```go
if len(r.Form["username"][0]) == 0 {

}
```
### r.Form对不同类型的表单元素的留空有不同的处理，对于空文本框，空文本区域以及文件上传，元素的值为空值，而如果是未选中的复选框和单选按钮，则根本不会在r.Form中产生相应条目。我们需要通过r.Form.Get()来获取值，因为如果不存在，通过该方式获取的是空值。但是通过r.Form.Get()只能获取单个的值。
## 数字
### 想确保一个表单输入框中获取的只能是数字。如果我们是判断正整数，那么我们先转化成int类型，然后在处理。
```go
getint,err := strconv.Atoi(r.Form.Get("age"))
if err != nil {

}
if getint > 100 {

}
```
### 还有一种方式就是正则表达式匹配
```go
if m,_ := regexp.MatchString("^[0-9]+$",r.Form.Get("age")); !m {
    return false
}
```
## 中文
### 有时候我们想通过表单元素获取一个用户的中文名，但是为了保证获取的是正确的中文，我们需要进行验证，而不是用户随便输入一些。对于中文有两种验证方式，可以使用Unicode包提供的func Is(rangeTab *RangeTable,r rune)bool来验证，也可以使用正则表达式来验证。
```go
if m,_ := regexp.MatchString("^\\p{Han}+$",r.Form.Get("realname")) ; !m {
    return false
}
```
## 英文
### 我们期望通过表单获取一个英文值。
```go
if m,_ := regexp.MatchString("^[a-zA-Z+$",r.Form.Get("engname")) ; !m {
    return false
}
```
## 电子邮件地址
```go
if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
	fmt.Println("no")
}else{
	fmt.Println("yes")
}
```
## 手机号码
```go
if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
	return false
}
```
## 下拉菜单
### 如果我们想要判断表单里<select>元素生成的下拉菜单中是否有被选中的项目。
```html
<select>
<option value="apple"> apple</option>
<option value="pear"> pear</option>
<option value="banana"> banana</option>
</select>
```
```go
slice := []string{"apple","pear","banana"}
v := r.Form.Get("fruit")
for _,item := range slice {
    if item == v {
        return true
    }
}
return false
```
## 单选按钮
### 如果我们想要判断radio按钮是否有一个被选中了。
```html
<input type="radio" name="gender" value="1">男
<input type="radio" name="gender" value="1">女
```
```go
slice:=[]string{"1","2"}

for _, v := range slice {
	if v == r.Form.Get("gender") {
		return true
	}
}
return false
```
## 复选框
```html
<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球
```
```go
slice:=[]string{"football","basketball","tennis"}
a:=Slice_diff(r.Form["interest"],slice)
if a == nil{
	return true
}

return false
```
## 日期和时间
### 想确定用户填写的日期或时间是否有效
```go
t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
fmt.Printf("Go launched at %s\n", t.Local())
```
## 省份证号码
```go
//验证15位身份证，15位的是全部数字
if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
	return false
}

//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
	return false
}
```