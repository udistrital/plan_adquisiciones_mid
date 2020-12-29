package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//IngresoRegistroCodigoArka ingresar elemento en la tabla de Registro_Plan_adquisiciones_codigo_arka
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

//GuardarCodigoArka descompone array para ingresar uno a uno los registro codigo arka
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

//ObtenerRegistroCodigoArkaByIDPlanAdquisicion regresa registro codigo arka segun ID de registro plan adquisiciones
func ObtenerRegistroCodigoArkaByIDPlanAdquisicion(idStr string) (CodigoArka []map[string]interface{}, outputError interface{}) {
	var codigoArka []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Codigo_arka/?query=RegistroPlanAdquisicionesId.id:"+idStr+",Activo:true", &codigoArka)
	if error != nil {
		return nil, error
	} else {
		for index := range codigoArka {
			ElementoCodigoArka, error := CatalogoElementosArka(codigoArka[index]["CodigoArka"].(string))
			if error != nil {
				return nil, error
			} else {
				codigoArka[index]["Descripcion"] = ElementoCodigoArka["Descripcion"]
			}
		}
		return codigoArka, nil
	}

}

//ObtenerRegistroCodigoArkaByID regresa registro codigo arka segun ID
func ObtenerRegistroCodigoArkaByID(idStr string) (CodigoArka map[string]interface{}, outputError interface{}) {
	var codigoArka []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Codigo_arka/?query=id:"+idStr, &codigoArka)
	if error != nil {
		return nil, error
	} else {
		return codigoArka[0], nil
	}

}

//CodigoArkaModificado descompone array para actualizar codigo arka uno a uno
func CodigoArkaModificado(registroPlanAdquisicion map[string]interface{}, idStr string) (outputError interface{}) {
	CodigosArka := registroPlanAdquisicion["CodigoArka"].([]interface{})
	for Index := range CodigosArka {
		CodigoArka := CodigosArka[Index].(map[string]interface{})
		_, errRegistroCodigoArka := ActualizarRegistroCodigoArka(CodigoArka, fmt.Sprintf("%v", CodigoArka["Id"]), idStr)
		if errRegistroCodigoArka != nil {
			return errRegistroCodigoArka
		}
	}
	return nil
}

//ActualizarRegistroCodigoArka actualiza un registro de codigo arka y si no existe lo crea
func ActualizarRegistroCodigoArka(registroCodigoArka map[string]interface{}, idStr string, idStrPlanAdquisicion string) (registroCodigoArkaRespuesta map[string]interface{}, outputError interface{}) {
	CodigoArkaPut := make(map[string]interface{})
	CodigoArkaActualizar := make(map[string]interface{})
	RegistroCodigoArkaAntiguo, error := ObtenerRegistroCodigoArkaByID(idStr)
	if error != nil {
		return nil, error
	} else {
		if len(RegistroCodigoArkaAntiguo) == 0 {
			//fmt.Println("No existe CodigoARKA toca crearlo")
			idint, _ := strconv.Atoi(idStrPlanAdquisicion)
			registroCodigoArka["RegistroPlanAdquisicionesId"] = idint
			RegistroRegistroCodigoArka, errRegistroRegistroCodigoArka := IngresoRegistroCodigoArka(registroCodigoArka)
			if errRegistroRegistroCodigoArka != nil {
				return nil, errRegistroRegistroCodigoArka
			} else {
				return RegistroRegistroCodigoArka, nil
			}

		} else {
			validacion := RegistroCodigoArkaValidacion(registroCodigoArka, RegistroCodigoArkaAntiguo)
			if validacion {
				//fmt.Println("existe Codigo ARKA No toca modificar")
			} else {
				//fmt.Println("existe Codigo ARKA toca modificar")
				idint, _ := strconv.Atoi(idStrPlanAdquisicion)
				registroCodigoArka["RegistroPlanAdquisicionesId"] = idint
				CodigoArkaActualizar = map[string]interface{}{
					"RegistroPlanAdquisicionesId": map[string]interface{}{"Id": registroCodigoArka["RegistroPlanAdquisicionesId"]},
					"FechaCreacion":               RegistroCodigoArkaAntiguo["FechaCreacion"],
					"CodigoArka":                  registroCodigoArka["CodigoArka"],
					"Activo":                      registroCodigoArka["Activo"],
				}
				error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones-Codigo_arka/"+idStr, "PUT", &CodigoArkaPut, CodigoArkaActualizar)
				if error != nil {
					return nil, error
				} else {
					return CodigoArkaPut, nil
				}
			}
		}

		return CodigoArkaPut, nil
	}

}

//RegistroCodigoArkaValidacion valida si se debe modificar campos del codigo arka
func RegistroCodigoArkaValidacion(registroCodigoArka map[string]interface{}, RegistroCodigoArkaAntiguo map[string]interface{}) (validacion bool) {
	registroCodigoArkaActual := make(map[string]interface{})

	registroCodigoArkaActual = map[string]interface{}{
		"FechaCreacion": RegistroCodigoArkaAntiguo["FechaCreacion"],
		"CodigoArka":    RegistroCodigoArkaAntiguo["CodigoArka"],
		"Activo":        RegistroCodigoArkaAntiguo["Activo"],
	}

	if reflect.DeepEqual(registroCodigoArkaActual["CodigoArka"], registroCodigoArka["CodigoArka"]) &&
		reflect.DeepEqual(registroCodigoArkaActual["Activo"], registroCodigoArka["Activo"]) {
		return true
	} else {
		return false
	}

}

//CatalogoElementosArka Consulta nombre en el catalogo de elementos de arka
func CatalogoElementosArka(idStr string) (NombreElemento map[string]interface{}, outputError interface{}) {
	var ElementoCodigoArka []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("catalogo_elementos_arka_url")+"elemento?fields=Id,Descripcion&query=Id:"+idStr, &ElementoCodigoArka)
	if error != nil {
		return nil, error
	} else {
		return ElementoCodigoArka[0], nil
	}

}
