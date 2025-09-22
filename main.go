package main

import (
	"encoding/json"
	"net/http"
	"notes-golang/config"
	"notes-golang/src/controller"
	"notes-golang/src/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func main() {
	db := config.OpenConectionDb()

	validate := validator.New()

	router := httprouter.New()

	service := service.NewUserServiceImpl(db, validate)
	controller := controller.NewUserControllerImpl(service)

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		encoder := json.NewEncoder(w)
		encoder.Encode("Hello World")
	})
	router.POST("/api/v1/auth/login", controller.UserLogin)
	router.POST("/api/v1/auth/register", controller.UserRegister)

	port := 8000
	server := http.Server{
		Addr:    "127.0.0.1:" + strconv.Itoa(port),
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}
