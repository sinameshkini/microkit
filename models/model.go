package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelTime struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ModelIID basic model with auto incremental id
type ModelIID struct {
	ID IID `gorm:"primarykey" json:"id"`
	ModelTime
}

// ModelSID basic model with snowflake id
type ModelSID struct {
	ID SID `gorm:"primarykey" json:"id"`
	ModelTime
}

// ModelUUID basic model with UUID
type ModelUUID struct {
	ID uuid.UUID `gorm:"primarykey" json:"id"`
	ModelTime
}

// Model basic model with string id
type Model struct {
	ID string `gorm:"primarykey" json:"id"`
	ModelTime
}
