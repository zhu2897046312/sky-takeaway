package models

import (
	"github.com/sky-takeaway/utils"
	"gorm.io/gorm"
)

//COMMENT='套餐菜品关系';

type SetMealDish struct {
	ID        uint    `gorm:"column:id;primaryKey"`
	SetMealID uint    `gorm:"column:setmeal_id"`
	DishID    uint    `gorm:"column:dish_id"`
	Name      string  `gorm:"column:name"`
	Price     float64 `gorm:"column:price"`
	Copies    int     `gorm:"column:copies"`
}

func (s *SetMealDish) TableName() string {
	return "setmeal_dish"
}

func (u *SetMealDish) Insert(category *SetMealDish) *gorm.DB {
	return utils.DB_MySQL.Model(&SetMealDish{}).Create(category)
}

func (u *SetMealDish) Update(category *SetMealDish) *gorm.DB {
	return utils.DB_MySQL.Model(&SetMealDish{}).Where("name =?", category.Name).Updates(&category)
}

func (u *SetMealDish) Delete(name string) *gorm.DB {
	return utils.DB_MySQL.Model(&SetMealDish{}).Where("name =?", name).Delete(&SetMealDish{})
}

func (u *SetMealDish) FindByName(name string) (*SetMealDish, *gorm.DB) {
	category := SetMealDish{}
	return &category, utils.DB_MySQL.Model(&SetMealDish{}).Where("name =?", name).Find(&category)
}

func (u *SetMealDish) FindBySetMealID(SetMealID string) (*SetMealDish, *gorm.DB) {
	category := SetMealDish{}
	return &category, utils.DB_MySQL.Model(&SetMealDish{}).Where("setMeal_id =?", SetMealID).Find(&category)
}

func (u *SetMealDish) FindByDishID(DishID string) (*SetMealDish, *gorm.DB) {
	category := SetMealDish{}
	return &category, utils.DB_MySQL.Model(&SetMealDish{}).Where("dish_id =?", DishID).Find(&category)
}

func (u *SetMealDish) PageQuery(page int, pageSize int) ([]SetMealDish, *gorm.DB) {
	category := make([]SetMealDish, 0)
	var total int64

	utils.DB_MySQL.Model(&SetMealDish{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&SetMealDish{}).Limit(pageSize).Offset(offset).Find(&category)
	return category, query
}
