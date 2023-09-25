package simaland

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func tryError(resp *http.Response) error {
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		obj := &Error{
			Msg: "Non 2xx status",
		}

		buffer := new(bytes.Buffer)
		buffer.ReadFrom(resp.Body)
		err := json.NewDecoder(buffer).Decode(&obj.Data)
		if err != nil {
			obj.Data = map[string]interface{}{
				"json_error": err.Error(),
				"raw_body":   buffer.String(),
			}
		}

		return obj
	}

	return nil
}

type Error struct {
	Msg  string
	Data map[string]interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Msg: %v, Data: %v", e.Msg, e.Data)
}
