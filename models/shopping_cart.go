package models

import (
	"time"

	"github.com/sky-takeaway/utils"
	"gorm.io/gorm"
)

// COMMENT='购物车';
type ShoppingCart struct {
	ID         uint      `gorm:"column:id;primaryKey"`
	Name       string    `gorm:"column:name"`
	Image      string    `gorm:"column:image"`
	UserID     uint      `gorm:"column:user_id"`
	DishID     uint      `gorm:"column:dish_id"`
	SetMealID  uint      `gorm:"column:setmeal_id"`
	DishFlavor string    `gorm:"column:dish_flavor"`
	Number     int       `gorm:"column:number;default:1"`
	Amount     float64   `gorm:"column:amount"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (s *ShoppingCart) TableName() string {
	return "shopping_cart"
}

func (u *ShoppingCart) Insert(category *ShoppingCart) *gorm.DB {
	return utils.DB_MySQL.Model(&ShoppingCart{}).Create(category)
}

func (u *ShoppingCart) Update(category *ShoppingCart) *gorm.DB {
	return utils.DB_MySQL.Model(&ShoppingCart{}).Where("name =?", category.Name).Updates(&category)
}

func (u *ShoppingCart) Delete(name string) *gorm.DB {
	return utils.DB_MySQL.Model(&ShoppingCart{}).Where("name =?", name).Delete(&ShoppingCart{})
}

func (u *ShoppingCart) FindByName(name string) (*ShoppingCart, *gorm.DB) {
	category := ShoppingCart{}
	return &category, utils.DB_MySQL.Model(&ShoppingCart{}).Where("name =?", name).Find(&category)
}

func (u *ShoppingCart) FindBySetMealID(SetMealID string) (*ShoppingCart, *gorm.DB) {
	category := ShoppingCart{}
	return &category, utils.DB_MySQL.Model(&ShoppingCart{}).Where("setMeal_id =?", SetMealID).Find(&category)
}

func (u *ShoppingCart) FindByDishID(DishID string) (*ShoppingCart, *gorm.DB) {
	category := ShoppingCart{}
	return &category, utils.DB_MySQL.Model(&ShoppingCart{}).Where("dish_id =?", DishID).Find(&category)
}

func (u *ShoppingCart) FindByUserID(UserID string) (*ShoppingCart, *gorm.DB) {
	category := ShoppingCart{}
	return &category, utils.DB_MySQL.Model(&ShoppingCart{}).Where("user_id =?", UserID).Find(&category)
}

func (u *ShoppingCart) PageQuery(page int, pageSize int) ([]ShoppingCart, *gorm.DB) {
	category := make([]ShoppingCart, 0)
	var total int64

	utils.DB_MySQL.Model(&ShoppingCart{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&ShoppingCart{}).Limit(pageSize).Offset(offset).Find(&category)
	return category, query
}
