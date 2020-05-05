package common

import (
	"encoding/json"
	"net/http"
)

type Playground struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(w *http.ResponseWriter, code int, message string, data interface{}) {
	playground := Playground{
		Code:    code,
		Message: message,
		Data:    data,
	}
	jsonValue, err := json.Marshal(playground)
	if err != nil {
		playground.Code = Error
		playground.Message = err.Error()
		playground.Data = nil
		jsonValue, _ = json.Marshal(playground)
		(*w).Write(jsonValue)
	} else {
		(*w).Header().Add("Content-Type", "application/json")
		(*w).Write(jsonValue)
	}
}
