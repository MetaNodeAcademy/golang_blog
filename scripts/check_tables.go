package main

import (
	"fmt"

	"golang_blog/config"
	"golang_blog/models"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found, using system environment variables")
	}

	// 初始化日志
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Checking database tables...")

	// 初始化数据库连接
	config.InitDatabase()
	db := config.GetDB()

	// 检查表是否存在
	tables := []struct {
		model interface{}
		name  string
		desc  string
	}{
		{&models.User{}, "users", "用户表"},
		{&models.Post{}, "posts", "文章表"},
		{&models.Comment{}, "comments", "评论表"},
	}

	fmt.Println("\n=== 数据库表检查结果 ===")
	
	for _, table := range tables {
		if db.Migrator().HasTable(table.model) {
			fmt.Printf("✓ %s (%s) - 存在\n", table.name, table.desc)
			
			// 获取表的列信息
			columnTypes, err := db.Migrator().ColumnTypes(table.model)
			if err != nil {
				logrus.WithError(err).Errorf("Failed to get column types for %s", table.name)
				continue
			}
			
			fmt.Printf("  列信息:\n")
			for _, col := range columnTypes {
				fmt.Printf("    - %s: %s\n", col.Name(), col.DatabaseTypeName())
			}
			fmt.Println()
		} else {
			fmt.Printf("✗ %s (%s) - 不存在\n", table.name, table.desc)
		}
	}

	// 检查外键关系
	fmt.Println("=== 外键关系检查 ===")
	
	// 检查posts表的user_id外键
	var postCount int64
	db.Model(&models.Post{}).Count(&postCount)
	fmt.Printf("posts表记录数: %d\n", postCount)
	
	// 检查comments表的外键
	var commentCount int64
	db.Model(&models.Comment{}).Count(&commentCount)
	fmt.Printf("comments表记录数: %d\n", commentCount)
	
	var userCount int64
	db.Model(&models.User{}).Count(&userCount)
	fmt.Printf("users表记录数: %d\n", userCount)

	logrus.Info("Database table check completed!")
}