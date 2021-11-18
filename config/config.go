package config

import (
	"tmi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	/*
		username := "postgres"
		password := "12345678"
		host := "tcp(127.0.0.1:5432)"
		database := "db_jcc_tmi"
	*/
	dsn := "host=localhost user=postgres password=12345678 dbname=db_jcc_tmi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Mahasiswa{})

	return db
}
