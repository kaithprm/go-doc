package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

type Service struct {
	Route string
	Name  string
}

var (
	MySQLClient *sql.DB //连接池对象
)

func Init() {
	//数据库
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/nacos"
	//连接数据集
	client, err := sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = client.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功~")
	MySQLClient = client

}

func Query(route string) string {
	//1.查询单挑记录的sql语句  ?是参数
	sqlStr := "select * from route_nacos where route = ?;"
	//2.执行
	rowObj := MySQLClient.QueryRow(sqlStr, route) //从连接池里取一个连接出来去数据库查询单挑记录
	if rowObj == nil {
		fmt.Println("QueryRow returned nil")
		return "QueryRow returned nil"
	}
	////3.拿到结果
	var s1 Service
	rowObj.Scan(&s1.Route, &s1.Name)
	////打印结果
	fmt.Printf("s1 is:%#v\n", s1.Name)
	return s1.Name
}
