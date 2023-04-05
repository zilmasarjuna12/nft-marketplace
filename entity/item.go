package entity

import (
	"github.com/google/uuid"
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
)

func (Item) TableName() string {
	return "items"
}
