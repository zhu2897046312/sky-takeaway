package models
import(
	
	"time"
)

//COMMENT='订单表';
type Order struct {
    ID                     uint       `gorm:"column:id;primaryKey"`
    Number                 string     `gorm:"column:number"`
    Status                 int        `gorm:"column:status;not null;default:1"`
    UserID                 uint       `gorm:"column:user_id;not null"`
    AddressBookID          uint       `gorm:"column:address_book_id;not null"`
    OrderTime              time.Time  `gorm:"column:order_time;not null"`
    CheckoutTime           time.Time `gorm:"column:checkout_time"`
    PayMethod              int        `gorm:"column:pay_method;not null;default:1"`
    PayStatus              int8       `gorm:"column:pay_status;not null;default:0"`
    Amount                 float64    `gorm:"column:amount;not null"`
    Remark                 string     `gorm:"column:remark"`
    Phone                  string     `gorm:"column:phone"`
    Address                string     `gorm:"column:address"`
    UserName               string     `gorm:"column:user_name"`
    Consignee              string     `gorm:"column:consignee"`
    CancelReason           string     `gorm:"column:cancel_reason"`
    RejectionReason        string     `gorm:"column:rejection_reason"`
    CancelTime             time.Time `gorm:"column:cancel_time"`
    EstimatedDeliveryTime  time.Time `gorm:"column:estimated_delivery_time"`
    DeliveryStatus         int8       `gorm:"column:delivery_status;not null;default:1"`
    DeliveryTime           time.Time `gorm:"column:delivery_time"`
    PackAmount             int        `gorm:"column:pack_amount"`
    TablewareNumber        int        `gorm:"column:tableware_number"`
    TablewareStatus        int8       `gorm:"column:tableware_status;not null;default:1"`
}

func (o *Order) TableName() string {
    return "orders"
}