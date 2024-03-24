# 微服务网关
## 什么是API网关
### 定义
**API网关是一个服务器，是系统的唯一入口。**
### 概念
* 从面向对象设计的角度看，它与外观模式类似。
* 外观模式(Facade Pattern)：外部与子系统的通信通过一个统一的外观对象进行，为子系统中的一组接口提供一个统一的入口。
* 一个客户类需要和多个业务类交互，有时候这些需要交互的业务类会作为一个整体出现，这时引入一个新的外观类(Facade)来负责和多个业务类【子系统(Subsystem)】进行交互，而客户类只需与外观类交互。
* **API网关封装了系统内部架构，为每个客户端提供一个定制的API。**
## 本微服务网关想解决的问题
* 1.集成kubernetes
## Traefik
### 前置软件
* docker
* docker-compose
### 安装
* 1.创建docker-compose.yml文件，复制以下内容:
```yml
version: '3'

services:
  reverse-proxy:
    # The official v2 Traefik docker image
    image: traefik:v2.11
    # Enables the web UI and tells Traefik to listen to docker
    command: --api.insecure=true --providers.docker
    ports:
      # The HTTP port 端口映射
      - "80:80"
      # The Web UI (enabled by --api.insecure=true)
      - "8080:8080"
    volumes:
      # So that Traefik can listen to the Docker events 容器和主机之间的数据卷映射
      - /var/run/docker.sock:/var/run/docker.sock
```
* 2.在yml地址下执行以下命令:
```shell
docker-compose up -d reverse-proxy
```
* 3.验证
```
http://localhost:8080/api/rawdata
```
* 4.部署service
```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: whoami
  labels:
    app: whoami
spec:
  containers:
    - name: whoami
      image: traefik/whoami:latest
      ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: whoami
spec:
  ports:
    - port: 80
      protocol: TCP
      targetPort: 80
  selector:
    app: whoami
  type: ClusterIP
```
* 5.创建路由规则
* 创建完路由规则后外部可以访问 通过扩展CRD的方式来扩展
```yaml
# cat ingressroute.yaml 
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: whoami-route
spec:
  entryPoints:
  - web
  routes:
  - match: Host(`whoami.coolops.cn`)
    kind: Rule
    services:
      - name: whoami
        port: 80 
```
* 6.定义入口点（端口号和对应的端口名称）
启动 Traefik 时会定义入口点（端口号和对应的端口名称），这时 Kubernetes 集群外部就可以通过访问 Traefik 服务器地址和配置的入口点对 Traefik 服务进行访问。在访问时一般会带上“域名”+“入口点端口”，然后 Traefik 会根据域名和入口点端口在 Traefik 路由规则表中进行匹配，如果匹配成功，则将流量发送到 Kubernetes 内部应用中与外界进行交互。其中，域名与入口点与对应后台服务关联的规则，即是 Traefik 路由规则。
### 组件
### 创建路由规则
* 1.原生Ingress写法
* 2.使用CRD IngressRoute方式
* 3.使用GatewayAPI的方式
## k3s设置
## 安装
* 1.运行安装版本
```shell
curl -sfL https://rancher-mirror.rancher.cn/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn sh -
```
* 2.初始化脚本
```shell
curl -sfL https://rancher-mirror.rancher.cn/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn K3S_URL=https://myserver:6443 K3S_TOKEN=mynodetoken sh -
```
## 问题
* 使用kubectl get pod验证时 显示api server无法访问
* 解决方法:可能是由于初始化时没完全启动k3s组件，使用以下步骤重启k3s后该问题不再存在
```shell
sudo k3s-killall.sh
```
```shell
sudo k3s server
```
## K3s资源
一些资源的解释
### 工作负载型资源(workload):
* Pods：最小的部署单位，每个Pod包含一个或多个容器。
* 部署（Deployments）：管理Pod的创建和更新。
### 服务发现及负载均衡型资源(ServiceDiscoveryLoadBalance): 
* 服务（Services）：定义了如何访问Pod，例如负载均衡和服务发现。
### 集群级资源：
* 节点（Nodes）：集群的物理或虚拟机器。
* Namespace
* ClusterRole：角色 用来绑定到应用程序来进行权限管理
### 特殊类型的存储卷
* ConfigMap : 当配置中心来使用的资源类型
