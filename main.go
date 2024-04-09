package main

import (
	"fmt"

	"github.com/sky-takeaway/models"
	"github.com/sky-takeaway/routers"
	"github.com/sky-takeaway/utils"
)

func main() {
	fmt.Println("Hello, World!")

	utils.InitConfig("")
	utils.InitMySQL()
	utils.InitRedis()
	exists := utils.DB_MySQL.Migrator().HasTable(&models.Category{})
    if !exists {
        // 表不存在，创建表
        utils.DB_MySQL.AutoMigrate(&models.Category{})
        println("Table created")
    } else {
        println("Table already exists")
    }
	r := routers.Router()
	r.Run()
}
