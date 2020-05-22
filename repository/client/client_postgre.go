package client

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go-gorm-crud-postgresql/model"
)

type IClientRepository struct {
}

func (repo IClientRepository) Save(client *model.Client, db *gorm.DB) error {
	b := db.NewRecord(client) // => returns `true` as primary key is blank
	if !b {
		errors.New("Require ID as zero value")
	}
	db.Create(client)
	defer db.Close()
	b = db.NewRecord(client) // => return `false` after `user` created
	if b {
		errors.New("Cannot insert object into database")
	}
	return nil
}
func (repo IClientRepository) FindById(id uint32, db *gorm.DB) (model.Client, error) {
	client := model.Client{
		Id: uint(id),
	}
	// Get record with primary key (only works for integer primary key)
	db.First(&client, id)
	defer db.Close()
	return client, nil
}
func (repo IClientRepository) FindAll(db *gorm.DB) ([]model.Client, error) {
	clients := []model.Client{}
	// Get all records
	db.Find(&clients)
	defer db.Close()
	return clients, nil
}
func (repo IClientRepository) Update(client *model.Client, db *gorm.DB) (*model.Client, error) {
	db.Save(&client)
	defer db.Close()
	return client, nil
}
func (repo IClientRepository) Delete(client *model.Client, db *gorm.DB) error {
	// Delete an existing record
	db.Delete(&client)
	defer db.Close()
	return nil
}
