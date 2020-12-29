package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerRegistroPlanAdquisicionActividad regresa todos los elementos de la tabla registro_plan_adquisicion_Actividad
func ObtenerRegistroPlanAdquisicionActividad() (registroPlanAdquisicionActividad []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividad []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Actividad/", &RegistroPlanAdquisicionActividad)

	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividad, nil
	}

}

//IngresoRegistroActividad crea un nuevo elemento en la tabla registro_plan_adquisicion_Actividad
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

//RegistroActividadModificado descompone un array para actualizar uno a uno un registro de actividad
func RegistroActividadModificado(registroPlanAdquisicion map[string]interface{}, idStr string) (outputError interface{}) {
	RegistroActividades := registroPlanAdquisicion["RegistroPlanAdquisicionActividad"].([]interface{})
	for Index := range RegistroActividades {
		RegistroActividad := RegistroActividades[Index].(map[string]interface{})
		idFloat, _ := strconv.ParseFloat(idStr, 64)
		RegistroActividad["RegistroPlanAdquisicionesId"] = idFloat
		_, errRegistroActividad := ActualizarRegistroActividad(RegistroActividad, fmt.Sprintf("%v", RegistroActividad["RegistroActividadId"]))
		if errRegistroActividad != nil {
			return errRegistroActividad
		}
	}
	return nil
}

//ActualizarRegistroActividad actualiza registro de actividad, si no existe se crea
func ActualizarRegistroActividad(registroActividad map[string]interface{}, idStr string) (registroActividadRespuesta map[string]interface{}, outputError interface{}) {
	registroActividadPut := make(map[string]interface{})
	registroActividadActualizar := make(map[string]interface{})
	RegistroPlanAdquisicionActividad, error := ObtenerRegistroPlanAdquisicionActividadByID(idStr)
	if error != nil {
		return nil, error
	} else {
		if len(RegistroPlanAdquisicionActividad) == 0 {
			//fmt.Println("No existe registro Actividad toca crearlo")
			_, errRegistroActividad := IngresoRegistroActividad(registroActividad)
			if errRegistroActividad != nil {
				return nil, errRegistroActividad
			} else {
				return registroActividad, nil
			}

		} else {
			validacion := RegistroActividadValidacion(registroActividad, RegistroPlanAdquisicionActividad, idStr)
			if validacion {
				//fmt.Println("existe registro Actividad y no toca modificarlo")
				error := FuenteFinanciamientoModificado(registroActividad, idStr)
				if error != nil {
					return nil, error
				} else {
					return RegistroPlanAdquisicionActividad, nil
				}
			} else {
				//fmt.Println("existe registro Actividad y  toca modificarlo")
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

//GuardarPlanAdquisicionActividad descompone array para ingresar uno a uno los registros de actividad
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

//ObtenerRegistroPlanAdquisicionActividadByID un registro de actividad segun su ID
func ObtenerRegistroPlanAdquisicionActividadByID(idStr string) (registroPlanAdquisicionActividad map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividad []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Actividad/?query=id:"+idStr, &RegistroPlanAdquisicionActividad)
	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividad[0], nil
	}

}

//RegistroActividadValidacion valida si se modificaron campos de un registro de actividad
func RegistroActividadValidacion(registroActividad map[string]interface{}, RegistroPlanAdquisicionActividad map[string]interface{}, idStr string) (validacion bool) {
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
	//id = registroActividadActual["RegistroPlanAdquisicionesId"].(map[string]interface{})
	//idRegistroPlan := id["Id"].(map[string]interface{})

	if reflect.DeepEqual(idActividad["Id"], registroActividad["ActividadId"]) &&
		reflect.DeepEqual(registroActividadActual["Valor"], registroActividad["Valor"]) &&
		reflect.DeepEqual(registroActividadActual["Activo"], registroActividad["Activo"]) {
		return true
	} else {
		return false
	}

}

//FuenteFinanciamientoModificado valida si se modificaron campos de una fuente de financiamiento
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
