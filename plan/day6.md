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
*server端 
```
  type server struct {
	pb.UnimplementedUserInfoServiceServer // !!必须嵌入 UnimplementedUserInfoServiceServer 才能实现向前兼容。就是说必须嵌入该方法才能继承接口
}

var u = server{}

func (s *server) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	fmt.Println("get function begin")
	name := req.Name
	// 数据里查用户信息
	if name == "zs" {
		resp = &pb.UserResponse{
			Id:   1,
			Name: name,
			Age:  22,
		}
	}
	return
}

func main() {
	addr := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("tcp listener default :", err)
	}
	fmt.Println("listener starting,port is :", addr)
	s := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(s, &u)
	s.Serve(listener)
}
 ```

* client端
```
func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("connect failed,err:", err)
	}
	defer conn.Close()
	client := pb.NewUserInfoServiceClient(conn)
	req := new(pb.UserRequest)
	req.Name = "zs"
	response, err := client.GetUserInfo(context.Background(), req) //Background 返回一个非 nil、空的 Context。
	if err != nil {
		fmt.Println("响应异常 ", err)
	}
	fmt.Printf("响应结果： %v\n", response)

}
```
# daily part3
实现与js交互，写持久层逻辑
