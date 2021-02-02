package models

type BaseModel interface {
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func Ok(data BaseModel) *Response {
	return &Response{
		Code: 200,
		Data: data,
	}
}
