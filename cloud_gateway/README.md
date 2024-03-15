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