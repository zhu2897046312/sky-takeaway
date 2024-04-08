package main

import (
	"fmt"
	"github.com/sky-takeaway/utils"
	"github.com/sky-takeaway/routers"
)

func main() {
	fmt.Println("Hello, World!")

	utils.InitConfig("")
	utils.InitMySQL()
	utils.InitRedis()
	r := routers.Router()
	r.Run()
}
