package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//IngresoRegistroProductosAsociados ingresa un elemento a la tabla modalid de seleccion
func IngresoRegistroProductosAsociados(registroProductosAsociados map[string]interface{}) (registroProductosAsociadosRespuesta map[string]interface{}, outputError interface{}) {
	registroProductosAsociadosIngresado := make(map[string]interface{})
	registroProductosAsociadosPost := make(map[string]interface{})

	registroProductosAsociadosIngresado = map[string]interface{}{
		"RegistroPlanAdquisicionesId": 	map[string]interface{}{"Id": registroProductosAsociados["RegistroPlanAdquisicionesId"]},
		"ProductoAsociadoId":        	registroProductosAsociados["ProductoAsociadoId"],
		"PorcentajeDistribucion":      	registroProductosAsociados["PorcentajeDistribucion"],
		"Activo":                      	registroProductosAsociados["Activo"],
	}
	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Productos_Asociados/", "POST", &registroProductosAsociadosPost, registroProductosAsociadosIngresado)
	if error != nil {
		return nil, error
	} else {
		m := ExtraerDataPeticion(registroProductosAsociadosPost)
		return m, nil
	}

}

//GuardarProductosAsociados descompone el array de modalidad de selección para crear uno a uno
func GuardarProductosAsociados(ProductosAsociados []interface{}, idPost interface{}) (registroProductosAsociadosRespuesta []map[string]interface{}, outputError interface{}) {
	resultModalidad := make([]map[string]interface{}, 0)
	for Index := range ProductosAsociados {
		ProductosAsociados := ProductosAsociados[Index].(map[string]interface{})
		ProductosAsociados["RegistroPlanAdquisicionesId"] = idPost
		RegistroProductosAsociados, errRegistroProductosAsociados := IngresoRegistroProductosAsociados(ProductosAsociados)
		if errRegistroProductosAsociados != nil {
			return nil, errRegistroProductosAsociados
		} else {
			resultModalidad = append(resultModalidad, RegistroProductosAsociados)
		}
	}
	return resultModalidad, nil
}

//ObtenerRegistroProductosAsociadosByIDPlanAdquisicion regresa una registro de la tabla modalidad de seleccioón segun un Id de un registro_plan_adquisicion
func ObtenerRegistroProductosAsociadosByIDPlanAdquisicion(idStr string) (ProductosAsociados []map[string]interface{}, outputError interface{}) {
	var productosAsociados map[string]interface{}
	// var nombreProductosAsociados []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Productos_Asociados/?query=RegistroPlanAdquisicionesId.Id:"+idStr+",Activo:true&fields=Id,ProductoAsociadoId,Activo,FechaCreacion,FechaModificacion", &productosAsociados)
	if error != nil {
		return nil, error
	} else {
		m := ExtraerDataPeticionArreglo(productosAsociados)
		for index := range m {
			producto, error := ObtenerProductoByID(MapToString(m[index]["ProductoAsociadoId"]))
			if error != nil {
				return nil, error
			} else {
				m[index]["ProductoData"] = producto
			}
		}
		return m, nil
	}

}

//ObtenerRegistroProductosAsociadosByID regresa una registro de la tabla modalidad de seleccioón segun el ID
func ObtenerRegistroProductosAsociadosByID(idStr string) (ProductosAsociados map[string]interface{}, outputError interface{}) {
	var productosAsociados map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Productos_Asociados/?query=Id:"+idStr+"&fields=Id,ProductoAsociadoId,Activo,FechaCreacion,FechaModificacion", &productosAsociados)
	if error != nil {
		return nil, error
	} else {
		// fmt.Println(productosAsociados["Data"])
		m := ExtraerDataPeticionArreglo(productosAsociados)
		return m[0], nil
	}

}

//ProductosAsociadosModificado descompone el array de modalidad de selección para actualizar uno a uno
func ProductosAsociadosModificado(registroPlanAdquisicion map[string]interface{}, idStr string) (outputError interface{}) {
	ProductosAsociados := registroPlanAdquisicion["ProductosAsociados"].([]interface{})
	for Index := range ProductosAsociados {
		ProductosAsociados := ProductosAsociados[Index].(map[string]interface{})
		_, errRegistroProductosAsociados := ActualizarRegistroProductosAsociados(ProductosAsociados, fmt.Sprintf("%v", ProductosAsociados["Id"]), idStr)
		if errRegistroProductosAsociados != nil {
			return errRegistroProductosAsociados
		}
	}
	return nil
}

