package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
func Start()  {
	conf:=fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", "localhost",5432, "goland",
		"gorm", "goland")

	db, err := gorm.Open(postgres.Open(conf),&gorm.Config{
		DisableForeignKeyConstraintWhenMigrating:true,
	})

	if err != nil{
		panic(err)
	}
	DB = db
}
