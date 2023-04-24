package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// func ParseBody(r *http.Request, x interface{}) {
// 	if body, err := io.ReadAll(r.Body); err == nil {
// 		if err := json.Unmarshal([]byte(body), x); err != nil {
// 			return
// 		}
// 	}
// }

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, x); err != nil {
		return err
	}

	return nil
}
