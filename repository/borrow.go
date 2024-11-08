package repository

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/storage"

	"gorm.io/gorm"
)

type BorrowRepository struct {
	db *gorm.DB
}

func NewBorrowRepo(db *gorm.DB) *BorrowRepository {
	return &BorrowRepository{db: db}
}

func (r *BorrowRepository) BorrowBook(newItem *dao.Borrow) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	tx := r.db.WithContext(ctx).Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *BorrowRepository) UpdateBorrow(id uint, newItem *dao.Borrow) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var item dao.Borrow
	tx := r.db.WithContext(ctx).Model(&item).Updates(&newItem)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *BorrowRepository) DeleteBorrowData(id uint) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var data dao.Borrow
	tx := r.db.WithContext(ctx).Delete(id, &data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *BorrowRepository) GetDataByID(id uint) (*dao.Borrow, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var data dao.Borrow
	tx := r.db.WithContext(ctx).First(&data, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return nil, tx.Error
}

func (r *BorrowRepository) GetListdata(params *dto.Filter) ([]dao.Borrow, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	var data []dao.Borrow
	tx := r.db.WithContext(ctx).Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return data, tx.Error
}
