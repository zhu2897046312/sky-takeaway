package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sky-takeaway/models"
)
func Login(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(user)
	c.JSON(200, gin.H{
        "code": 200,
        "msg": "success",
        "data": user,
    })
}

func Register(c *gin.Context){
	var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Println(user)
	user.CreateTime = time.Now()
	err := user.Insert(&user).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "failed to insert user",
			"data": user,
		})
		return
	}
    c.JSON(200, gin.H{
        "code": 200,
        "msg": "success",
        "data": nil,
    })
}