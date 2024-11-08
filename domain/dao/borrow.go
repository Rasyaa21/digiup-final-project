package dao

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	BorrowDate   time.Time  `gorm:"type:datetime;not null" json:"borrow_date"`
	ReturnDate   *time.Time `gorm:"type:datetime" json:"return_date"`
	BookID       uint64     `gorm:"not null" json:"book_id"`
	PersonID     uint64     `gorm:"not null" json:"person_id"`
	BorrowBookID Book       `gorm:"foreignKey:BookID"`
	UserID       Person     `gorm:"foreignKey:PersonID"`
}
