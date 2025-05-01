package model

import "time"

type ArticleTag struct {
	ArticleID string `gorm:"primaryKey;type:varchar(32)"`
	TagID     string `gorm:"primaryKey;type:varchar(32)"`
	CreatedAt time.Time
}
