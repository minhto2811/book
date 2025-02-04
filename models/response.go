package models

type Response struct {
	Message interface{} `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func FullResponse(message interface{}, err interface{}, data interface{}) *Response {
	return &Response{
		Message: message,
		Error:   err,
		Data:    data,
	}
}

func ErrorResponse(err interface{}) *Response {
	return &Response{
		Message: nil,
		Error:   err,
		Data:    nil,
	}
}

func MessageResponse(message interface{}) *Response {
	return &Response{
		Message: message,
		Error:   nil,
		Data:    nil,
	}
}

func DataResponse(data interface{}) *Response {
	return &Response{
		Message: nil,
		Error:   nil,
		Data:    data,
	}
}

func DataAndMessageResponse(data interface{}, message interface{}) *Response {
	return &Response{
		Message: message,
		Data:    data,
		Error:   nil,
	}
}
