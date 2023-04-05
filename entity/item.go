package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type (
	Item struct {
		ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
		Name            string    `gorm:"column:name"`
		Rating          int       `gorm:"column:rating"`
		Category        string    `gorm:"column:category"`
		Image           string    `gorm:"column:image"`
		Price           int       `gorm:"column:price"`
		Availibility    int       `gorm:"column:availibility"`
		ReputationValue int       `gorm:"column:reputation_value"`
		ReputationBadge string    `gorm:"column:reputation_badge"`
		CreatorID       uuid.UUID `json:"creator_id"`

		Creator User

		CreatedAt int64 `gorm:"column:created_at"`
		UpdatedAt int64 `gorm:"column:updated_at"`
	}

	RangeInput struct {
		Gte *int
		Lte *int
	}

	ItemQuery struct {
		Rating          *int        `json:"rating,omitempty"`
		ReputationBadge *string     `json:"reputationBadge,omitempty"`
		Category        *string     `json:"category,omitempty"`
		Availability    *RangeInput `json:"availability,omitempty"`
		CreatorID       []string    `json:"creator_id,omitempty"`
	}

	ItemInput struct {
		Name         *string `validate:"longer_10,word_alert" json:"name"`
		Rating       *int    `validate:"0_5" json:"rating"`
		Category     *string `validate:"category" json:"category"`
		Image        *string `validate:"url" json:"image"`
		Reputation   *int    `validate:"0_1000" json:"reputation"`
		Price        *int    `json:"price"`
		Availibility *int    `json:"availibility"`
	}
)

func (Item) TableName() string {
	return "items"
}

func (u *Item) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()

	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()

	return
}

func (item *Item) SetReputationBadge() {
	if item.ReputationValue <= 500 {
		item.ReputationBadge = "red"
	} else if item.ReputationValue <= 799 {
		item.ReputationBadge = "yellow"
	} else {
		item.ReputationBadge = "green"
	}
}
