package service

import (
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
	return &BookService{repo: bookRepo}
}

func (s *BookService) Create(params *dto.BookReq) error {
	newItem := params.ToEntity()
	return s.repo.Create(&newItem)
}

func (s *BookService) Update(id uint, params *dto.BookUpdateReq) error {
	updateItem := params.ToEntity()
	return s.repo.Update(id, &updateItem)
}

func (s *BookService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *BookService) GetByID(id uint) (dto.BookRes, error) {
	var resp dto.BookRes
	item, err := s.repo.GetByID(id)
	if err != nil {
		return resp, err
	}
	if item == nil {
		return resp, exception.ErrDataNotFound
	}
	resp.FromEntity(item)
	return resp, nil
}

func (s *BookService) GetDataList(params *dto.Filter) ([]dto.BookRes, error) {
	var res []dto.BookRes
	items, err := s.repo.GetListData(params)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		var t dto.BookRes
		t.FromEntity(&item)
		res = append(res, t)
	}
	return res, nil
}
