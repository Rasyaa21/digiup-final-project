package dto

import (
	"base-gin/domain/dao"
	"time"
)

type BookReq struct {
	Title       string `json:"title" binding:"required,min=1,max=56"`
	Subtitle    string `json:"subtitle" binding:"omitempty,max=64"`
	PublisherID uint   `json:"publisher_id" binding:"required"`
	AuthorID    uint   `json:"author_id" binding:"required"`
}

type BookRes struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle,omitempty"`
	PublisherID uint      `json:"publisher_id"`
	AuthorID    uint      `json:"author_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type BookUpdateReq struct {
	Title       *string `json:"title" binding:"omitempty,min=1,max=56"`
	Subtitle    *string `json:"subtitle" binding:"omitempty,max=64"`
	PublisherID *uint   `json:"publisher_id" binding:"omitempty"`
	AuthorID    *uint   `json:"author_id" binding:"omitempty"`
}

func (req *BookReq) ToEntity() dao.Book {
	var item dao.Book
	item.Title = req.Title
	item.Subtitle = &req.Subtitle
	item.PublisherID = req.PublisherID
	item.AuthorID = req.AuthorID
	return item
}

func (req *BookUpdateReq) ToEntity() dao.Book {
	var item dao.Book
	item.Title = *req.Title
	item.Subtitle = req.Subtitle
	item.PublisherID = *req.PublisherID
	item.AuthorID = *req.AuthorID
	return item
}

func (res *BookRes) FromEntity(item *dao.Book) {
	item.Title = res.Title
	item.Subtitle = &res.Subtitle
	item.AuthorID = res.AuthorID
	item.PublisherID = res.PublisherID
}
