package dao

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string    `gorm:"size:56;not null;"`
	Subtitle      *string   `gorm:"size:64;"`
	PublisherID   uint      `gorm:"not null;"`
	BookPublisher Publisher `gorm:"foreignKey:PublisherID;"`
	AuthorID      uint      `gorm:"not null;"`
	BookAuthor    Author    `gorm:"foreignKey:AuthorID;"`
}
