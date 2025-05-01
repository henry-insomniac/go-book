package model

import (
	"time"

	"github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Article struct {
	ID          string    `gorm:"primaryKey;type:varchar(32)" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	ContentMD   string    `gorm:"type:longtext;not null" json:"content_md"`
	Summary     string    `gorm:"type:text" json:"summary"`
	CoverImage  string    `gorm:"type:varchar(512)" json:"cover_image"`
	CoverThumb  string    `gorm:"type:varchar(512)" json:"cover_thumb"`
	ReadTime    int       `gorm:"default:0" json:"read_time"`
	PublishedAt time.Time `gorm:"not null" json:"published_at"`
	AuthorID    string    `gorm:"type:varchar(32);not null;index" json:"author_id"`
	Author      User      `gorm:"foreignKey:AuthorID" json:"author"`
	Tags        []Tag     `gorm:"many2many:article_tags;" json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == "" {
		a.ID, err = gonanoid.New()
	}
	return
}
