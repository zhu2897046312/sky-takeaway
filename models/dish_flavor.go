package models

import(
	
)
//COMMENT='菜品口味关系表';
type DishFlavor struct {
    ID     uint   `gorm:"column:id;primaryKey"`
    DishID uint   `gorm:"column:dish_id"`
    Name   string `gorm:"column:name"`
    Value  string `gorm:"column:value"`
}

func (d *DishFlavor) TableName() string {
    return "dish_flavor"
}