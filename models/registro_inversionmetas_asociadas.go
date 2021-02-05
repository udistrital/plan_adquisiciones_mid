package models

import (
	"fmt"
	// "reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//IngresoRegistroMetasAsociadas ingresa un elemento a la tabla modalid de seleccion
func IngresoRegistroMetasAsociadas(registroMetasAsociadas map[string]interface{}) (registroMetasAsociadasRespuesta map[string]interface{}, outputError interface{}) {
	registroMetasAsociadasIngresado := make(map[string]interface{})
	registroMetasAsociadasPost := make(map[string]interface{})

	registroMetasAsociadasIngresado = map[string]interface{}{
		"RegistroPlanAdquisicionesId": 	map[string]interface{}{"Id": registroMetasAsociadas["RegistroPlanAdquisicionesId"]},
		"MetaId":        				map[string]interface{}{"Id": registroMetasAsociadas["MetaId"]},
		"Activo":                      	registroMetasAsociadas["Activo"],
	}
	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Metas_Asociadas/", "POST", &registroMetasAsociadasPost, registroMetasAsociadasIngresado)
	if error != nil {
		return nil, error
	} else {
		m := ExtraerDataPeticion(registroMetasAsociadasPost)
		return m, nil
	}

}

//GuardarMetasAsociadas descompone el array de modalidad de selección para crear uno a uno
func GuardarMetasAsociadas(Metas_Asociadas []interface{}, idPost interface{}) (registroMetasAsociadasRespuesta []map[string]interface{}, outputError interface{}) {
	resultModalidad := make([]map[string]interface{}, 0)
	for Index := range Metas_Asociadas {
		MetasAsociadas := Metas_Asociadas[Index].(map[string]interface{})
		MetasAsociadas["RegistroPlanAdquisicionesId"] = idPost
		RegistroMetasAsociadas, errRegistroMetasAsociadas := IngresoRegistroMetasAsociadas(MetasAsociadas)
		if errRegistroMetasAsociadas != nil {
			return nil, errRegistroMetasAsociadas
		} else {
			resultModalidad = append(resultModalidad, RegistroMetasAsociadas)
		}
	}
	return resultModalidad, nil
}

//ObtenerRegistroMetasAsociadasByIDPlanAdquisicion regresa una registro de la tabla modalidad de seleccioón segun un Id de un registro_plan_adquisicion
func ObtenerRegistroMetasAsociadasByIDPlanAdquisicion(idStr string) (MetasAsociadas []map[string]interface{}, outputError interface{}) {
	var metasAsociadas map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Metas_Asociadas/?query=RegistroPlanAdquisicionesId.Id:"+idStr+",Activo:true&fields=Id,MetaId,Activo", &metasAsociadas)
	if error != nil {
		return nil, error
	} else {
		fmt.Println(metasAsociadas)
		m := ExtraerDataPeticionArreglo(metasAsociadas)
		return m, nil
	}

}

//ObtenerRegistroMetasAsociadasByID regresa una registro de la tabla modalidad de seleccioón segun el ID
func ObtenerRegistroMetasAsociadasByID(idStr string) (MetasAsociadas map[string]interface{}, outputError interface{}) {
	var metasAsociadas map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Metas_Asociadas/?query=Id:"+idStr+"&fields=Id,MetaId,Activo", &metasAsociadas)
	if error != nil {
		return nil, error
	} else {
		fmt.Println(metasAsociadas)
		m := ExtraerDataPeticionArreglo(metasAsociadas)
		return m[0], nil
	}

}

//MetasAsociadasModificado descompone el array de modalidad de selección para actualizar uno a uno
func MetasAsociadasModificado(registroPlanAdquisicion map[string]interface{}, idStr string) (outputError interface{}) {
	Metas_Asociadas := registroPlanAdquisicion["MetasAsociadas"].([]interface{})
	for Index := range Metas_Asociadas {
		MetasAsociadas := Metas_Asociadas[Index].(map[string]interface{})
		_, errRegistroMetasAsociadas := ActualizarRegistroMetasAsociadas(MetasAsociadas, fmt.Sprintf("%v", MetasAsociadas["Id"]), idStr)
		if errRegistroMetasAsociadas != nil {
			return errRegistroMetasAsociadas
		}
	}
	return nil
}

//ActualizarRegistroMetasAsociadas Actualiza la modalidad de selección y la crea en caso de que no exista
func ActualizarRegistroMetasAsociadas(registroMetasAsociadas map[string]interface{}, idStr string, idStrPlanAdquisicion string) (registroMetasAsociadasRespuesta map[string]interface{}, outputError interface{}) {
	MetasAsociadasPut := make(map[string]interface{})
	MetasAsociadasActualizar := make(map[string]interface{})
	

	if registroMetasAsociadas["Id"] == nil {

		idint, _ := strconv.Atoi(idStrPlanAdquisicion)
		registroMetasAsociadas["RegistroPlanAdquisicionesId"] = idint
		RegistroMetasAsociadas, errRegistroMetasAsociadas := IngresoRegistroMetasAsociadas(registroMetasAsociadas)
		if errRegistroMetasAsociadas != nil {
			return nil, errRegistroMetasAsociadas
		} else {
			return RegistroMetasAsociadas, nil
		}
	} else {
		RegistroMetasAsociadasAntiguo, _ := ObtenerRegistroMetasAsociadasByID(idStr)
		idint, _ := strconv.Atoi(idStrPlanAdquisicion)
		registroMetasAsociadas["RegistroPlanAdquisicionesId"] = idint
		MetasAsociadasActualizar = map[string]interface{}{
			"RegistroPlanAdquisicionesId": 	map[string]interface{}{"Id": registroMetasAsociadas["RegistroPlanAdquisicionesId"]},
			"FechaCreacion":               	RegistroMetasAsociadasAntiguo["FechaCreacion"],
			"MetaId":        				map[string]interface{}{"Id": registroMetasAsociadas["MetaId"]},
			"Activo":                      	registroMetasAsociadas["Activo"],
		}
		error2 := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Metas_Asociadas/"+idStr, "PUT", &MetasAsociadasPut, MetasAsociadasActualizar)
		if error2 != nil {
			return nil, error2
		} else {
			m := ExtraerDataPeticion(MetasAsociadasPut)
			return m, nil
		}
	}
}

//RegistroMetasAsociadasValidacion Valida si se requiere actualizar campos , regresa false en caso que se requiera actualizar
// func RegistroMetasAsociadasValidacion(registroMetasAsociadas map[string]interface{}, RegistroMetasAsociadasAntiguo map[string]interface{}) (validacion bool) {
// 	registroMetasAsociadasActual := make(map[string]interface{})

// 	registroMetasAsociadasActual = map[string]interface{}{
// 		"FechaCreacion":        RegistroMetasAsociadasAntiguo["FechaCreacion"],
// 		"MetaId": 				RegistroMetasAsociadasAntiguo["MetaId"]["Id"],
// 		"Activo":               RegistroMetasAsociadasAntiguo["Activo"],
// 	}

// 	if reflect.DeepEqual(registroMetasAsociadasActual["MetaId"]["Id"], registroMetasAsociadas["MetaId"]["Id"]) &&
// 		reflect.DeepEqual(registroMetasAsociadasActual["Activo"], registroMetasAsociadas["Activo"]) {
// 		return true
// 	} else {
// 		return false
// 	}

// }
