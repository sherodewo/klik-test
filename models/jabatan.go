package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Jabatan Struct
type Jabatan struct {
	ID          string    `gorm:"type:varchar(50);column:id;primary_key:true"`
	Name        string    `gorm:"type:varchar(50);column:name"`
	Description string    `gorm:"type:varchar(255);column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (c Jabatan) TableName() string {
	return "jabatan"
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Jabatan) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
