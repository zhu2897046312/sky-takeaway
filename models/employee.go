package models

import (
	"time"

	"github.com/sky-takeaway/utils"
	"gorm.io/gorm"
)

// COMMENT='员工信息';
type Employee struct {
	ID         uint      `gorm:"column:id;primaryKey" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	UserName   string    `gorm:"column:user_name;not null;unique" json:"user_name"`
	Password   string    `gorm:"column:password;not null" json:"-"`
	Phone      string    `gorm:"column:phone;not null" json:"phone"`
	Sex        string    `gorm:"column:sex;not null" json:"sex"`
	IDNumber   string    `gorm:"column:id_number;not null" json:"id_number"`
	Status     int       `gorm:"column:status;not null;default:1" json:"status"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateUser uint      `gorm:"column:create_user" json:"create_user"`
	UpdateUser uint      `gorm:"column:update_user" json:"update_user"`
}

func (e *Employee) TableName() string {
	return "employee"
}

func (u *Employee) Insert(employee *Employee) *gorm.DB {
	return utils.DB_MySQL.Model(&Employee{}).Create(employee)
}

func (u *Employee) Update(employee *Employee) *gorm.DB {
	return utils.DB_MySQL.Model(&Employee{}).Where("user_name =?", employee.UserName).Updates(&employee)
}

func (u *Employee) Delete(user_name string) *gorm.DB {
	return utils.DB_MySQL.Model(&Employee{}).Where("user_name =?", user_name).Delete(&Employee{})
}

func (u *Employee) FindByUserName(user_name string) (*Employee, *gorm.DB) {
	employee := Employee{}
	return &employee, utils.DB_MySQL.Model(&Employee{}).Where("user_name =?", user_name).Find(&employee)
}

func (u *Employee) PageQuery(page int, pageSize int) ([]Employee, *gorm.DB) {
	employees := make([]Employee, 0)
	var total int64

	utils.DB_MySQL.Model(&Employee{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&Employee{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}
