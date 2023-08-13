package config

import (
	"SynergizeTest/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func initPostgres() (err error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "irvandy2"
		password = "koinworks"
		dbname   = "Synergize"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "users.", // schema name
		SingularTable: false,
	}})

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")
	return
}

func Paginate(form models.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (form.Page - 1) * form.Limit
		return db.Offset(offset).Limit(form.Limit)
	}
}
