package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`

	BuyerID uuid.UUID `gorm:"column:buyer_id"`
	ItemID  uuid.UUID `gorm:"column:item_id"`

	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (u *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()

	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()

	return
}
