package service

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sky-takeaway/models"
)
//添加菜品
func AddCategory(c *gin.Context){
	category := models.Category{}
    //TODO: 接收json数据
    if err := c.ShouldBindJSON(&category); err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
        return
    }
	//TODO: 是否已存在
	sql_category, db := category.FindByName(category.Name)
	if db.Error != nil {
		log.Println(db.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad request",
			"data":    nil,
		})
		return
	}
	if sql_category.Name != "" {
		log.Println(sql_category)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "已存在",
			"data":    sql_category,
		})
		return
	}
	//TODO: 插入到数据库中
	if err := category.Insert(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "注册菜品失败",
			"data":    nil,
		})
		return
	}
    //TODO: 成功
    c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "添加成功",
        "data":    nil,
    })
}

// 分页查询菜品
func PageQueryCategory(c *gin.Context){
	//接收数据
	page := c.Query("page")
	pageSize := c.Query("page_size")

	page_, _ := strconv.Atoi(page)
	pageSize_, _ :=strconv.Atoi(pageSize)

	// 分页查询
	employee := models.Category{}
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
//修改菜品数据
func UpdateCategory(c *gin.Context){
	category := models.Category{}
	//接收数据
	if err := c.ShouldBindJSON(&category); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	//判断是否存在 如：存在，修改 否则 返回
	_, db := category.FindByName(category.Name)
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
	if err := category.Update(&category).Error; err != nil {
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

//删除菜品
func DeleteCategory(c *gin.Context){
	category := models.Category{}
    //TODO: 接收json数据
    if err := c.ShouldBindJSON(&category); err != nil {
        log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
        return
    }
	//TODO: 是否已存在
	_, db := category.FindByName(category.Name)
	if db.Error != nil {
		log.Println(db.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "不存在该菜品",
			"data":    nil,
		})
		return
	}
	//TODO: 从数据库中删除
	if err := category.Delete(category.Name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "删除菜品失败",
			"data":    nil,
		})
		return
	}
    //TODO: 成功
    c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "删除成功",
        "data":    nil,
    })
}





