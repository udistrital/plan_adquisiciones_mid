package utils

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

func ConvertirInterface(in interface{}, out interface{}) (outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "ConvertirInterface - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	var err error
	var v []byte
	if v, err = json.Marshal(&in); err == nil {
		if err = json.Unmarshal(v, &out); err == nil {
			return nil
		}
	}
	logs.Error(err)
	outputError = map[string]interface{}{
		"funcion": "ConvertirInterface - Marshal o Unmarshal elementos",
		"err":     err,
		"status":  "500",
	}
	return outputError
}
