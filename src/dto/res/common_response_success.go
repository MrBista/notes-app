package res

type CommonResponseSuccess struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

func NewCommonResponseSuccess(data interface{}, message string, status int) *CommonResponseSuccess {

	return &CommonResponseSuccess{
		Data:    data,
		Message: message,
		Status:  status,
	}

}
