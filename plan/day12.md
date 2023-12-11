###关于切片的数据结构
* Slice 的数据结构定义如下:
```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
* [!imag1](https://img.halfrost.com/Blog/ArticleImage/57_2.png)
* 切片的结构体由3部分构成，Pointer 是指向一个数组的指针，len 代表当前切片的长度，cap 是当前切片的容量。cap 总是大于等于 len 的。
* [原文地址](https://halfrost.com/go_slice/#:~:text=%E6%B7%B1%E5%85%A5%E8%A7%A3%E6%9E%90%20Go%20%E4%B8%AD%20Slice%20%E5%BA%95%E5%B1%82%E5%AE%9E%E7%8E%B0%201%20%E4%B8%80.%20%E5%88%87%E7%89%87%E5%92%8C%E6%95%B0%E7%BB%84,...%205%20%E4%BA%94.%20%E5%88%87%E7%89%87%E6%8B%B7%E8%B4%9D%20Slice%20%E4%B8%AD%E6%8B%B7%E8%B4%9D%E6%96%B9%E6%B3%95%E6%9C%892%E4%B8%AA%E3%80%82%20Go%20)
