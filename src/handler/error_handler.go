package handler

import (
	"encoding/json"
	"net/http"
	"notes-golang/src/exception"
)

func HandleError(w http.ResponseWriter, err error) {

	if customErr, ok := err.(*exception.ErrorCustom); ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(customErr.GetStatusHttp())

		encoder := json.NewEncoder(w)

		err := encoder.Encode(customErr)

		if err != nil {
			panic(err)
		}

		return
	}

	unhandleError := exception.NewUnhandleError(err.Error())

	w.Header().Add("Content-Type", "applicatin/json")
	w.WriteHeader(http.StatusInternalServerError)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(&unhandleError); err != nil {
		panic(err)
	}

}
