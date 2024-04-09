package models

import (
	"time"

	"github.com/sky-takeaway/utils"
	"gorm.io/gorm"
)

// COMMENT='菜品';
type Dish struct {
	ID          uint      `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;not null;unique"`
	CategoryID  uint      `gorm:"column:category_id"` //逻辑外键
	Price       float64   `gorm:"column:price"`
	Image       string    `gorm:"column:image"`
	Description string    `gorm:"column:description"`
	Status      int       `gorm:"column:status;default:1"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
	CreateUser  uint      `gorm:"column:create_user"`
	UpdateUser  uint      `gorm:"column:update_user"`
}

func (d *Dish) TableName() string {
	return "dish"
}

func (u *Dish) Insert(category *Dish) *gorm.DB {
	return utils.DB_MySQL.Model(&Dish{}).Create(category)
}

func (u *Dish) Update(category *Dish) *gorm.DB {
	return utils.DB_MySQL.Model(&Dish{}).Where("name =?", category.Name).Updates(&category)
}

func (u *Dish) Delete(name string) *gorm.DB {
	return utils.DB_MySQL.Model(&Dish{}).Where("name =?", name).Delete(&Dish{})
}

func (u *Dish) FindByName(name string) (*Dish, *gorm.DB) {
	category := Dish{}
	return &category, utils.DB_MySQL.Model(&Dish{}).Where("name =?", name).Find(&category)
}

func (u *Dish) FindByID(id string) (*Dish, *gorm.DB) {
	category := Dish{}
	return &category, utils.DB_MySQL.Model(&Dish{}).Where("id =?", id).Find(&category)
}

func (u *Dish) FindByCategoryID(CategoryID string) (*Dish, *gorm.DB) {
	category := Dish{}
	return &category, utils.DB_MySQL.Model(&Dish{}).Where("category_id =?", CategoryID).Find(&category)
}

func (u *Dish) PageQuery(page int, pageSize int) ([]Dish, *gorm.DB) {
	category := make([]Dish, 0)
	var total int64

	utils.DB_MySQL.Model(&Dish{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&Dish{}).Limit(pageSize).Offset(offset).Find(&category)
	return category, query
}
