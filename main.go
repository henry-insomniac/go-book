package main

import (
	"fmt"
	"github.com/henry-insomniac/go-book/model"
	"log"

	// 导入 database 包
	"github.com/henry-insomniac/go-book/database"
	"github.com/henry-insomniac/go-book/router"
)

func main() {
	// 初始化数据库连接
	database.InitDB()

	// 自动迁移：GORM 会根据模型自动创建或更新数据库表
	if err := database.DB.AutoMigrate(&model.Book{}, &model.User{}); err != nil {
		log.Fatalf("Auto migration failed: %v", err)
		return
	}

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	if err := r.Run(":8087"); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
