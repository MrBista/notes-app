package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// db := config.OpenConectionDb()

	router := httprouter.New()

	router.GET("/api/v1/auth/login", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	})

	port := 8000
	server := http.Server{
		Addr: "127.0.0.1:" + strconv.Itoa(port),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
