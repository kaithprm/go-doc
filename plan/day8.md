# daily part1

# daily part2
* proto文件用后方标识号传递，不用字段名传递
*
```
message UserResponse{
  int32 id(字段名) = 1(标识号);
  string name = 2;
  int32 age = 3;
}
```
# daily part3
