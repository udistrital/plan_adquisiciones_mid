package models

import (
	"fmt"
	"reflect"
	"strconv"

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

//ObtenerRegistroModalidadSeleccionByIDPlanAdquisicion ...
func ObtenerRegistroModalidadSeleccionByIDPlanAdquisicion(idStr string) (ModalidadSeleccion []map[string]interface{}, outputError interface{}) {
	var modalidadSeleccion []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_funcionamiento-Modalidad_seleccion/?query=RegistroPlanAdquisicionesId.id%3A"+idStr+"%2CActivo%3Atrue", &modalidadSeleccion)
	if error != nil {
		return nil, error
	} else {
		return modalidadSeleccion, nil
	}

}

//ObtenerRegistroModalidadSeleccionByID ...
func ObtenerRegistroModalidadSeleccionByID(idStr string) (ModalidadSeleccion map[string]interface{}, outputError interface{}) {
	var modalidadSeleccion map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_funcionamiento-Modalidad_seleccion/"+idStr, &modalidadSeleccion)
	if error != nil {
		return nil, error
	} else {
		return modalidadSeleccion, nil
	}

}

//ModalidadSeleccionModificado ...
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

//ActualizarRegistroModalidadSeleccion ...
func ActualizarRegistroModalidadSeleccion(registroModalidadSeleccion map[string]interface{}, idStr string, idStrPlanAdquisicion string) (registroModalidadSeleccionRespuesta map[string]interface{}, outputError interface{}) {
	ModalidadSeleccionPut := make(map[string]interface{})
	ModalidadSeleccionActualizar := make(map[string]interface{})
	RegistroModalidadSeleccionAntiguo, error := ObtenerRegistroModalidadSeleccionByID(idStr)
	if error != nil {
		return nil, error
	} else {
		if RegistroModalidadSeleccionAntiguo["Status"] != nil {
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

//RegistroModalidadSeleccionValidacion ...
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
