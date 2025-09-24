package main

import (
	"net/http"
	"notes-golang/config"
	"notes-golang/src/controller"
	"notes-golang/src/repository"
	"notes-golang/src/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func main() {
	db := config.OpenConectionDb()

	logrus.SetLevel(logrus.DebugLevel)

	validate := validator.New()

	router := httprouter.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserServiceImpl(db, validate, userRepository)
	userController := controller.NewUserControllerImpl(userService)

	router.POST("/api/v1/auth/login", userController.UserLogin)
	router.POST("/api/v1/auth/register", userController.UserRegister)

	noteRepository := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(db, noteRepository)
	noteController := controller.NewNoteController(noteService)

	router.GET("/api/v1/notes", noteController.DeleteNoteById)
	router.GET("/api/v1/notes/:id", noteController.FindNoteById)
	router.POST("/api/v1/notes", noteController.CreateNote)
	router.PUT("/api/v1/notes/:id", noteController.UpdateNote)
	router.DELETE("/api/v1/notes/:id", noteController.DeleteNoteById)

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
