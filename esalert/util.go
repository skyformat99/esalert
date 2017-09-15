package esalert

import (
	"bytes"
	"encoding/json"
)

func ToBuffer(i interface{}) *bytes.Buffer {
	json, err:= json.Marshal(i)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(json)
}
