package esalert

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ToBuffer(i interface{}) *bytes.Buffer {
	json, err:= json.Marshal(i)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(json)
}

func QueryToJson(query interface{}) ([]byte,error) {
	//query = cleanupMapValue(query)
	return json.Marshal(query)
}

func cleanupInterfaceArray(in []interface{}) []interface{} {
	res := make([]interface{}, len(in))
	for i, v := range in {
		res[i] = cleanupMapValue(v)
	}
	return res
}

func cleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range in {
		res[fmt.Sprintf("%v", k)] = cleanupMapValue(v)
	}
	return res
}

func cleanupMapValue(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return cleanupInterfaceArray(v)
	case map[interface{}]interface{}:
		return cleanupInterfaceMap(v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}