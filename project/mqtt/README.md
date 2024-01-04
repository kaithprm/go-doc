# 关于mqtt订阅收发消息

```
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
```
* 根据发送的消息做反应应基于如上形式 而不是加for进行监听
## 关于json结构体转换
```
type Received struct {
	Action string `json:"action"`
	Cnt    int    `json:"cnt"`
}
var rec Received
err := json.Unmarshal(msg.Payload(), &rec)
```
