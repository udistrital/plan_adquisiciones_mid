package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//IngresoRegistroModalidadSeleccion ingresa un elemento a la tabla modalid de seleccion
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

//GuardarModalidadSeleccion descompone el array de modalidad de selección para crear uno a uno
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

//ObtenerRegistroModalidadSeleccionByIDPlanAdquisicion regresa una registro de la tabla modalidad de seleccioón segun un Id de un registro_plan_adquisicion
func ObtenerRegistroModalidadSeleccionByIDPlanAdquisicion(idStr string) (ModalidadSeleccion []map[string]interface{}, outputError interface{}) {
	var modalidadSeleccion []map[string]interface{}
	var nombreModalidadSeleccion []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_funcionamiento-Modalidad_seleccion/?query=RegistroPlanAdquisicionesId.id:"+idStr+",Activo:true", &modalidadSeleccion)
	if error != nil {
		return nil, error
	} else {
		for index := range modalidadSeleccion {
			s := fmt.Sprintf("%.0f", modalidadSeleccion[index]["Id"].(float64))
			error := request.GetJson(beego.AppConfig.String("administrativa_crud_api_url")+"modalidad_seleccion/?query=Id:"+s+"&fields=Nombre", &nombreModalidadSeleccion)
			if error != nil {
				return nil, error
			} else {
				modalidadSeleccion[index]["Nombre"] = nombreModalidadSeleccion[0]["Nombre"]
			}
		}
		return modalidadSeleccion, nil
	}

}

//ObtenerRegistroModalidadSeleccionByID regresa una registro de la tabla modalidad de seleccioón segun el ID
func ObtenerRegistroModalidadSeleccionByID(idStr string) (ModalidadSeleccion map[string]interface{}, outputError interface{}) {
	var modalidadSeleccion []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_funcionamiento-Modalidad_seleccion/?query=Id:"+idStr, &modalidadSeleccion)
	if error != nil {
		return nil, error
	} else {
		return modalidadSeleccion[0], nil
	}

}

//ModalidadSeleccionModificado descompone el array de modalidad de selección para actualizar uno a uno
func ModalidadSeleccionModificado(registroPlanAdquisicion map[string]interface{}, idStr string) (outputError interface{}) {
	ModalidadesSeleccion := registroPlanAdquisicion["ModalidadSeleccion"].([]interface{})
	for Index := range ModalidadesSeleccion {
		ModalidadSeleccion := ModalidadesSeleccion[Index].(map[string]interface{})
		_, errRegistroModalidadSeleccion := ActualizarRegistroModalidadSeleccion(ModalidadSeleccion, fmt.Sprintf("%v", ModalidadSeleccion["Id"]), idStr)
		if errRegistroModalidadSeleccion != nil {
			return errRegistroModalidadSeleccion
		}
	}
	return nil
}

//ActualizarRegistroModalidadSeleccion Actualiza la modalidad de selección y la crea en caso de que no exista
func ActualizarRegistroModalidadSeleccion(registroModalidadSeleccion map[string]interface{}, idStr string, idStrPlanAdquisicion string) (registroModalidadSeleccionRespuesta map[string]interface{}, outputError interface{}) {
	ModalidadSeleccionPut := make(map[string]interface{})
	ModalidadSeleccionActualizar := make(map[string]interface{})
	RegistroModalidadSeleccionAntiguo, error := ObtenerRegistroModalidadSeleccionByID(idStr)
	if error != nil {
		return nil, error
	} else {
		if len(RegistroModalidadSeleccionAntiguo) == 0 {
			//fmt.Println("No existe ModalidadSeleccion toca crearlo")
			idint, _ := strconv.Atoi(idStrPlanAdquisicion)
			registroModalidadSeleccion["RegistroPlanAdquisicionesId"] = idint
			RegistroModalidadSeleccion, errRegistroModalidadSeleccion := IngresoRegistroModalidadSeleccion(registroModalidadSeleccion)
			if errRegistroModalidadSeleccion != nil {
				return nil, errRegistroModalidadSeleccion
			} else {
				return RegistroModalidadSeleccion, nil
			}

		} else {
			validacion := RegistroModalidadSeleccionValidacion(registroModalidadSeleccion, RegistroModalidadSeleccionAntiguo)
			if validacion {
				//fmt.Println("existe ModalidadSeleccion No toca modificar")
			} else {
				//fmt.Println("existe ModalidadSeleccion toca modificar")
				idint, _ := strconv.Atoi(idStrPlanAdquisicion)
				registroModalidadSeleccion["RegistroPlanAdquisicionesId"] = idint
				ModalidadSeleccionActualizar = map[string]interface{}{
					"RegistroPlanAdquisicionesId": map[string]interface{}{"Id": registroModalidadSeleccion["RegistroPlanAdquisicionesId"]},
					"FechaCreacion":               RegistroModalidadSeleccionAntiguo["FechaCreacion"],
					"IdModalidadSeleccion":        registroModalidadSeleccion["IdModalidadSeleccion"],
					"Activo":                      registroModalidadSeleccion["Activo"],
				}
				error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_funcionamiento-Modalidad_seleccion/"+idStr, "PUT", &ModalidadSeleccionPut, ModalidadSeleccionActualizar)
				if error != nil {
					return nil, error
				} else {
					return ModalidadSeleccionPut, nil
				}
			}
		}

		return ModalidadSeleccionPut, nil
	}

}

//RegistroModalidadSeleccionValidacion Valida si se requiere actualizar campos , regresa false en caso que se requiera actualizar
func RegistroModalidadSeleccionValidacion(registroModalidadSeleccion map[string]interface{}, RegistroModalidadSeleccionAntiguo map[string]interface{}) (validacion bool) {
	registroModalidadSeleccionActual := make(map[string]interface{})

	registroModalidadSeleccionActual = map[string]interface{}{
		"FechaCreacion":        RegistroModalidadSeleccionAntiguo["FechaCreacion"],
		"IdModalidadSeleccion": RegistroModalidadSeleccionAntiguo["IdModalidadSeleccion"],
		"Activo":               RegistroModalidadSeleccionAntiguo["Activo"],
	}

	if reflect.DeepEqual(registroModalidadSeleccionActual["IdModalidadSeleccion"], registroModalidadSeleccion["IdModalidadSeleccion"]) &&
		reflect.DeepEqual(registroModalidadSeleccionActual["Activo"], registroModalidadSeleccion["Activo"]) {
		return true
	} else {
		return false
	}

}
