package models

import(
	
	"time"
)

//COMMENT='购物车';
type ShoppingCart struct {
    ID          uint       `gorm:"column:id;primaryKey"`
    Name        string     `gorm:"column:name"`
    Image       string     `gorm:"column:image"`
    UserID      uint       `gorm:"column:user_id"`
    DishID      uint       `gorm:"column:dish_id"`
    SetMealID   uint       `gorm:"column:setmeal_id"`
    DishFlavor  string     `gorm:"column:dish_flavor"`
    Number      int        `gorm:"column:number;default:1"`
    Amount      float64    `gorm:"column:amount"`
    CreateTime  time.Time  `gorm:"column:create_time"`
}

func (s *ShoppingCart) TableName() string {
    return "shopping_cart"
}