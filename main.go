package main

import (
	"SynergizeTest/config"
	"SynergizeTest/delivery"

	"gorm.io/gorm"
)

var db gorm.DB

func main() {
	config.Init()
	delivery.InitRouter()
}
