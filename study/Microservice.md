# 微服务概念
* 使用一套小服务来开发单个应用的方式，每个服务运行在独立的进程里，一般采用轻量级的通讯机制互联(RPC)，并且它们可以通过自动化的方式部署
# 微服务生态
## 1.通信层
* 1.1 网络传输，用RPC（远程过程调用）
* HTTP传输，GET POST PUT DELETE
* 基于TCP，更靠底层，RPC基于TCP，Dubbo（18年底改成支持各种语言），Grpc，Thrift
* 1.2 服务注册和发现
* 需要分布式数据同步：etcd，consul，zk
## 2.应用平台层
* 云管理平台、监控平台、日志管理平台，需要他们支持
* 服务管理平台，测试发布平台
* 服务治理平台
# RPC
关于RPC与protocol的理解:RPC是实现远程方法调用 而protocol是调用方与被调用方交互的一种序列化形式
# protocol
* [](https://blog.csdn.net/qq_31347869/article/details/93189219)
