package handler

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func NewGormConnection() *gorm.DB {
	USER := getEnv("DB_USER", "root")
	PASSWORD := getEnv("DB_PASSWORD", "password")
	HOSTNAME := getEnv("DB_HOSTNAME", "localhost")
	DB_NAME := getEnv("DB_NAME", "examples")
	PORT := getEnv("DB_PORT", "3306")
	PROTOCOL := fmt.Sprintf("tcp(%v:%v)", HOSTNAME, PORT)

	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DB_NAME

	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		return nil
	}

	return db
}
