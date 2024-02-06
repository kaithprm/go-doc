# 用go语言实现一个图书管理系统过程
## 1.http server搭建
```
http.ListenAndServe(":8080", handler)
```
* 该方法完成监听端口(listen)并处理进入端口的连接(serve)
### 关于handler
* 设置handler就是设置处理的程序
* 路由的概念:当接收到客户端发来的HTTP请求，会根据请求的URL，来找到相应的映射函数，然后执行该函数，并将函数的返回值发送给客户端.
* 定义handler仅需要实现ServeHTTP(ResponseWriter, *Request)接口即可
```
// 设置路由
type handlerImp struct {
}

func (imp handlerImp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.ListenAndServe(":8080", handlerImp{})
}
```
