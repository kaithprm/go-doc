package service

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var client naming_client.INamingClient

/**for
 * 初始化Nacos客户端
 */
func InitNacosClient() {

	sc := []constant.ServerConfig{
		{
			IpAddr: "localhost", // Nacos的服务地址
			Port:   8848,        // Nacos的服务端口
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "",                                 // ACM的命名空间Id 当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,                               // 请求Nacos服务端的超时时间，默认是10000ms
		NotLoadCacheAtStart: true,                               // 在启动的时候不读取缓存在CacheDir的service信息
		LogDir:              "/Users/xu_nuo/mydata/nacos/logs",  // 日志存储路径
		CacheDir:            "/Users/xu_nuo/mydata/nacos/cache", // 缓存service信息的目录，默认是当前运行目录
		LogLevel:            "debug",                            // 日志默认级别，值必须是：debug,info,warn,error，默认值是info
	}

	// 创建命名客户端
	clientTemp, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	client = clientTemp
	println("初始化nacos成功")

}
func Query(name string) string {
	return ServiceClient_GetService(client, vo.GetServiceParam{
		ServiceName: name,
	})

}

// 获取服务
func ServiceClient_GetService(client naming_client.INamingClient, param vo.GetServiceParam) string {
	var IP string
	service, _ := client.GetService(param)
	fmt.Printf("获取服务:%+v, result:%+v \n\n", param, service)
	// 在这里 服务注册时目前只支持一个IP
	if len(service.Hosts) > 0 {
		for _, host := range service.Hosts {
			// 获取每个 Host 结构体的 Ip 字段
			IP = host.Ip
		}
	} else {
		println("Hosts is empty")
	}
	return IP
}
