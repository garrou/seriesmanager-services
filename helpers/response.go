package helpers

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string, data interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Data:    data,
	}
}
