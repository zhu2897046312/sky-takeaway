package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sky-takeaway/models"
	"github.com/sky-takeaway/utils"
)

func Login(c *gin.Context) {
	employee := models.Employee{}
	//接收json数据
	if err := c.ShouldBindJSON(&employee); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	//查询数据库
	sql_employee, db := employee.FindByUserName(employee.UserName)
	if db.Error != nil {
		log.Println(db.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": db.Error.Error(),
			"data":    nil,
		})
		return
	}
	//md5加密
	password, err := utils.Md5(employee.Password)
	if err!= nil {
		log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
        return
	}

	//密码比对
	if password != sql_employee.Password{
		log.Println("密码错误")
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": "密码错误",
            "data":    nil,
        })
        return
	}
	//生成token
	token , err := utils.GenerateToken(employee.UserName)
	if err != nil {
		log.Println("token生成失败")
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": "token生成失败",
            "data":    nil,
        })
        return
	}
	//TODO: 存入redis

	//TODO: 返回token
	c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "登录成功",
        "data":    token,
    })
}

func Register(c *gin.Context) {
	employee := models.Employee{}
	//接收json数据
	if err := c.ShouldBindJSON(&employee); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	//TODO: 是否已经存在该用户
	sql_employee, db := employee.FindByUserName(employee.UserName)
	if db.Error == nil {
		log.Println(db.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户已注册",
			"data":    sql_employee,
		})
		return
	}

	//md5加密
	password, err := utils.Md5(employee.Password)
	if err!= nil {
		log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
        return
	}
	employee.Password = password
	//插入数据库
	if employee.Insert(&employee).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "注册失败",
			"data":    nil,
		})
		return
	}

	//返回成功
	c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "注册成功",
        "data":    nil,
    })
}

func PageQuery(c *gin.Context){
	//接收数据
	page := c.Query("page")
	pageSize := c.Query("page_size")

	page_, _ := strconv.Atoi(page)
	pageSize_, _ :=strconv.Atoi(pageSize)

	// 分页查询
	employee := models.Employee{}
	data, db := employee.PageQuery(page_,pageSize_)
	if db.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "分页查找失败",
			"data":    nil,
		})
		return
	}
	//返回数据
	c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "分页查找成功",
        "data":    data,
    })
}

func UpdateEmployee(c *gin.Context){
	employee := models.Employee{}
	//接收数据
	if err := c.ShouldBindJSON(&employee); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	//判断是否存在 如：存在，修改 否则 返回
	_, db := employee.FindByUserName(employee.UserName)
	if db.Error != nil {
		log.Println(db.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "不存在该条数据",
			"data":    nil,
		})
		return
	}

	//修改
	if err := employee.Update(&employee).Error; err != nil {
		log.Println(db.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "修改失败",
			"data":    nil,
		})
		return 
	}

	//返回数据
	c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "修改成功",
        "data":    nil,
    })
}
