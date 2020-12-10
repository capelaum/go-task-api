package main

// going to build this API: https://github.com/L04DB4L4NC3R/getgoing/tree/master/section4/todo-api

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
