package main

import (
	"main/controller"
	"main/database"
)

func main() {
	//go ws.InitWSClient()
	database.RedisInit()
	controller.Router()
}
