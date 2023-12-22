package response

type Response struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ClientResponse(StatusCode int, message string, data interface{}, err interface{}) Response {
	return Response{
		StatusCode: StatusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}

}
