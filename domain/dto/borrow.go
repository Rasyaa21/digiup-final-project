package dto

import (
	"base-gin/domain/dao"
	"fmt"
	"time"
)

type BorrowReq struct {
	BorrowDate *time.Time `json:"borrow_date"`
	ReturnDate string     `json:"return_date" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	BookID     uint       `json:"book_id" binding:"required"`
	PersonID   uint       `json:"person_id" binding:"required"`
}

type BorrowRes struct {
	ID         uint       `json:"id"`
	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"`
	BookID     uint       `json:"book_id"`
	PersonID   uint       `json:"person_id"`
}

type UpdateBorrowReq struct {
	BorrowDate *time.Time `json:"borrow_date,omitempty"`
	ReturnDate string     `json:"return_date,omitempty" binding:"omitempty,datetime=2006-01-02 15:04:05"`
	BookID     *uint      `json:"book_id,omitempty"`
	PersonID   *uint      `json:"person_id,omitempty"`
}

func (req *UpdateBorrowReq) ToEntity() dao.Borrow {
	var item dao.Borrow
	item.BorrowDate = *req.BorrowDate
	item.BookID = uint64(*req.BookID)
	item.PersonID = uint64(*req.PersonID)
	return item
}

func (res *BorrowRes) FromEntity(item *dao.Borrow) {
	res.ID = uint(item.ID)
	res.BorrowDate = item.BorrowDate
	if item.ReturnDate != nil {
		res.ReturnDate = item.ReturnDate
	} else {
		res.ReturnDate = nil
	}
	res.BookID = uint(item.BookID)
	res.PersonID = uint(item.PersonID)
}

func (req *BorrowReq) ToEntity() dao.Borrow {
	var item dao.Borrow
	item.BorrowDate = *req.BorrowDate
	if req.ReturnDate != "" {
		parsedReturnDate, err := time.Parse("2006-01-02 15:04:05", req.ReturnDate)
		if err == nil {
			item.ReturnDate = &parsedReturnDate
		}
	}
	item.BookID = uint64(req.BookID)
	item.PersonID = uint64(req.PersonID)
	return item
}

func NewBorrowReq(returnDate string, bookID, personID uint) BorrowReq {
	now := time.Now()
	return BorrowReq{
		BorrowDate: &now,
		ReturnDate: returnDate,
		BookID:     bookID,
		PersonID:   personID,
	}
}

func ConvertToBorrowRes(req BorrowReq) (BorrowRes, error) {
	var returnDate *time.Time
	if req.ReturnDate != "" {
		parsedReturnDate, err := time.Parse("2006-01-02 15:04:05", req.ReturnDate)
		if err != nil {
			return BorrowRes{}, fmt.Errorf("invalid return date format: %v", err)
		}
		returnDate = &parsedReturnDate
	}
	return BorrowRes{
		BorrowDate: *req.BorrowDate,
		ReturnDate: returnDate,
		BookID:     req.BookID,
		PersonID:   req.PersonID,
	}, nil
}
