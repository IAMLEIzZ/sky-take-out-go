package main

import (
	"sky-take-out-go/db"
	"sky-take-out-go/router"
)

func main() {
	// 连接数据库
	db.InitDB()
	// 路由初始化
	r := router.InitRouter()
	err := r.Run(":8081")

	if err != nil {
		panic(err)
	}
}
