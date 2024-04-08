package models

import(
	"time"
)

//COMMENT='菜品及套餐分类'
type Category struct {
    ID          uint      `gorm:"column:id;primaryKey"`
    Type        int       `gorm:"column:type"`
    Name        string    `gorm:"column:name;not null;unique"`
    Sort        int       `gorm:"column:sort;not null;default:0"`
    Status      int       `gorm:"column:status"`
    CreateTime  time.Time `gorm:"column:create_time"`
    UpdateTime  time.Time `gorm:"column:update_time"`
    CreateUser  uint      `gorm:"column:create_user"`
    UpdateUser  uint      `gorm:"column:update_user"`
}

func (c *Category) TableName() string {
    return "category"
}