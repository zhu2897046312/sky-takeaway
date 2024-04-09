package models

import (
	"github.com/sky-takeaway/utils"
	"gorm.io/gorm"
)

// COMMENT='菜品口味关系表';
type DishFlavor struct {
	ID     uint   `gorm:"column:id;primaryKey"`
	DishID uint   `gorm:"column:dish_id"` //逻辑外键
	Name   string `gorm:"column:name"`
	Value  string `gorm:"column:value"`
}

func (d *DishFlavor) TableName() string {
	return "dish_flavor"
}

func (u *DishFlavor) Insert(category *DishFlavor) *gorm.DB {
	return utils.DB_MySQL.Model(&DishFlavor{}).Create(category)
}

func (u *DishFlavor) Update(category *DishFlavor) *gorm.DB {
	return utils.DB_MySQL.Model(&DishFlavor{}).Where("name =?", category.Name).Updates(&category)
}

func (u *DishFlavor) Delete(name string) *gorm.DB {
	return utils.DB_MySQL.Model(&DishFlavor{}).Where("name =?", name).Delete(&DishFlavor{})
}

func (u *DishFlavor) FindByName(name string) (*Dish, *gorm.DB) {
	category := Dish{}
	return &category, utils.DB_MySQL.Model(&DishFlavor{}).Where("name =?", name).Find(&category)
}

func (u *DishFlavor) PageQuery(page int, pageSize int) ([]DishFlavor, *gorm.DB) {
	category := make([]DishFlavor, 0)
	var total int64

	utils.DB_MySQL.Model(&Dish{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&DishFlavor{}).Limit(pageSize).Offset(offset).Find(&category)
	return category, query
}
