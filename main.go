package main

import (
	"net/http"
	"notes-golang/config"
	"notes-golang/src/app"
	"notes-golang/src/controller"
	"notes-golang/src/repository"
	"notes-golang/src/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func main() {
	db := config.OpenConectionDb()

	logrus.SetLevel(logrus.DebugLevel)

	validate := validator.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserServiceImpl(db, validate, userRepository)
	userController := controller.NewUserControllerImpl(userService)

	noteRepository := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(db, noteRepository)
	noteController := controller.NewNoteController(noteService)

	router := app.NewRouter(userController, noteController)

	port := 8001
	server := http.Server{
		Addr:    "127.0.0.1:" + strconv.Itoa(port),
		Handler: router,
	}
	logrus.Info("App runing in port ", port)

	err := server.ListenAndServe()

	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}
