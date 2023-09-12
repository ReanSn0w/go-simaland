package simaland

import (
	"bytes"
	"fmt"
	"net/http"
)

func tryError(resp *http.Response) error {
	if resp.StatusCode != 200 {
		obj := &Error{
			Msg: "Status Code != 200",
		}

		buffer := new(bytes.Buffer)
		buffer.ReadFrom(resp.Body)
		obj.Data = buffer.String()
		return obj
	}

	return nil
}

type Error struct {
	Msg  string
	Data string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Msg: %v, Data: %v", e.Msg, e.Data)
}
