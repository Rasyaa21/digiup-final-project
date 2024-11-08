package service

import (
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/repository"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(AuthorRepo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: AuthorRepo}
}

func (s *AuthorService) Create(params *dto.AuthorReq) error {
	newItem := params.ToEntity()
	return s.repo.Create(&newItem)
}

func (s *AuthorService) Update(params *dto.AuthorUpdateReq) error {
	if params.Fullname == "" {
		return exception.ErrDataNotFound
	}
	updatedItem := params.ToEntity()
	return s.repo.Update(&updatedItem)
}

func (s *AuthorService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *AuthorService) GetList(params *dto.Filter) ([]dto.AuthorRes, error) {
	var res []dto.AuthorRes
	items, err := s.repo.GetAuthorDataList(params)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		var t dto.AuthorRes
		t.FromEntity(&item)
		res = append(res, t)
	}
	return res, nil
}

func (s *AuthorService) GetByID(id uint) (dto.AuthorRes, error) {
	var res dto.AuthorRes
	item, err := s.repo.FindDataByID(id)
	if err != nil {
		return res, err
	}
	if item == nil {
		return res, err
	}
	res.FromEntity(item)
	return res, err
}
