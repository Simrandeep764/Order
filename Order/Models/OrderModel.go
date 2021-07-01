package Models

type Order struct {
	OrderId         uint   `gorm:"primary_key;auto_increment:true" json:"oid"`
	PaymentId       uint   `json:"pid"`
	CustomerId      uint   `json:"custid"`
	ShipmentDetails string `json:"sid"`
}

func (o *Order) TableName() string {
	return "order"
}
