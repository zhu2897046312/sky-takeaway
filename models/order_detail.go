package models

import (
	"github.com/sky-takeaway/utils"
	"gorm.io/gorm"
)

// COMMENT='订单明细表';
type OrderDetail struct {
	ID         uint    `gorm:"column:id;primaryKey"`
	Name       string  `gorm:"column:name"`
	Image      string  `gorm:"column:image"`
	OrderID    uint    `gorm:"column:order_id;not null"`
	DishID     uint    `gorm:"column:dish_id"`
	SetmealID  uint    `gorm:"column:setmeal_id"`
	DishFlavor string  `gorm:"column:dish_flavor"`
	Number     int     `gorm:"column:number;not null;default:1"`
	Amount     float64 `gorm:"column:amount;not null"`
}

func (o *OrderDetail) TableName() string {
	return "order_detail"
}

func (u *OrderDetail) Insert(category *OrderDetail) *gorm.DB {
	return utils.DB_MySQL.Model(&OrderDetail{}).Create(category)
}

func (u *OrderDetail) Update(category *OrderDetail) *gorm.DB {
	return utils.DB_MySQL.Model(&OrderDetail{}).Where("name =?", category.Name).Updates(&category)
}

func (u *OrderDetail) Delete(name string) *gorm.DB {
	return utils.DB_MySQL.Model(&OrderDetail{}).Where("name =?", name).Delete(&OrderDetail{})
}

func (u *OrderDetail) FindByName(name string) (*OrderDetail, *gorm.DB) {
	category := OrderDetail{}
	return &category, utils.DB_MySQL.Model(&OrderDetail{}).Where("name =?", name).Find(&category)
}

func (u *OrderDetail) FindByOrderID(OrderID string) (*OrderDetail, *gorm.DB) {
	category := OrderDetail{}
	return &category, utils.DB_MySQL.Model(&OrderDetail{}).Where("order_id =?", OrderID).Find(&category)
}

func (u *OrderDetail) FindByDishID(DishID string) (*OrderDetail, *gorm.DB) {
	category := OrderDetail{}
	return &category, utils.DB_MySQL.Model(&OrderDetail{}).Where("dish_id =?", DishID).Find(&category)
}

func (u *OrderDetail) FindBySetmealID(SetmealID string) (*OrderDetail, *gorm.DB) {
	category := OrderDetail{}
	return &category, utils.DB_MySQL.Model(&OrderDetail{}).Where("setmeal_id =?", SetmealID).Find(&category)
}

func (u *OrderDetail) PageQuery(page int, pageSize int) ([]OrderDetail, *gorm.DB) {
	category := make([]OrderDetail, 0)
	var total int64

	utils.DB_MySQL.Model(&OrderDetail{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&OrderDetail{}).Limit(pageSize).Offset(offset).Find(&category)
	return category, query
}
