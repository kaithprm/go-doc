# daily part1
学习了powerpoint一些特性技巧，实现了grpc框架C&S
# daily part2
## 编写protocol
```
syntax = "proto3";

package proto;


message UserRequest{
  string name = 1;
}
message UserResponse{
  int32 id = 1;
  string name = 2;
  int32 age = 3;
}
service UserInfoService{
  rpc GetUserInfo(UserRequest)returns(UserResponse){
  }
}
```
## 编写server和client
# daily part3
实现与js交互，写持久层逻辑
