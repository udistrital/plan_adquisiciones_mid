package models

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//IngresoRegistroCodigoArka ...
func IngresoRegistroCodigoArka(registroCodigoArka map[string]interface{}) (registroCodigoArkaRespuesta map[string]interface{}, outputError interface{}) {
	registroCodigoArkaIngresado := make(map[string]interface{})
	registroCodigoArkaPost := make(map[string]interface{})

	registroCodigoArkaIngresado = map[string]interface{}{
		"RegistroPlanAdquisicionesId": map[string]interface{}{"Id": registroCodigoArka["RegistroPlanAdquisicionesId"]},
		"CodigoArka":                  registroCodigoArka["CodigoArka"],
		"Activo":                      registroCodigoArka["Activo"],
	}
	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Codigo_arka/", "POST", &registroCodigoArkaPost, registroCodigoArkaIngresado)
	if error != nil {
		return nil, error
	} else {
		return registroCodigoArkaPost, nil
	}

}

//GuardarCodigoArka ...
func GuardarCodigoArka(CodigosArka []interface{}, idPost interface{}) (registroCodigosArkaRespuesta []map[string]interface{}, outputError interface{}) {
	resultCodigo := make([]map[string]interface{}, 0)
	for Index := range CodigosArka {
		CodigoArka := CodigosArka[Index].(map[string]interface{})
		CodigoArka["RegistroPlanAdquisicionesId"] = idPost
		RegistroCodigoArka, errRegistroCodigoArka := IngresoRegistroCodigoArka(CodigoArka)
		if errRegistroCodigoArka != nil {
			return nil, errRegistroCodigoArka
		} else {
			resultCodigo = append(resultCodigo, RegistroCodigoArka)
		}
	}
	return resultCodigo, nil
}

//ObtenerRegistroCodigoArkaByID ...
func ObtenerRegistroCodigoArkaByID(idStr string) (CodigoArka []map[string]interface{}, outputError interface{}) {
	var codigoArka []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Codigo_arka/?query=RegistroPlanAdquisicionesId.id%3A"+idStr, &codigoArka)
	if error != nil {
		return nil, error
	} else {
		return codigoArka, nil
	}

}
