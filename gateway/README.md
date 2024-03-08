# api gateway
## 实现网关前的准备 
### 1.使用docker搭建本地nacos
* 1.拉取nacos镜像
```shell
docker pull nacos/nacos-server
```
* 2.创建容器网络
* 实现容器之间的通信和数据交换。创建容器网络可以提高容器之间的隔离性，并简化容器的网络配置和管理。
```shell
docker network create nacos_network
```
* 3.启动nacos
* 这个命令会启动一个名为 nacos 的容器，并将其绑定到本地机器的 8848 端口。同时，它还会将容器添加到之前创建的 nacos_network 容器网络中，并设置容器模式为 standalone。
```shell
docker run --name nacos -d \
-p 8848:8848 \
--network nacos_network \
-e MODE=standalone \
nacos/nacos-server
```
* 4.再次启动
```shell
docker start -a nacos
```
### 2.注册服务
* 1.注册服务demo.go
## 网关实现
### 1.实现反向代理
* 连接nacos，根据name获取ip
