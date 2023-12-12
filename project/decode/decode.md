# 关于编解包项目的总结
## 项目主要目标
* 本项目根据定义的文档，完成自定协议的编解包
## 得到经验
* 1.位运算
* 2.io读写
* 3.go test
* 4.debug
* 5.切片性质
* 6.大小端
## 具体经验


### 2.io读写
```
var reader io.Reader
bufreader := bufio.NewReader(reader)
n, err1 := bufreader.Read(result)
if err1 != nil {
println(err1)
}
```

```
var buf bytes.Buffer
buf.Write(result)
reader := &buf

p1, err := Decode(reader)
```
* 问题出在bufio.NewReader(reader)这一行。具体来说，问题在于在这里尝试对一个未初始化的io.Reader进行读取操作，这会导致空指针引用错误。在这里，reader是一个未初始化的接口类型变量，尝试使用bufio.NewReader对其进行初始化是不正确的。
* 而在第二个代码片段中，您首先创建了一个bytes.Buffer类型的变量buf，并使用buf.Write(result)将数据写入到缓冲区中，然后将buf用作io.Reader类型的变量进行解码操作。这是一个正确的做法，因为bytes.Buffer类型实现了io.Reader接口，可以直接用于读取数据
* 

### 5.切片性质
* 当切片由如下初始化时，切片需要一个空间
```
var slice []uint32{1,2,3,4}
```
* 当切片由如下初始化时，初始化需要的空间是3个，因为需要一个位置来存储len和cap，另外还需要一个位置来存储切片中的元素。
```
a := []uint32{1}
```
