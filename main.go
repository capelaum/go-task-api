package main

import (
	"fmt"
	"log"
	"net/http"

	"./controller"
	"./model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":3000", mux))
	fmt.Println("Serving on port :3000")
}
