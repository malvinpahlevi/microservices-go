package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

type MasterToken struct {
	ID        string    `json:"id" gorm:"type:char(36);primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Token     string    `json:"token" gorm:"type:text"`
	IssuedAt  time.Time `json:"issued_at, omitempty"`
	ExpiredAt time.Time `json:"expired_at, omitempty"`
}

func main() {

	DB, err := gorm.Open(mysql.Open("appuser:admin@tcp(127.0.0.1:3306)/boilerplate_go?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
	DB.AutoMigrate(&MasterToken{})
	fmt.Println("Success Migrated !")
}
