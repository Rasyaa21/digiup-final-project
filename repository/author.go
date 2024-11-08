package repository

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/storage"
	"errors"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) Create(newAuthor *dao.Author) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Create(&newAuthor)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *AuthorRepository) Update(updateItem *dao.Author) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Updates(&updateItem)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *AuthorRepository) Delete(id uint) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var item dao.Author
	tx := r.db.WithContext(ctx).Delete(id, &item)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *AuthorRepository) FindDataByID(id uint) (*dao.Author, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var item dao.Author
	tx := r.db.WithContext(ctx).First(&item, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, exception.ErrUserNotFound
		}
		return nil, tx.Error
	}
	return &item, tx.Error
}

func (r *AuthorRepository) GetAuthorDataList(params *dto.Filter) ([]dao.Author, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var authors []dao.Author
	tx := r.db.WithContext(ctx).Find(&authors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return authors, tx.Error
}
