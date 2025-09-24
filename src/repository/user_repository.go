package repository

import (
	"errors"
	"notes-golang/src/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByIdentifier(identifier string) (models.User, error)
	CreateUser(user models.User) error
	CreateUserWithTx(user models.User) error
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (f *UserRepositoryImpl) FindUserByIdentifier(identifier string) (models.User, error) {
	var findUser models.User
	err := f.Db.Take(&findUser, "username = ? or email = ? ", identifier, identifier).Error

	return findUser, err

}

func (f *UserRepositoryImpl) CreateUser(user models.User) error {
	createUser := f.Db.Create(&user)

	err := createUser.Error

	if err != nil {
		return err
	}

	if createUser.RowsAffected == 0 {
		return errors.New("no row affacted")
	}

	return nil

}

func (f *UserRepositoryImpl) CreateUserWithTx(user models.User) error {

	err := f.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&user).Error

		return err
	})

	return err

}
