package main

import (
	"log"

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
	logrus.Info("Starting database migration...")

	// 初始化数据库连接
	config.InitDatabase()
	db := config.GetDB()

	// 执行数据库迁移
	logrus.Info("Running database migrations...")
	
	err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	logrus.Info("Database migration completed successfully!")
	
	// 显示创建的表信息
	logrus.Info("Created tables:")
	logrus.Info("- users (用户表)")
	logrus.Info("- posts (文章表)")
	logrus.Info("- comments (评论表)")
	
	// 检查表是否存在
	if db.Migrator().HasTable(&models.User{}) {
		logrus.Info("✓ users table created successfully")
	}
	if db.Migrator().HasTable(&models.Post{}) {
		logrus.Info("✓ posts table created successfully")
	}
	if db.Migrator().HasTable(&models.Comment{}) {
		logrus.Info("✓ comments table created successfully")
	}
}