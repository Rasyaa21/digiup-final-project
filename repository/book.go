package repository

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/storage"
	"errors"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func newBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(newItem *dao.Book) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	tx := r.db.WithContext(ctx).Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *BookRepository) Update(id uint, newItem *dao.Book) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var item dao.Book
	tx := r.db.WithContext(ctx).Model(&item).Updates(&newItem)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *BookRepository) Delete(id uint) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var item dao.Book
	tx := r.db.WithContext(ctx).Delete(id, &item)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *BookRepository) GetByID(id uint) (*dao.Book, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var item dao.Book
	tx := r.db.WithContext(ctx).Joins("BookPublisher").
		First(&item, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, exception.ErrUserNotFound
		}
		return nil, tx.Error
	}
	return &item, nil
}

func (r *BookRepository) GetListData(params *dto.Filter) ([]dao.Book, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var book []dao.Book
	tx := r.db.WithContext(ctx).Find(&book)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return book, tx.Error
}
