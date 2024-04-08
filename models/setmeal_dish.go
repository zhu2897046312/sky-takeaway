package models

import(
	
)
//COMMENT='套餐菜品关系';

type SetMealDish struct {
    ID        uint   `gorm:"column:id;primaryKey"`
    SetMealID uint   `gorm:"column:setmeal_id"`
    DishID    uint   `gorm:"column:dish_id"`
    Name      string `gorm:"column:name"`
    Price     float64 `gorm:"column:price"`
    Copies    int    `gorm:"column:copies"`
}

func (s *SetMealDish) TableName() string {
    return "setmeal_dish"
}