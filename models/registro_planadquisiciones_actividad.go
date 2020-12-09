package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerRegistroPlanAdquisicionActividad ...
func ObtenerRegistroPlanAdquisicionActividad() (registroPlanAdquisicionActividad []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividad []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Actividad/", &RegistroPlanAdquisicionActividad)

	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividad, nil
	}

}

//IngresoRegistroActividad ...
func IngresoRegistroActividad(registroActividad map[string]interface{}) (registroActividadRespuesta []map[string]interface{}, outputError interface{}) {
	registroActividadIngresado := make(map[string]interface{})
	registroActividadPost := make(map[string]interface{})
	result := []map[string]interface{}{}

	registroActividadIngresado = map[string]interface{}{
		"ActividadId":                 map[string]interface{}{"Id": registroActividad["ActividadId"]},
		"RegistroPlanAdquisicionesId": map[string]interface{}{"Id": registroActividad["RegistroPlanAdquisicionesId"]},
		"Valor":                       registroActividad["Valor"],
		"Activo":                      registroActividad["Activo"],
	}
	FuentesFinanciamiento := registroActividad["FuentesFinanciamiento"].([]interface{})
	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Actividad/", "POST", &registroActividadPost, registroActividadIngresado)
	if error != nil {
		return nil, error
	} else {
		result = append(result, registroActividadPost)
		for fuenteIndex := range FuentesFinanciamiento {
			FuenteFinanciamiento := FuentesFinanciamiento[fuenteIndex].(map[string]interface{})
			FuenteFinanciamiento["ActividadId"] = registroActividadPost["Id"]
			RegistroActividadFuente, errRegistroActividadFuente := IngresoRegistroPlanInversionActividadFuente(FuenteFinanciamiento)
			if errRegistroActividadFuente != nil {
				return nil, errRegistroActividadFuente
			} else {
				result = append(result, RegistroActividadFuente)
			}
		}
		return result, nil
	}

}

//ActualizarRegistroActividad ...
func ActualizarRegistroActividad(registroActividad map[string]interface{}, idStr string) (registroActividadRespuesta map[string]interface{}, outputError interface{}) {
	registroActividadPut := make(map[string]interface{})
	registroActividadActualizar := make(map[string]interface{})
	RegistroPlanAdquisicionActividad, error := ObtenerRegistroPlanAdquisicionActividadByID(idStr)
	if error != nil {
		return nil, error
	} else {
		if RegistroPlanAdquisicionActividad["Status"] != nil {
			//fmt.Println("No existe registro toca crearlo")
			_, errRegistroActividad := IngresoRegistroActividad(registroActividad)
			if errRegistroActividad != nil {
				return nil, errRegistroActividad
			} else {
				return registroActividad, nil
			}

		} else {
			validacion := RegistroModificado(registroActividad, RegistroPlanAdquisicionActividad, idStr)
			if validacion {
				//fmt.Println("existe registro y no toca modificarlo")
				error := FuenteFinanciamientoModificado(registroActividad, idStr)
				if error != nil {
					return nil, error
				} else {
					return RegistroPlanAdquisicionActividad, nil
				}
			} else {
				//fmt.Println("existe registro y  toca modificarlo")
				registroActividadActualizar = map[string]interface{}{
					"ActividadId":                 map[string]interface{}{"Id": registroActividad["ActividadId"]},
					"RegistroPlanAdquisicionesId": map[string]interface{}{"Id": registroActividad["RegistroPlanAdquisicionesId"]},
					"FechaCreacion":               RegistroPlanAdquisicionActividad["FechaCreacion"],
					"Valor":                       registroActividad["Valor"],
					"Activo":                      registroActividad["Activo"],
				}

				error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Actividad/"+idStr, "PUT", &registroActividadPut, registroActividadActualizar)
				if error != nil {
					return nil, error
				} else {
					error := FuenteFinanciamientoModificado(registroActividad, idStr)
					if error != nil {
						return nil, error
					} else {
						return registroActividadPut, nil
					}
				}

			}
		}
	}
}

//GuardarPlanAdquisicionActividad ...
func GuardarPlanAdquisicionActividad(PlanAdquisicionActividades []interface{}, idPost interface{}) (registroPlanAdquisicionActividadRespuesta []map[string]interface{}, outputError interface{}) {
	resultPlanAdquisicionActividad := make([]map[string]interface{}, 0)
	for Index := range PlanAdquisicionActividades {
		PlanAdquisicionActividad := PlanAdquisicionActividades[Index].(map[string]interface{})
		PlanAdquisicionActividad["RegistroPlanAdquisicionesId"] = idPost
		RegistroPlanAdquisicionActividad, errRegistroPlanAdquisicionActividad := IngresoRegistroActividad(PlanAdquisicionActividad)
		if errRegistroPlanAdquisicionActividad != nil {
			return nil, errRegistroPlanAdquisicionActividad
		} else {
			resultPlanAdquisicionActividad = append(resultPlanAdquisicionActividad, RegistroPlanAdquisicionActividad...)
		}
	}
	return resultPlanAdquisicionActividad, nil
}

//ObtenerRegistroPlanAdquisicionActividadByID ...
func ObtenerRegistroPlanAdquisicionActividadByID(idStr string) (registroPlanAdquisicionActividad map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividad map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Actividad/"+idStr, &RegistroPlanAdquisicionActividad)
	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividad, nil
	}

}

//RegistroModificado
func RegistroModificado(registroActividad map[string]interface{}, RegistroPlanAdquisicionActividad map[string]interface{}, idStr string) (validacion bool) {
	registroActividadActual := make(map[string]interface{})

	registroActividadActual = map[string]interface{}{
		"ActividadId":                 map[string]interface{}{"Id": RegistroPlanAdquisicionActividad["ActividadId"]},
		"RegistroPlanAdquisicionesId": map[string]interface{}{"Id": RegistroPlanAdquisicionActividad["RegistroPlanAdquisicionesId"]},
		"FechaCreacion":               RegistroPlanAdquisicionActividad["FechaCreacion"],
		"Valor":                       RegistroPlanAdquisicionActividad["Valor"],
		"Activo":                      RegistroPlanAdquisicionActividad["Activo"],
	}

	id := registroActividadActual["ActividadId"].(map[string]interface{})
	idActividad := id["Id"].(map[string]interface{})
	id = registroActividadActual["RegistroPlanAdquisicionesId"].(map[string]interface{})
	idRegistroPlan := id["Id"].(map[string]interface{})

	if reflect.DeepEqual(idActividad["Id"], registroActividad["ActividadId"]) && reflect.DeepEqual(idRegistroPlan["Id"], registroActividad["RegistroPlanAdquisicionesId"]) && reflect.DeepEqual(registroActividadActual["Valor"], registroActividad["Valor"]) && reflect.DeepEqual(registroActividadActual["Activo"], registroActividad["Activo"]) {
		return true
	} else {
		return false
	}

}

func FuenteFinanciamientoModificado(registroActividad map[string]interface{}, idStr string) (outputError interface{}) {
	FuentesFinanciamiento := registroActividad["FuentesFinanciamiento"].([]interface{})
	for fuenteIndex := range FuentesFinanciamiento {
		FuenteFinanciamiento := FuentesFinanciamiento[fuenteIndex].(map[string]interface{})
		_, errRegistroActividadFuente := ActualizarRegistroActividadFuente(FuenteFinanciamiento, fmt.Sprintf("%v", FuenteFinanciamiento["Id"]), idStr)
		if errRegistroActividadFuente != nil {
			return errRegistroActividadFuente
		}
	}
	return nil
}
