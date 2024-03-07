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
### 2.注册服务
* 1.注册服务代码
```go
/**
 * 输入指令
 */
func enterInstruction() int {
	var x int
	fmt.Println("请输入指令")
	fmt.Scan(&x)
	return x
}

/**
 * 流程判断
 */
func process(instruction int) {

	switch {
	case instruction == 1:
		t1()
	case instruction == 2:
		t2()
	case instruction == 3:
		t3()
	case instruction == 4:
		t4()
	case instruction == 5:
		t5()
	case instruction == 6:
		t6()
	case instruction == 7:
		t7()
	case instruction == 8:
		t8()
	case instruction == 9:
		t9()
	case instruction == 10:
		t10()
	case instruction == 11:
		t11()
	case instruction == 12:
		t12()
	case instruction == 13:
		t13()
	case instruction == 14:
		t14()
	case instruction == 15:
		t15()
	}

}

func main() {

	//初始化Nacos客户端
	intiNacosClient()

	for {
		instruction := enterInstruction()
		fmt.Printf("当前指令==>:%d\n", instruction)
		if instruction == 0 {
			return
		}
		process(instruction)
	}

}

/**
 * 注册默认集群和组
 * ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
 */
func t1() {
	ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo1.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})
}

/**
 * 使用集群名称注册
 * GroupName=DEFAULT_GROUP
 */
func t2() {
	ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-a",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
}

/**
 * 注册不同的集群
 * GroupName=DEFAULT_GROUP
 */
func t3() {
	ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.12",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-b",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
}

/**
 * 注册不同的组
 */
func t4() {
	ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.13",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-b",
		GroupName:   "group-a",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})

	ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.14",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		ClusterName: "cluster-b",
		GroupName:   "group-b",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
}

/**
 * 注销 使用 ip,port,serviceName 注销  ClusterName=DEFAULT, GroupName=DEFAULT_GROUP
 * Note:ip=10.0.0.10,port=8848 应该属于集群 DEFAULT and the group of DEFAULT_GROUP.
 */
func t5() {
	ExampleServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demo.go",
		Ephemeral:   true, //it must be true
	})
}

/**
 * 注销 使用 ip,port,serviceName,cluster 注销 GroupName=DEFAULT_GROUP
 * Note:ip=10.0.0.10,port=8848,cluster=cluster-a 应该属于 DEFAULT_GROUP.
 */
func t6() {
	ExampleServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo.go",
		Cluster:     "cluster-a",
		Ephemeral:   true, //it must be true
	})
}

/**
 * 注销 ip,port,serviceName,cluster,group
 */
func t7() {
	ExampleServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.14",
		Port:        8848,
		ServiceName: "demo.go",
		Cluster:     "cluster-b",
		GroupName:   "group-b",
		Ephemeral:   true, //it must be true
	})
}

/**
 * 使用 serviceName 获取服务
 * ClusterName=DEFAULT, GroupName=DEFAULT_GROUP
 */
func t8() {
	ExampleServiceClient_GetService(client, vo.GetServiceParam{
		ServiceName: "demo.go",
	})
}

/**
 * 使用 serviceName 和 cluster 获取服务
 * GroupName=DEFAULT_GROUP
 */
func t9() {
	ExampleServiceClient_GetService(client, vo.GetServiceParam{
		ServiceName: "demo.go",
		Clusters:    []string{"cluster-a", "cluster-b"},
	})
}

/**
 * 获取服务 serviceName
 * ClusterName=DEFAULT
 */
func t10() {
	ExampleServiceClient_GetService(client, vo.GetServiceParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
	})
}

/**
 * 查询所有服务 返回所有实例 ,包括 healthy=false,enable=false,weight<=0
 * ClusterName=DEFAULT, GroupName=DEFAULT_GROUP
 */
func t11() {
	ExampleServiceClient_SelectAllInstances(client, vo.SelectAllInstancesParam{
		ServiceName: "demo.go",
	})
}

/**
 * 查询所有服务
 * GroupName=DEFAULT_GROUP
 */
func t12() {
	ExampleServiceClient_SelectAllInstances(client, vo.SelectAllInstancesParam{
		ServiceName: "demo.go",
		Clusters:    []string{"cluster-a", "cluster-b"},
	})
}

/**
 * 查询所有服务
 * ClusterName=DEFAULT
 */
func t13() {
	ExampleServiceClient_SelectAllInstances(client, vo.SelectAllInstancesParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
	})
}

/**
 * 查询说有实例 只返回的实例 healthy=${HealthyOnly},enable=true and weight>0
 * ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
 */
func t14() {
	ExampleServiceClient_SelectInstances(client, vo.SelectInstancesParam{
		ServiceName: "demo.go",
	})
}

/**
 * 选择一个健康的实例 通过 WRR 策略返回一个实例进行负载均衡
 * 并且实例应该是 health=true,enable=true and weight>0
 * ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
 */
func t15() {
	ExampleServiceClient_SelectOneHealthyInstance(client, vo.SelectOneHealthInstanceParam{
		ServiceName: "demo.go",
	})
}
func ExampleServiceClient_RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, _ := client.RegisterInstance(param)
	fmt.Printf("注册服务实例:%+v,result:%+v \n\n", param, success)
}

func ExampleServiceClient_DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	success, _ := client.DeregisterInstance(param)
	fmt.Printf("注销服务实例:%+v,result:%+v \n\n", param, success)
}

func ExampleServiceClient_GetService(client naming_client.INamingClient, param vo.GetServiceParam) {
	service, _ := client.GetService(param)
	fmt.Printf("获取服务:%+v, result:%+v \n\n", param, service)
}

func ExampleServiceClient_SelectAllInstances(client naming_client.INamingClient, param vo.SelectAllInstancesParam) {
	instances, _ := client.SelectAllInstances(param)
	fmt.Printf("选择所有实例:%+v, result:%+v \n\n", param, instances)
}

func ExampleServiceClient_SelectInstances(client naming_client.INamingClient, param vo.SelectInstancesParam) {
	instances, _ := client.SelectInstances(param)
	fmt.Printf("选择实例:%+v, result:%+v \n\n", param, instances)
}

func ExampleServiceClient_SelectOneHealthyInstance(client naming_client.INamingClient, param vo.SelectOneHealthInstanceParam) {
	instances, _ := client.SelectOneHealthyInstance(param)
	fmt.Printf("选择实例:%+v, result:%+v \n\n", param, instances)
}

func ExampleServiceClient_Subscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Subscribe(param)
}

func ExampleServiceClient_UnSubscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Unsubscribe(param)
}

func ExampleServiceClient_GetAllService(client naming_client.INamingClient, param vo.GetAllServiceInfoParam) {
	service, _ := client.GetAllServicesInfo(param)
	fmt.Printf("获取所有服务:%+v, result:%+v \n\n", param, service)
}
```
## 网关实现
### 1.实现反向代理
* 连接nacos，根据name获取ip
