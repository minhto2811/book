package main

import (
	"book/api"
	"book/db"
	"book/utils"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	utils.ValidateInit()
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		return
	}
	db, err := database.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	routes := routes.Routes{Db: db}
	routes.Init()
}
