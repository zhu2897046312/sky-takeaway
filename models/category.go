package models

import (
	"time"

	"github.com/sky-takeaway/utils"
	"gorm.io/gorm"
)

// COMMENT='菜品及套餐分类'
type Category struct {
	ID         uint      `gorm:"column:id;primaryKey" json:"id"`
	Type       int       `gorm:"column:type" json:"type"`                    //'类型   1 菜品分类 2 套餐分类'
	Name       string    `gorm:"column:name;not null;unique" json:"name"`    //套餐'分类名称'
	Sort       int       `gorm:"column:sort;not null;default:0" json:"sort"` //'顺序'
	Status     int       `gorm:"column:status" json:"status"`                //'分类状态 0:禁用，1:启用'
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`      //'创建时间'
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`      //'更新时间'
	CreateUser uint      `gorm:"column:create_user" json:"create_user"`      //'创建人'
	UpdateUser uint      `gorm:"column:update_user" json:"update_user"`      //'修改人'
}

func (c *Category) TableName() string {
	return "category"
}

func (u *Category) Insert(category *Category) *gorm.DB {
	return utils.DB_MySQL.Model(&Category{}).Create(category)
}

func (u *Category) Update(category *Category) *gorm.DB {
	return utils.DB_MySQL.Model(&Category{}).Where("name =?", category.Name).Updates(&category)
}

func (u *Category) Delete(name string) *gorm.DB {
	return utils.DB_MySQL.Model(&Category{}).Where("name =?", name).Delete(&Category{})
}

func (u *Category) FindByName(name string) (*Category, *gorm.DB) {
	category := Category{}
	return &category, utils.DB_MySQL.Model(&Category{}).Where("name =?", name).Find(&category)
}

func (u *Category) FindByID(id string) (*Category, *gorm.DB) {
	category := Category{}
	return &category, utils.DB_MySQL.Model(&Category{}).Where("id =?", id).Find(&category)
}

func (u *Category) PageQuery(page int, pageSize int) ([]Category, *gorm.DB) {
	category := make([]Category, 0)
	var total int64

	utils.DB_MySQL.Model(&Category{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&Category{}).Limit(pageSize).Offset(offset).Find(&category)
	return category, query
}
