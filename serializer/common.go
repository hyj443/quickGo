package serializer

import (
	"time"
)

type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Error     int         `json:"error"`
	TimeStamp int64       `json:"timestamp"`
}

func (r Response) Format() Response {
	r.TimeStamp = time.Now().Unix()
	return r
}
