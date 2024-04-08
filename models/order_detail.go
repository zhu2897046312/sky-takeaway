package models

import(
	
)
//COMMENT='订单明细表';
type OrderDetail struct {
    ID          uint    `gorm:"column:id;primaryKey"`
    Name        string  `gorm:"column:name"`
    Image       string  `gorm:"column:image"`
    OrderID     uint    `gorm:"column:order_id;not null"`
    DishID      uint    `gorm:"column:dish_id"`
    SetmealID   uint    `gorm:"column:setmeal_id"`
    DishFlavor  string  `gorm:"column:dish_flavor"`
    Number      int     `gorm:"column:number;not null;default:1"`
    Amount      float64 `gorm:"column:amount;not null"`
}

func (o *OrderDetail) TableName() string {
    return "order_detail"
}