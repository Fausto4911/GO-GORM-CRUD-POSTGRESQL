package repository

import (
	"github.com/jinzhu/gorm"
	"go-gorm-crud-postgresql/model"
)

type ClientRepository interface {
	Save(client *model.Client, db *gorm.DB) error
	FindById(id uint32, db *gorm.DB) (model.Client, error)
	FindAll(db *gorm.DB) ([]model.Client, error)
	Update(client *model.Client, db *gorm.DB) (*model.Client, error)
	Delete(client *model.Client, db *gorm.DB) error
}
