package app

import (
	"notes-golang/src/controller"
	"notes-golang/src/middleware"

	"github.com/julienschmidt/httprouter"
)

func userRouter(router *httprouter.Router, userController controller.UserController) {
	router.POST("/api/v1/auth/login", userController.UserLogin)
	router.POST("/api/v1/auth/register", userController.UserRegister)
}

func noteRouter(router *httprouter.Router, noteController controller.NoteController) {
	router.GET("/api/v1/notes", middleware.AuthMiddleware(noteController.FindAllNote))
	router.GET("/api/v1/notes/:id", middleware.AuthMiddleware(noteController.FindNoteById))
	router.POST("/api/v1/notes", middleware.AuthMiddleware(noteController.CreateNote))
	router.PUT("/api/v1/notes/:id", middleware.AuthMiddleware(noteController.UpdateNote))
	router.DELETE("/api/v1/notes/:id", middleware.AuthMiddleware(noteController.DeleteNoteById))

}

func NewRouter(userController controller.UserController, noteController controller.NoteController) *httprouter.Router {
	router := httprouter.New()

	userRouter(router, userController)
	noteRouter(router, noteController)

	return router
}
