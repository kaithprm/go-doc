# go-doc Go语言编程阅读总结
## 一.go项目目录结构
### [原文地址](https://juejin.cn/post/7103440474152632328)
* ![GO](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/bd7b248fc6864f95b191b059133b947b~tplv-k3u1fbpfcp-zoom-in-crop-mark:1512:0:0:0.awebp?)
### 重要Go package解释
* 1./web
前端代码存放目录。
存放Web 应用程序特定的组件：主要有静态 Web 资源，服务器端模板和单页应用（Single-Page App，SPA）等。
* 2./cmd
存放当前项目的可执行文件。
cmd 目录下的每一个子目录名称都应该匹配可执行文件。例如，把组件 main 函数所在的文件夹统一放在 /cmd 目录下。
不要在 /cmd 目录中放置太多的代码，我们应该将公有代码放置到 /pkg 中，将私有代码放置到 /internal 中并在 /cmd 中引入这些包，保证 main 函数中的代码尽可能简单和少。
* 3./internal
存放私有应用和库代码。
如果一些代码，你不希望被其他项目/库导入，可以将这部分代码放至/internal目录下。一般存储一些比较专属于当前项目的代码包。这是在代码编译阶段就会被限制的，该目录下的代码不可被外部访问到。一般有以下子目录：
在/internal目录下应存放每个组件的源码目录，当项目变大、组件增多时，扔可以将新增的组件代码存放到/internal目录下internal目录并不局限在根目录，在各级子目录中也可以有internal子目录，也会同样起到作用。
* 4./router
路由
* 5./application
存放命令与查询
* 6./command
query
* 7./middleware 中间件
* 8./model 模型定义
* 9./repository 仓储层，封装数据库操作
* 10./response 响应
* 11./errmsg 错误处理
* 12./pkg
存放可以被外部应用使用的代码
/pkg目录下时可以被其他项目引用的包，所以我们将代码放入该目录下时候一定要慎重。
在非根目录的下也是可以很好的加入pkg目录，很多项目会在internal目录下加入pkg表示内部共享包库。
个人建议：一开始将所有的共享代码存放在/internal/pkg目录下，当确认可以对外开发时，再转至到根目录的/pkg目录下
* 13./vendor
存放项目依赖
可以通过命令行go mod wendor创建
如果创建的是一个Go库，不要提交wendor依赖包
* 14./third_party
存放放一些第三方的资源工具文件。
* 15./test
存放整个应用的测试、测试数据及一些集成测试等，
相较于单元测试在每个go文件对应的目录下，test目录偏向于整体
在某些子项目内也会有局部项目的测试会放在子项目的test中。
需要Go忽略该目录中的内容，可以使用/test/data或/test/testdata目录下
Go会忽略.或_开头的目录或文件
* 16./config或/configs
配置文件或者配置文件模板所在的文件夹。
配置中不能携带敏感信息，可用占位符代替
* 17./init
存放初始化系统和进程管理配置文件
/deployments 或 /deploy
存放 Iaas、PaaS 系统和容器编排部署配置和模板。
### 2.项目管理
* 1./Makefile
对项目进行管理
执行静态代码检查、单元测试、编译等功能。
* 2./build
存放安装包和持续集成相关的文件。
* 3./website
如果不使用 Github pages，则在这里放置项目的网站数据。
* 4.assets
项目使用的其他资源 (如图片等)。
* 5.tools
存放这个项目的支持工具。
## 二、go mod及Makefile对项目管理
### 1.Makefile语法规则
	目标 ... : 依赖 ...
		命令1
		命令2
		. . .


