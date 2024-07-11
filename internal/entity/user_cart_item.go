package entity

type UserCartItem struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    string
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID uint64
	Quantity  int
}
