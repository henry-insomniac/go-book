package model

import (
	"time"

	"github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Tag struct {
	ID        string    `gorm:"primaryKey;type:varchar(32)" json:"id"`
	Name      string    `gorm:"type:varchar(100);unique;not null" json:"name"`
	Articles  []Article `gorm:"many2many:article_tags;" json:"articles,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == "" {
		t.ID, err = gonanoid.New()
	}
	return
}
