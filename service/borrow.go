package service

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/repository"
)

type BorrowService struct {
	repo *repository.BorrowRepository
}

func NewBorrowService(borrowRepo *repository.BorrowRepository) *BorrowService {
	return &BorrowService{repo: borrowRepo}
}

func (s *BorrowService) BorrowBook(params *dto.BorrowReq) (*dao.Borrow, error) {
	newBorrow := params.ToEntity()
	err := s.repo.BorrowBook(&newBorrow)
	if err != nil {
		return nil, err
	}
	return &newBorrow, nil
}

func (s *BorrowService) ReturnBorrow(id uint) error {
	return s.repo.DeleteBorrowData(id)
}

func (s *BorrowService) GetByID(id uint) (dto.BorrowRes, error) {
	var resp dto.BorrowRes

	item, err := s.repo.GetDataByID(id)
	if err != nil {
		return resp, err
	}
	if item == nil {
		return resp, exception.ErrDataNotFound
	}
	resp.FromEntity(item)
	return resp, nil
}

func (s *BorrowService) GetList(params *dto.Filter) ([]dto.BorrowRes, error) {
	var resp []dto.BorrowRes

	items, err := s.repo.GetListdata(params)
	if err != nil {
		return nil, err
	}
	if len(items) < 1 {
		return nil, exception.ErrDataNotFound
	}
	for _, item := range items {
		var t dto.BorrowRes
		t.FromEntity(&item)
		resp = append(resp, t)
	}
	return resp, nil
}

func (s *BorrowService) Update(id uint, params *dto.UpdateBorrowReq) error {
	updateItem := params.ToEntity()
	return s.repo.UpdateBorrow(id, &updateItem)
}
