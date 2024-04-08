package models
import(
	
	"time"
)

//COMMENT='套餐';
type SetMeal struct {
    ID           uint       `gorm:"column:id;primaryKey"`
    CategoryID   uint       `gorm:"column:category_id;not null"`
    Name         string     `gorm:"column:name;not null"`
    Price        float64    `gorm:"column:price;not null"`
    Status       int        `gorm:"column:status;default:1"`
    Description  string     `gorm:"column:description"`
    Image        string     `gorm:"column:image"`
    CreateTime   time.Time  `gorm:"column:create_time"`
    UpdateTime   time.Time  `gorm:"column:update_time"`
    CreateUser   uint       `gorm:"column:create_user"`
    UpdateUser   uint       `gorm:"column:update_user"`
}

func (s *SetMeal) TableName() string {
    return "setmeal"
}