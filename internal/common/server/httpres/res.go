package httpres

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data any, code ...int) {
	statusCode := http.StatusOK
	if len(code) > 0 {
		statusCode = code[0]
	}

	w.WriteHeader(statusCode)

	res, _ := json.Marshal(data)

	w.Write(res)
}

func Bind[T any](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, err
	}

	return req, nil
}
