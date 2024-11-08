package dto

import (
	"base-gin/domain/dao"
	"errors"
	"time"
)

type AuthorReq struct {
	Fullname     string    `json:"fullname" binding:"required,min=2,max=50"`
	Gender       string    `json:"gender" binding:"required,oneof=m f"`
	BirthDateStr string    `json:"birth_date" binding:"required,datetime=2006-01-02"`
	BirthDate    time.Time `json:"-"`
}

type AuthorRes struct {
	Fullname  string    `json:"fullname"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
}

type AuthorUpdateReq struct {
	ID       uint   `json:"id" binding:"required"`
	Fullname string `json:"fullname" binding:"omitempty,min=2,max=50"`
}

func (req *AuthorReq) ParseBirthDate() error {
	layout := "2006-01-02"
	parseData, err := time.Parse(layout, req.BirthDateStr)
	if err != nil {
		return errors.New(err.Error())
	}
	req.BirthDate = parseData
	return nil
}

func (req *AuthorReq) ToEntity() dao.Author {
	var item dao.Author
	item.Fullname = req.Fullname
	item.Gender = req.Gender
	item.BirthDate = req.BirthDate
	return item
}

func (req *AuthorRes) FromEntity(item *dao.Author) {
	req.Fullname = item.Fullname
	req.Gender = item.Gender
	req.BirthDate = item.BirthDate
}

func (req *AuthorUpdateReq) ToEntity() dao.Author {
	var item dao.Author
	item.Fullname = req.Fullname
	return item
}
