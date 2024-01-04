package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"os/signal"
	"time"
)

type Received struct {
	Action string `json:"action"`
	Cnt    int    `json:"cnt"`
}

var rec Received

// mqtt client
func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("订阅: 当前话题是 [%s]; 信息是 [%s] \n", msg.Topic(), msg.Payload())
	err := json.Unmarshal(msg.Payload(), &rec)
	if err != nil {
		log.Println(err)
	}
	// 现在rec中应该包含了解析后的数据
	if rec.Action == "acc" {
		message := fmt.Sprintf("嗨，这趟行程您已经急加速 %d 次了，慢下来也许会收获更美的旅程哦～", rec.Cnt)
		token := client.Publish("vehiceEventChars/SGL0000413032020C", 0, false, message)
		token.Wait()
	} else if rec.Action == "dec" {
		message := fmt.Sprintf("嗨，这趟行程您已经急减速 %d 次了，慢下来也许会收获更美的旅程哦～", rec.Cnt)
		token := client.Publish("vehiceEventChars/SGL0000413032020C", 0, false, message)
		token.Wait()
	}
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("47.94.135.4:11883")
	opts.SetClientID("go-mqtt-client")
	opts.SetDefaultPublishHandler(onMessageReceived) // 设置默认的消息处理函数

	// 创建 MQTT 客户端实例
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	// 在连接成功后进行订阅发布
	go func() {
		// 1.订阅主题
		if token := client.Subscribe("vehiceEvent/SGL0000413032020C", 0, nil); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}

		// 2.收到并解析json -- > 在onMessageReceived

		// 3. 判断 action并发布消息

		time.Sleep(time.Second)

	}()

	// 等待退出信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// 断开与 MQTT 服务器的连接
	client.Disconnect(250)
}
