# daily part1

# daily part2
## 关于服务发现，服务注册
* ！[image](https://pic4.zhimg.com/80/v2-636cdc84e5139cf5380dd38c88c03b97_720w.webp)
* 注册中心，用于服务端注册远程服务以及客户端发现服务 zookeeper/etcd..
* 服务端，对外提供后台服务，将自己的服务信息注册到注册中心
* 客户端，从注册中心获取远程服务的注册信息，然后进行远程过程调用
## protocol生成的pb.go文件与_grpc.pb.go文件：
* pb.go 文件是 Protocol Buffers 的消息定义和序列化/反序列化代码。
* _grpc.pb.go 文件是 gRPC 的服务定义和客户端/服务器实现代码。它包含了根据 .proto 文件生成的 gRPC 服务端和客户端的接口和方法，用于在应用程序中创建和实现 gRPC 服务以及创建和调用 gRPC 客户端。
# stage
