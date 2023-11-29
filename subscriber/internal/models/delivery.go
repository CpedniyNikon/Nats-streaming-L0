package models

type Delivery struct {
	ID      uint   `gorm:"primarykey" json:"-"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
	OrderId uint   `gorm:"foreignKey:OrderId" json:"-"`
}
