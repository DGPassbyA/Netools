package model

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model
	Name    string `json:"name"`
	Content string `json:"content"`
	Type    string `json:"type"`
}
