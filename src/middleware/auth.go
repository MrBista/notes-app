package middleware

import (
	"encoding/json"
	"net/http"
	"notes-golang/src/dto/res"
	"notes-golang/src/helper"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")

		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := res.CommonResponseSuccess{
				Data:    false,
				Status:  http.StatusUnauthorized,
				Message: "Invalid header",
			}

			encoder := json.NewEncoder(w)
			encoder.Encode(&webResponse)
			return
		}

		splitBearerToken := strings.Split(bearerToken, " ")

		if len(splitBearerToken) != 2 || splitBearerToken[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := res.CommonResponseSuccess{
				Data:    false,
				Status:  http.StatusUnauthorized,
				Message: "Invalid header type name",
			}

			encoder := json.NewEncoder(w)
			encoder.Encode(&webResponse)
			return
		}

		accessToken := splitBearerToken[1]

		err := helper.VerifyToken(accessToken)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := res.CommonResponseSuccess{
				Data:    false,
				Status:  http.StatusUnauthorized,
				Message: "Invalid or expired token",
			}

			encoder := json.NewEncoder(w)
			encoder.Encode(&webResponse)
			return
		}

		handle(w, r, p)
	}
}
