package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
	"fmt"
	//"github.com/astaxie/beego/logs"
)

func main() {

	var dbConfig = map[string]interface{}{
		"Default":         "mysql_dev", // 默认数据库配置
		"SetMaxOpenConns": 0,           // (连接池)最大打开的连接数，默认值为0表示不限制
		"SetMaxIdleConns": 1,           // (连接池)闲置的连接数, 默认1

		"Connections": map[string]map[string]string{
			"mysql_dev": {// 定义名为 mysql_dev 的数据库配置
				"host": "192.168.4.30", // 数据库地址
				"username": "root",     // 数据库用户名
				"password": "111111",   // 数据库密码
				"port": "3306",         // 端口
				"database": "test",     // 链接的数据库名字
				"charset": "utf8",      // 字符集
				"protocol": "tcp",      // 链接协议
				"prefix": "",           // 表前缀
				"driver": "mysql",      // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
			},
			"sqlite_dev": {
				"database": "./foo.db",
				"prefix":   "",
				"driver":   "sqlite3",
			},
		},
	}

	// 初始化数据库链接, 默认会链接配置中 default 指定的值
	// 也可以在第二个参数中指定对应的数据库链接, 见下边注释的那一行链接示例
	connection, err := gorose.Open(dbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	// close DB
	defer connection.Close()

	db := connection.NewDB()

	//db.Begin()
	//
	//list := make([]map[string] interface{}, 1000)
	//
	//for i := 0; i < 1000; i++ {
	//	logs.Info("i: %d", i)
	//	list[i] = map[string]interface{}{"id": i, "name": "ss"}
	//}
	//
	//res2, err := db.Table("user_info").Data(list).Insert()
	//
	//if (res2 == 0 || err != nil) {
	//	logs.Info("发生一个错误: %s", err.Error())
	//	db.Rollback()
	//}
	//
	//db.Commit()

	res, err := db.Table("user_info").Where("id", ">", 33).Limit(10).Get()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
