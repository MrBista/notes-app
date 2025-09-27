package controller

import (
	"encoding/json"
	"net/http"
	"notes-golang/src/dto/res"
	"notes-golang/src/handler"
	"notes-golang/src/service"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type UserController interface {
	UserRegister(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UserLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) UserRegister(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// 1. ambil dulu bodynya

	decoder := json.NewDecoder(r.Body)

	var userReqBody service.ReqisterReq

	err := decoder.Decode(&userReqBody)

	if err != nil {
		logrus.Error("salah ", err)
		handler.HandleError(w, err)
		return
	}

	result, err := c.UserService.RegisterUser(userReqBody)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)

	responseData := res.NewCommonResponseSuccess(result, "Successfully regis user", http.StatusOK)

	encoder.Encode(responseData)

}

func (c *UserControllerImpl) UserLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// 1. ambil dulu bodynya

	decoder := json.NewDecoder(r.Body)

	var userReqBody service.LoginRequest

	err := decoder.Decode(&userReqBody)

	if err != nil {
		logrus.Error("salah ", err)
		handler.HandleError(w, err)
		return
	}

	data, err := c.UserService.LoginUser(userReqBody)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	webResponse := res.CommonResponseSuccess{
		Data:    data,
		Status:  http.StatusOK,
		Message: "Successfully login",
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)

	encoder.Encode(&webResponse)
}
