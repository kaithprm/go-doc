# daily part1

# daily part2
* proto文件用后方标识号传递，不用字段名传递
```
message UserResponse{
  int32 id(字段名) = 1(标识号);
  string name = 2;
  int32 age = 3;
}
```
* 关于proto生成文件类型思考：
* grpc相当于一个C/S结构，当使用proto生成文件时，只需在CS两端分别生成跟两端语言对应的文件类型即可
# daily part3
