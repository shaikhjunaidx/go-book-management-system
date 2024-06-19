package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody[T any](r *http.Request, x *T) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
