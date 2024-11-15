package repository

import (
	"base-gin/app/domain/dao"
	"base-gin/exception"
	"base-gin/storage"
	"errors"

	"gorm.io/gorm"
)

type PublishersRepository struct {
	db *gorm.DB
}

func newPublishersRepository(db *gorm.DB) *PublishersRepository {
	return &PublishersRepository{db: db}
}

func (r *PublishersRepository) Create(newItem *dao.Publishers) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *PublishersRepository) GetByID(id uint) (*dao.Publishers, error) {
    ctx, cancelFunc := storage.NewDBContext()
    defer cancelFunc()

    var item dao.Publishers

    tx := r.db.WithContext(ctx).First(&item, id)
    if tx.Error != nil {
        if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
            return nil, exception.ErrUserNotFound
        }

        return nil, tx.Error
    }

    return &item, nil
}