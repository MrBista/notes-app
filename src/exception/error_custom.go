package exception

import "net/http"

type ErrorCustom struct {
	Errors  map[string]string `json:"errors:omitempty"`
	Message string            `json:"message"`
	Code    string            `json:"code"`
	Status  int               `json:"status"`
}

func (err *ErrorCustom) Error() string {
	return err.Message
}

func (err *ErrorCustom) GetStatusHttp() int {

	if err.Status == 0 {
		return http.StatusInternalServerError
	}

	return err.Status
}

const (
	ERR_BAD_REQUEST       = "BAD_REQUEST"
	ERR_VALIDATION_ERROR  = "VALIDATION_ERROR"
	ERR_BUSNIS_VALIDATION = "BUSNIS_VALIDATION_ERROR"
	ERR_NOT_FOUND         = "NOT_FOUND"
	ERR_DB                = "DB_ERROR"
	ERR_UNDHANDLE         = "INTERNAL_SERVER_ERROR"
	ERR_COMMON            = "COMMON_ERR"
)

func NewErrorCommonMessage(message string, status int) *ErrorCustom {
	return &ErrorCustom{
		Message: message,
		Status:  status,
		Code:    ERR_COMMON,
	}
}

func NewBadRequestErr(message string) *ErrorCustom {
	return &ErrorCustom{
		Message: message,
		Status:  http.StatusBadRequest,
		Code:    ERR_BAD_REQUEST,
	}
}

func NewValidationtErr(message string, fieldErrors map[string]string) *ErrorCustom {
	return &ErrorCustom{
		Message: message,
		Status:  http.StatusBadRequest,
		Code:    ERR_VALIDATION_ERROR,
		Errors:  fieldErrors,
	}
}

func NewNotFoundError(message string) *ErrorCustom {
	return &ErrorCustom{
		Message: message,
		Status:  http.StatusNotFound,
		Code:    ERR_NOT_FOUND,
	}
}

func NewUnhandleError(message string) *ErrorCustom {
	return &ErrorCustom{
		Message: message,
		Status:  http.StatusInternalServerError,
		Code:    ERR_UNDHANDLE,
	}
}

func IsCustomError(err error) (*ErrorCustom, bool) {
	dataError, status := err.(*ErrorCustom)
	return dataError, status
}
