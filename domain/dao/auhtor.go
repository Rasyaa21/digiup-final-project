package dao

import "time"

type Author struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Fullname  string    `gorm:"type:varchar(50);not null" json:"fullname"`
	Gender    string    `gorm:"type:enum('m','f');default:null" json:"gender"`
	BirthDate time.Time `gorm:"type:datetime;default:null" json:"birth_date"`

	Book []Book `gorm:"foreignKey:AuthorID" json:"books"`
}
