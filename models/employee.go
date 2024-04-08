package models

import(
	"time"
)
//COMMENT='员工信息';
type Employee struct {
    ID          uint       `gorm:"column:id;primaryKey"`
    Name        string     `gorm:"column:name;not null"`
    Username    string     `gorm:"column:username;not null;unique"`
    Password    string     `gorm:"column:password;not null"`
    Phone       string     `gorm:"column:phone;not null"`
    Sex         string     `gorm:"column:sex;not null"`
    IDNumber    string     `gorm:"column:id_number;not null"`
    Status      int        `gorm:"column:status;not null;default:1"`
    CreateTime  time.Time  `gorm:"column:create_time"`
    UpdateTime  time.Time  `gorm:"column:update_time"`
    CreateUser  uint       `gorm:"column:create_user"`
    UpdateUser  uint       `gorm:"column:update_user"`
}

func (e *Employee) TableName() string {
    return "employee"
}