package app

import (
	"notes-golang/src/controller"

	"github.com/julienschmidt/httprouter"
)

func userRouter(router *httprouter.Router, userController controller.UserController) {
	router.POST("/api/v1/auth/login", userController.UserLogin)
	router.POST("/api/v1/auth/register", userController.UserRegister)
}

func noteRouter(router *httprouter.Router, noteController controller.NoteController) {
	router.GET("/api/v1/notes", noteController.DeleteNoteById)
	router.GET("/api/v1/notes/:id", noteController.FindNoteById)
	router.POST("/api/v1/notes", noteController.CreateNote)
	router.PUT("/api/v1/notes/:id", noteController.UpdateNote)
	router.DELETE("/api/v1/notes/:id", noteController.DeleteNoteById)

}

func NewRouter(userController controller.UserController, noteController controller.NoteController) *httprouter.Router {
	router := httprouter.New()

	userRouter(router, userController)
	noteRouter(router, noteController)

	return router
}
