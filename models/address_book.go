package models

//COMMENT='地址簿'
type AddressBook struct {
	ID           uint   `gorm:"column:id;primaryKey" json:"id"`
	UserID       uint   `gorm:"column:user_id;not null"`
	Consignee    string `gorm:"column:consignee"`
	Sex          string `gorm:"column:sex"`
	Phone        string `gorm:"column:phone;not null"`
	ProvinceCode string `gorm:"column:province_code"`
	ProvinceName string `gorm:"column:province_name"`
	CityCode     string `gorm:"column:city_code"`
	CityName     string `gorm:"column:city_name"`
	DistrictCode string `gorm:"column:district_code"`
	DistrictName string `gorm:"column:district_name"`
	Detail       string `gorm:"column:detail"`
	Label        string `gorm:"column:label"`
	IsDefault    bool   `gorm:"column:is_default;default:false"`
}

func (a *AddressBook) TableName() string {
	return "address_book"
}
