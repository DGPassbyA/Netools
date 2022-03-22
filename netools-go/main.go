package main

import (
	"main/controller"
	"main/database"
)

func main() {
	database.RedisInit()
	database.SqliteInit()
	database.SqliteInitSchema()
	controller.Router()
}
