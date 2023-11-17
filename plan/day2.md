# day2 part1
* 学习了使用命令行操作mysql
* 命令行相较于可视化软件操作mysql优点：对于学习来说能够更加清晰理解mysql操作 缺点：操作数据表太慢了
```
mysql -u username -p
# 登录
show databases
# 显示数据库列表
use datebasename
# 选择数据库
show tables
# 查看数据库表
describe tablename
# 查看表结构
```
* go init函数机制: 在main文件之前执行，可完成注册等
* 编写了auth模块，使用go实现了mysql CRUD
* 在实现CRUD时 借鉴rest api 应该比较可读
```
func Post(){} # Create
func Get(){} # Retrieve
func Put(){} # Update
func Delete(){} 
```
# day2 part2
* docker的部署可以延后
* web与mqtt创建为两个项目，两个项目分别启动两个mysql服务即可，在mqtt端关闭对数据库的增删改即没有数据不一致等问题
* web端提供一个register button
# day2 part3
