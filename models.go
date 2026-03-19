package skeleton

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id"`
	Active      bool       `json:"active" db:"active"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (Item) TableName() string {
	return "skeleton_items"
}
