package models

import (
	"time"

	"github.com/sky-takeaway/utils"
	"gorm.io/gorm"
)

// COMMENT='套餐';
type SetMeal struct {
	ID          uint      `gorm:"column:id;primaryKey"`
	CategoryID  uint      `gorm:"column:category_id;not null"`
	Name        string    `gorm:"column:name;not null"`
	Price       float64   `gorm:"column:price;not null"`
	Status      int       `gorm:"column:status;default:1"`
	Description string    `gorm:"column:description"`
	Image       string    `gorm:"column:image"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
	CreateUser  uint      `gorm:"column:create_user"`
	UpdateUser  uint      `gorm:"column:update_user"`
}

func (s *SetMeal) TableName() string {
	return "setmeal"
}

func (u *SetMeal) Insert(category *SetMeal) *gorm.DB {
	return utils.DB_MySQL.Model(&SetMeal{}).Create(category)
}

func (u *SetMeal) Update(category *SetMeal) *gorm.DB {
	return utils.DB_MySQL.Model(&SetMeal{}).Where("name =?", category.Name).Updates(&category)
}

func (u *SetMeal) Delete(name string) *gorm.DB {
	return utils.DB_MySQL.Model(&SetMeal{}).Where("name =?", name).Delete(&SetMeal{})
}

func (u *SetMeal) FindByName(name string) (*SetMeal, *gorm.DB) {
	category := SetMeal{}
	return &category, utils.DB_MySQL.Model(&SetMeal{}).Where("name =?", name).Find(&category)
}

func (u *SetMeal) FindByCategoryID(CategoryID string) (*SetMeal, *gorm.DB) {
	category := SetMeal{}
	return &category, utils.DB_MySQL.Model(&SetMeal{}).Where("category_id = ?", CategoryID).Find(&category)
}

func (u *SetMeal) PageQuery(page int, pageSize int) ([]SetMeal, *gorm.DB) {
	category := make([]SetMeal, 0)
	var total int64

	utils.DB_MySQL.Model(&SetMeal{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&SetMeal{}).Limit(pageSize).Offset(offset).Find(&category)
	return category, query
}
