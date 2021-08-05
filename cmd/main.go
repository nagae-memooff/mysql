package main

import (
	"log"
	"time"

	// "database/sql"
	"github.com/nagae-memooff/dameng"
	"gorm.io/gorm"
)

func init() {
}

type Model struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT" json:"id" yaml:"id" gorm:"column:ID"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at" gorm:"column:CREATED_AT"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at" gorm:"column:UPDATED_AT"`
	//   DeletedAt *time.Time
}

type User struct {
	Model
	Name string `json:"name" gorm:"index;column:NAME"`
}

func main() {
	addr := "dm://esns:minxing123@192.168.100.92:5236"
	// addr := "esns:minxing123@(127.0.0.1:3306)/esns_development?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(dameng.Open(addr), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Printf("error: %v", err)
	}

	user := User{}

	db.First(&user)

	log.Printf("user: %v", user)
}
