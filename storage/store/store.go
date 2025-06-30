package store

import (
	"fmt"

	"github.com/inventory/storage/models"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func (store *UserStore) CreateUser(user *models.User) (*models.User, error) {

	results := store.db.Create(&user)

	if results.Error != nil {
		fmt.Println("Error creating the user", results.Error)
		return nil, results.Error
	}

	return user, nil
}

func (store *UserStore) GetAllUsers() (*[]models.User, error) {
	users := &[]models.User{}
	results := store.db.Find(users)

	if results.Error != nil {
		fmt.Println("No users forund", results.Error)
		return nil, results.Error
	}

	return users, nil
}

func (store *UserStore) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}

	results := store.db.Where("id = ?", id).Find(user)
	if results.Error != nil {
		fmt.Println("this is the error", results.Error)
		return nil, results.Error
	}

	return user, nil
}

func (store *UserStore) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	results := store.db.Where("email = ?", email).Find(user)
	if results.Error != nil {
		fmt.Println("Doesn't exits", results.Error)
		return nil, results.Error
	}

	return user, nil
}

// this the sstore for

func (store *UserStore) CreateStore(storeDetails *models.Store) (*models.Store, error) {
	results := store.db.Create(&storeDetails)
	if results.Error != nil {
		fmt.Println("Doesn't exits", results.Error)
		return nil, results.Error
	}

	return storeDetails, nil
}

// get the store by id
