package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/sky-takeaway/utils"
)


// COMMENT='用户信息';
type User struct {
	OpenID     string    `gorm:"column:openid" json:"openid"`
	Name       string    `gorm:"column:name" json:"name"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	Sex        string    `gorm:"column:sex" json:"sex"`
	IDNumber   string    `gorm:"column:id_number" json:"id_number"`
	Avatar     string    `gorm:"column:avatar" json:"avatar"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

func (u *User) TableName() string {
	return "user"
}

func (u*User)Insert(user *User) *gorm.DB{
	return utils.DB_MySQL.Model(&User{}).Create(user)
}

func (u *User)Update(user User) *gorm.DB{
    return utils.DB_MySQL.Model(&User{}).Where("open_id =?", user.OpenID).Updates(&user)
}

func (u *User)Delete(open_id string) *gorm.DB{
    return utils.DB_MySQL.Model(&User{}).Where("open_id =?", open_id).Delete(&User{})
}
