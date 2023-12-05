# daily part1
配置mysql,测试dashboard
# daily part2
## 在linux系统中配置mysql 8.0
* 1.下载
```
# 先进入到local文件夹
cd usr/local

# 新建mysql文件夹
mkdir mysql

# 进入mysql文件夹
cd mysql

# 下载rpm，粘贴刚才复制的下载链接
wget https://dev.mysql.com/get/mysql80-community-release-el7-5.noarch.rpm
```
* 2.安装
```
# 通过本地的rpm进行安装源
yum localinstall -y mysql80-community-release-el7-5.noarch.rpm
# 在仓库中查询 mysql(可以找到刚下载的mysql源了)
yum search mysql-community
# 开始安装 mysql-community-server
yum install -y mysql-community-server
```
* 3.启动和检查
```
# 启动 mysql
systemctl start mysqld

# 查看mysql当前状态
systemctl status mysqld

# 设置开机启动
systemctl enable mysqld

# 找到默认密码(A temporary password)
vi /var/log/mysqld.log 
```
* 4.登录后修改密码
```mysql
alter user 'root'@'localhost' identified with mysql_native_password by 'your password';
```
* 5.设置root权限(远程登录)
```
update user set host = '%' where user = 'root';

#  使配置立即生效
flush privileges;
```
* 6.打开防火墙
```
# 添加放行3306端口
firewall-cmd --zone=public --permanent --add-port=3306/tcp

# 将防火墙配置重载
firewall-cmd --reload
```

