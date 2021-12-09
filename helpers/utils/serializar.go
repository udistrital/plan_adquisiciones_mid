package utils

import "encoding/json"

func Serializar(objeto interface{}) (respuesta string, err error) {
	var temp []byte
	if temp, err = json.Marshal(objeto); err != nil {
		return "", err
	}
	respuesta = string(temp)
	return
}
