package main

import (
	"fmt"
	"log"
	"net/http"
	"tmi-native/config"
	"tmi-native/controller"
	client "tmi-native/db"
	sqlClient "tmi-native/db/sqlc"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	client.DB = sqlClient.New(db)
	fmt.Println("Success")
	router := httprouter.New()
	router.POST("/admin/tantangan", controller.PostTantangan)
	router.POST("/admin/silabus", controller.PostSilabus)
	router.POST("/user/register", controller.PostUsers)
	router.POST("/user/login", controller.GetUserByEmailAndPassword)
	router.GET("/user/cek-user/:id", controller.GetUsersByID)
	log.Fatal(http.ListenAndServe(":8080", router))
}
