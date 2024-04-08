package models

import(
	
	"time"
)
//COMMENT='菜品';
type Dish struct {
    ID          uint       `gorm:"column:id;primaryKey"`
    Name        string     `gorm:"column:name;not null;unique"`
    CategoryID  uint       `gorm:"column:category_id"`
    Price       float64    `gorm:"column:price"`
    Image       string     `gorm:"column:image"`
    Description string     `gorm:"column:description"`
    Status      int        `gorm:"column:status;default:1"`
    CreateTime  time.Time  `gorm:"column:create_time"`
    UpdateTime  time.Time  `gorm:"column:update_time"`
    CreateUser  uint       `gorm:"column:create_user"`
    UpdateUser  uint       `gorm:"column:update_user"`
}

func (d *Dish) TableName() string {
    return "dish"
}