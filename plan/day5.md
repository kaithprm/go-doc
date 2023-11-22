# daily part1

# daily part2
* 使用RPC有以下几个重要规则：
* 1.结构体字段首字母要大写，可以别人调用**结构体中的参数同样需要大写**
* 2.函数名必须首字母大写
* 3.函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型
* 4.函数还必须有一个返回值error
* 否则会出现以下错误
```
 gob: type main.Requests has no exported fields
```
# daily part3
