package models

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//IngresoRegistroModalidadSeleccion ...
func IngresoRegistroModalidadSeleccion(registroModalidadSeleccion map[string]interface{}) (registroModalidadSeleccionRespuesta map[string]interface{}, outputError interface{}) {
	registroModalidadSeleccionIngresado := make(map[string]interface{})
	registroModalidadSeleccionPost := make(map[string]interface{})

	registroModalidadSeleccionIngresado = map[string]interface{}{
		"RegistroPlanAdquisicionesId": map[string]interface{}{"Id": registroModalidadSeleccion["RegistroPlanAdquisicionesId"]},
		"IdModalidadSeleccion":        registroModalidadSeleccion["IdModalidadSeleccion"],
		"Activo":                      registroModalidadSeleccion["Activo"],
	}
	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_funcionamiento-Modalidad_seleccion/", "POST", &registroModalidadSeleccionPost, registroModalidadSeleccionIngresado)
	if error != nil {
		return nil, error
	} else {
		return registroModalidadSeleccionPost, nil
	}

}

//GuardarModalidadSeleccion ...
func GuardarModalidadSeleccion(ModalidadesSeleccion []interface{}, idPost interface{}) (registroModalidadSeleccionRespuesta []map[string]interface{}, outputError interface{}) {
	resultModalidad := make([]map[string]interface{}, 0)
	for Index := range ModalidadesSeleccion {
		ModalidadSeleccion := ModalidadesSeleccion[Index].(map[string]interface{})
		ModalidadSeleccion["RegistroPlanAdquisicionesId"] = idPost
		RegistroModalidadSeleccion, errRegistroModalidadSeleccion := IngresoRegistroModalidadSeleccion(ModalidadSeleccion)
		if errRegistroModalidadSeleccion != nil {
			return nil, errRegistroModalidadSeleccion
		} else {
			resultModalidad = append(resultModalidad, RegistroModalidadSeleccion)
		}
	}
	return resultModalidad, nil
}

//ObtenerRegistroModalidadSeleccionByID ...
func ObtenerRegistroModalidadSeleccionByID(idStr string) (ModalidadSeleccion []map[string]interface{}, outputError interface{}) {
	var modalidadSeleccion []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_funcionamiento-Modalidad_seleccion/?query=RegistroPlanAdquisicionesId.id%3A"+idStr, &modalidadSeleccion)
	if error != nil {
		return nil, error
	} else {
		return modalidadSeleccion, nil
	}

}