//ActualizarRegistroProductosAsociados Actualiza la modalidad de selección y la crea en caso de que no exista
func ActualizarRegistroProductosAsociados(registroProductosAsociados map[string]interface{}, idStr string, idStrPlanAdquisicion string) (registroProductosAsociadosRespuesta map[string]interface{}, outputError interface{}) {
	ProductosAsociadosPut := make(map[string]interface{})
	ProductosAsociadosActualizar := make(map[string]interface{})
	

	if registroProductosAsociados["Id"].(float64) == 0 {
		//fmt.Println("No existe ProductosAsociados toca crearlo")
		idint, _ := strconv.Atoi(idStrPlanAdquisicion)
		registroProductosAsociados["RegistroPlanAdquisicionesId"] = idint
		RegistroProductosAsociados, errRegistroProductosAsociados := IngresoRegistroProductosAsociados(registroProductosAsociados)
		if errRegistroProductosAsociados != nil {
			return nil, errRegistroProductosAsociados
		} else {
			return RegistroProductosAsociados, nil
		}

	} else {
		// fmt.Println(idStr)
		RegistroProductosAsociadosAntiguo, _ := ObtenerRegistroProductosAsociadosByID(idStr)
		//fmt.Println("existe ProductosAsociados toca modificar")
		idint, _ := strconv.Atoi(idStrPlanAdquisicion)
		registroProductosAsociados["RegistroPlanAdquisicionesId"] = idint
		ProductosAsociadosActualizar = map[string]interface{}{
			"RegistroPlanAdquisicionesId": 	map[string]interface{}{"Id": registroProductosAsociados["RegistroPlanAdquisicionesId"]},
			"FechaCreacion":               	RegistroProductosAsociadosAntiguo["FechaCreacion"],
			"ProductoAsociadoId":     	   	registroProductosAsociados["ProductoAsociadoId"],
			"PorcentajeDistribucion":		registroProductosAsociados["PorcentajeDistribucion"],
			"Activo":                      	registroProductosAsociados["Activo"],
		}
		error2 := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Productos_Asociados/"+idStr, "PUT", &ProductosAsociadosPut, ProductosAsociadosActualizar)
		if error2 != nil {
			return nil, error2
		} else {
			m := ExtraerDataPeticion(ProductosAsociadosPut)
			return m, nil
		}
	}
}

//RegistroProductosAsociadosValidacion Valida si se requiere actualizar campos , regresa false en caso que se requiera actualizar
func RegistroProductosAsociadosValidacion(registroProductosAsociados map[string]interface{}, RegistroProductosAsociadosAntiguo map[string]interface{}) (validacion bool) {
	registroProductosAsociadosActual := make(map[string]interface{})

	registroProductosAsociadosActual = map[string]interface{}{
		"FechaCreacion":        RegistroProductosAsociadosAntiguo["FechaCreacion"],
		"IdProductosAsociados": RegistroProductosAsociadosAntiguo["IdProductosAsociados"],
		"Activo":               RegistroProductosAsociadosAntiguo["Activo"],
	}

	if reflect.DeepEqual(registroProductosAsociadosActual["IdProductosAsociados"], registroProductosAsociados["IdProductosAsociados"]) &&
		reflect.DeepEqual(registroProductosAsociadosActual["Activo"], registroProductosAsociados["Activo"]) {
		return true
	} else {
		return false
	}

}

func ExtraerDataPeticion(Peticion map[string]interface{}) (Data map[string]interface{}) {
	if Peticion["Data"] == nil {
		m := make(map[string]interface{})
		return m
	}
	m := Peticion["Data"].(interface{})
	Datos := m.(map[string]interface{})
	return Datos
}

func ExtraerDataPeticionArreglo(Peticion map[string]interface{}) (Data []map[string]interface{}) {
	if Peticion["Data"] == nil {
		m := make([]map[string]interface{},0)
		return m
	} else {
		m := Peticion["Data"].([]interface{})
		x:= make([]map[string]interface{},0)
		if len(m[0].(map[string]interface{})) == 0 {
			return make([]map[string]interface{},0)
		} else {
			for index := range m {
				x = append(x, m[index].(map[string]interface{}))
			}
			return x
		}
		
	}
}
