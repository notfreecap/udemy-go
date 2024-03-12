package models

import (
	"encoding/json"
	"io"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func (e *ErrorResponse) ReturnError(w io.Writer) {
	json.NewEncoder(w).Encode(e)
}
