package main

import (
	"tmi/config"
	"tmi/routes"
)

func main() {
	db := config.ConnectDatabase()
	sqlDB, _ := db.DB()

	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
