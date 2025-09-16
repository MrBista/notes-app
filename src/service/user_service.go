package service

import (
	"errors"
	"fmt"
	"notes-golang/src/models"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	FullName    string `json:"fullName"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `jsong:"password" validate:"required"`
}

type ReqisterReq struct {
	Username string `json:"username" validate:"required, min=3"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"fullName" validate:"required"`
}

type RegisterRes struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserService interface {
	RegisterUser(registerReq ReqisterReq) (RegisterRes, error)
	LoginUser(loginReq LoginRequest) (LoginResponse, error)
}

type UserServiceImpl struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ComparePassword(password, hashPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return fmt.Errorf("password tidak tepat %w", err)
		}
		return err
	}

	return nil
}

func (s *UserServiceImpl) RegisterUser(registerReq ReqisterReq) (RegisterRes, error) {

	registerResponse := RegisterRes{}

	if err := s.Validate.Struct(registerReq); err != nil {
		return registerResponse, err
	}

	// 1. hash passwordnya
	// 2. insert
	passwordHashed, errHash := HashPassword(registerReq.Password)
	if errHash != nil {
		return registerResponse, errHash
	}

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		user := models.User{
			FullName: registerReq.FullName,
			Email:    registerReq.Email,
			Username: registerReq.Username,
			Password: passwordHashed,
			Status:   1,
		}

		err := tx.Create(user).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return registerResponse, fmt.Errorf("terjadi kesalahan %w", err)
	}

	registerResponse.Email = registerReq.Email
	registerResponse.FullName = registerReq.FullName
	registerResponse.Username = registerReq.Username
	registerResponse.CreatedAt = time.Now()

	return registerResponse, nil

}

func (s *UserServiceImpl) LoginUser(loginReq LoginRequest) (LoginResponse, error) {
	responseLogin := LoginResponse{}

	err := s.Validate.Struct(loginReq)
	if err != nil {

		return responseLogin, err
	}
	// 1. cek dulu ada ga username/email nya
	// 2. kalau ga ada maka throw
	// 3. kalau ada maka cek apakah passowrdnya match
	// 4. kalau ga match maka throw
	// 5. lalu generate jwt sebagai response

	var user models.User
	err = s.DB.Take(&user, "email = ? or username = ?", loginReq.Identifier, loginReq.Identifier).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responseLogin, errors.New("User not found")
		}
		return responseLogin, err
	}

	if err := ComparePassword(loginReq.Password, user.Password); err != nil {
		return responseLogin, err
	}

	tokenGenerated := "ini adalah token generated"

	loginResponse := LoginResponse{
		AccessToken: tokenGenerated,
		FullName:    user.FullName,
	}

	return loginResponse, nil
}
