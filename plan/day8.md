# daily part1
开始更深刻的了解rpc协议
# daily part2
## proto文件用后方标识号传递，不用字段名传递
```
message UserResponse{
  int32 id(字段名) = 1(标识号);
  string name = 2;
  int32 age = 3;
}
```
## 关于proto生成文件类型思考：
* grpc相当于一个C/S结构，当使用proto生成文件时，只需在CS两端分别生成跟两端语言对应的文件类型即可
## 为什么说mqtt是个轻量化的协议呢？那rpc协议是吗？对轻量化的定义是什么？
* mqtt的轻量化体现在：1.mqtt消息头部很小，协议本身的开销很小。2.低带宽和低资源消耗：MQTT协议使用二进制格式传输数据，消息头部只有几个字节，传输效率高。此外，MQTT协议具有低功耗和低网络带宽消耗的特点，适用于资源受限的设备和网络环境。
* RPC协议通常不被称为轻量级协议，因为它的实现和功能较为复杂，需要更多的网络带宽和资源。
* 对于轻量化的定义，通常是指协议或系统具有较小的开销、较低的资源消耗、高效的传输和处理能力，适用于资源受限的环境，同时提供足够的功能满足特定的需求。
## 原生rpc
* 使用gob/json编解码
### 原生rpc实现rpc通信为什么要建立在http协议上？
* 广泛支持：HTTP协议是互联网上最常用的协议之一，几乎所有的网络设备和应用程序都能够支持HTTP协议。因此，使用HTTP作为底层协议可以保证RPC通信能够在各种环境和设备上进行。
* 跨平台和语言支持：HTTP协议是一种跨平台和语言无关的协议，任何支持HTTP协议的平台和编程语言都可以使用HTTP进行通信。这使得基于HTTP的RPC通信可以在不同的平台和使用不同编程语言的应用程序之间进行通信。
### 原生rpc没有像grpc框架一样建立在tcp协议上而是要在封装一层http?
* 因为原生的net/rpc库是在Go语言的标准库中提供的，而Go语言标准库中的net/http包提供了HTTP服务器和客户端的实现，因此使用HTTP作为底层协议更加方便和简易。
### net/rpc **服务端**的注册源码
```
func (server *Server) register(rcvr any, name string, useName bool) error {
	s := new(service)
	s.typ = reflect.TypeOf(rcvr)  //保存对象的type
	s.rcvr = reflect.ValueOf(rcvr) //保存对象的value
	sname := name //保存对象的名称
	if !useName {
		sname = reflect.Indirect(s.rcvr).Type().Name()
	}
	if sname == "" {
		s := "rpc.Register: no service name for type " + s.typ.String()
		log.Print(s)
		return errors.New(s)
	}
	if !useName && !token.IsExported(sname) {
		s := "rpc.Register: type " + sname + " is not exported"
		log.Print(s)
		return errors.New(s)
	}
	s.name = sname

	// Install the methods
	s.method = suitableMethods(s.typ, logRegisterError)

	if len(s.method) == 0 {
		str := ""

		// To help the user, see if a pointer receiver would work.
		method := suitableMethods(reflect.PointerTo(s.typ), false)
		if len(method) != 0 {
			str = "rpc.Register: type " + sname + " has no exported methods of suitable type (hint: pass a pointer to value of that type)"
		} else {
			str = "rpc.Register: type " + sname + " has no exported methods of suitable type"
		}
		log.Print(str)
		return errors.New(str)
	}

	if _, dup := server.serviceMap.LoadOrStore(sname, s); dup {
		return errors.New("rpc: service already defined: " + sname)
	}
	return nil
}
```
# daily part3
继续了解rpc协议底层，搞清楚每一步的意义，如何实现的，仔细看一下注册代码
