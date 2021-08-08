package handler

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := getEnv("DB_USER", "root")
	PASSWORD := getEnv("DB_PASSWORD", "password")
	HOSTNAME := getEnv("DB_HOSTNAME", "localhost")
	DB_NAME := getEnv("DB_NAME", "examples")
	PORT := getEnv("DB_PORT", "3306")
	PROTOCOL := fmt.Sprintf("tcp(%v:%v)", HOSTNAME, PORT)

	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DB_NAME

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		return nil
	}

	return db
}
