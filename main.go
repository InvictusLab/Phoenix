package main

import (
	"phoenix/models"
	"phoenix/router"
)

func main() {

	//初始化gorm.db
	models.NewGormDB()

	r := router.App()
	r.Run(":8081")
}
