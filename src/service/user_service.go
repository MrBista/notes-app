package service

import "gorm.io/gorm"

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	FullName    string `json:"fullName"`
}

type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `jsong:"password"`
}

type ReqisterReq struct {
}

type UserService interface {
	RegisterUser()
	LoginUser(loginReq LoginRequest) (LoginResponse, error)
}

type UserServiceImpl struct {
	DB *gorm.DB
}

func (s *UserServiceImpl) RegisterUser() {
	panic("not implemented") // TODO: Implement
}

func (s *UserServiceImpl) LoginUser(loginReq LoginRequest) (LoginResponse, error) {
	panic("not implemented") // TODO: Implement
}
