## 反射的应用场景
1. 结构体标签（Struct Tag）
在 Go 语言中，我们可以通过为结构体的字段添加标签（tag），来控制例如 JSON 编码/解码这样的行为。这就是一个反射的典型应用。
```
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    user := User{"Alice", 25}
    u := reflect.TypeOf(user)
    for i := 0; i < u.NumField(); i++ {
        field := u.Field(i)
        fmt.Printf("%s: %s\n", field.Name, field.Tag)
    }
}
```
