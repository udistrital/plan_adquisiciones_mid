package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerRegistroPlanInversionActividadFuente regresa los elementos de la tabla registros_inversion_actividad-fuente_financiamiento
func ObtenerRegistroPlanInversionActividadFuente() (registroPlanAdquisicionActividadFuente []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividadFuente []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/", &RegistroPlanAdquisicionActividadFuente)

	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividadFuente, nil
	}

}

//IngresoRegistroPlanInversionActividadFuente ingresa un elemento a la tabla los registros_inversion_actividad-fuente_financiamiento
func IngresoRegistroPlanInversionActividadFuente(registroActividadFuente map[string]interface{}) (registroActividadFuenteRespuesta map[string]interface{}, outputError interface{}) {
	registroActividadFuenteIngresado := make(map[string]interface{})
	registroActividadFuentePost := make(map[string]interface{})

	registroActividadFuenteIngresado = map[string]interface{}{
		"RegistroPlanAdquisicionesActividadId": map[string]interface{}{"Id": registroActividadFuente["ActividadId"]},
		"ValorAsignado":                        registroActividadFuente["ValorAsignado"],
		"FuenteFinanciamientoId":               registroActividadFuente["FuenteFinanciamientoId"],
		"Activo":                               registroActividadFuente["Activo"],
	}
	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/", "POST", &registroActividadFuentePost, registroActividadFuenteIngresado)
	if error != nil {
		return nil, error
	} else {
		return registroActividadFuentePost, nil
	}

}

//ActualizarRegistroActividadFuente actualiza un elemento de registro inversion fuente financiamiento, si no existe lo crea
func ActualizarRegistroActividadFuente(registroActividadFuente map[string]interface{}, idStr string, idStrActividad string) (registroActividadFuenteRespuesta map[string]interface{}, outputError interface{}) {
	fuenteActividadPut := make(map[string]interface{})
	fuenteActividadActualizar := make(map[string]interface{})
	RegistroPlanAdquisicionActividadFuente, error := ObtenerRegistroPlanAdquisicionActividadFuenteByID(idStr)
	if error != nil {
		return nil, error
	} else {
		if len(RegistroPlanAdquisicionActividadFuente) == 0 {
			//fmt.Println("No existe fuente toca crearla")
			idint, _ := strconv.Atoi(idStrActividad)
			registroActividadFuente["ActividadId"] = idint
			RegistroActividadFuente, errRegistroActividadFuente := IngresoRegistroPlanInversionActividadFuente(registroActividadFuente)
			if errRegistroActividadFuente != nil {
				return nil, errRegistroActividadFuente
			} else {
				return RegistroActividadFuente, nil
			}

		} else {
			validacion := RegistroFuenteModificado(registroActividadFuente, RegistroPlanAdquisicionActividadFuente)
			if validacion {
				//fmt.Println("existe fuente No toca modificar")
			} else {
				//fmt.Println("existe fuente toca modificar")
				idint, _ := strconv.Atoi(idStrActividad)
				registroActividadFuente["ActividadId"] = idint
				fuenteActividadActualizar = map[string]interface{}{
					"RegistroPlanAdquisicionesActividadId": map[string]interface{}{"Id": registroActividadFuente["ActividadId"]},
					"ValorAsignado":                        registroActividadFuente["ValorAsignado"],
					"FechaCreacion":                        RegistroPlanAdquisicionActividadFuente["FechaCreacion"],
					"FuenteFinanciamientoId":               registroActividadFuente["FuenteFinanciamientoId"],
					"Activo":                               registroActividadFuente["Activo"],
				}
				error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/"+idStr, "PUT", &fuenteActividadPut, fuenteActividadActualizar)
				if error != nil {
					return nil, error
				} else {
					return fuenteActividadPut, nil
				}
			}
		}

		return RegistroPlanAdquisicionActividadFuente, nil
	}

}

//ObtenerRegistroPlanAdquisicionActividadFuenteByID obtener un elemento segun el ID del registro inversion fuente financiamiento
func ObtenerRegistroPlanAdquisicionActividadFuenteByID(idStr string) (registroPlanAdquisicionActividadFuente map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividadFuente []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/?query=Id:"+idStr, &RegistroPlanAdquisicionActividadFuente)
	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividadFuente[0], nil
	}

}

//ObtenerRegistroTablaActividades regresa una tabla ordenada del registro de actividades con sus fuentes de financiamiento
func ObtenerRegistroTablaActividades(idStr string) (registroPlanAdquisicionActividadFuente []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividadFuente []map[string]interface{}
	var unicos []string
	registro := make(map[string]interface{})
	registros := make([]map[string]interface{}, 0)
	fuentesFinanciamiento := make([]map[string]interface{}, 0)
	query := "?query=RegistroPlanAdquisicionesActividadId.RegistroPlanAdquisicionesId.Id:" + idStr + ",RegistroPlanAdquisicionesActividadId.Activo:true,Activo:true&sortby=RegistroPlanAdquisicionesActividadId__Id&order=asc"
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/"+query, &RegistroPlanAdquisicionActividadFuente)

	if error != nil {
		return nil, error
	} else {
		if len(RegistroPlanAdquisicionActividadFuente) == 1 {
			if len(RegistroPlanAdquisicionActividadFuente[0]) == 0 {
				error := "No existe plan adquisicion"
				return nil, error
			}
		}
		Vigencia, AreaFuncional, error := VigenciaYAreaFuncional(idStr)
		if error != nil {
			return nil, error
		}
		for index := range RegistroPlanAdquisicionActividadFuente {
			ActividadID := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["ActividadId"].(map[string]interface{})["Id"]
			ValorActividad := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["Valor"]
			RegistroActividadID := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["Id"]
			ActivoActividad := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["Activo"]
			NombreActividad := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["ActividadId"].(map[string]interface{})["Nombre"]
			FechaCreacionActividad := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["FechaCreacion"]
			FechaModificacionActividad := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["FechaModificacion"]
			newdata := stringInSlice(fmt.Sprintf("%.0f", RegistroActividadID.(float64)), unicos)
			if !newdata {
				unicos = append(unicos, fmt.Sprintf("%.0f", RegistroActividadID.(float64)))
				registros = append(registros, registro)
				fuentesFinanciamiento = make([]map[string]interface{}, 0)
			}

			Fuente, errorFuente := ObtenerFuenteFinanciamientoByCodigo(RegistroPlanAdquisicionActividadFuente[index]["FuenteFinanciamientoId"].(string), Vigencia, AreaFuncional)
			if errorFuente != nil {
				return nil, errorFuente
			}
			fuenteFinanciamiento := map[string]interface{}{
				"Id":                   RegistroPlanAdquisicionActividadFuente[index]["Id"],
				"ValorAsignado":        RegistroPlanAdquisicionActividadFuente[index]["ValorAsignado"],
				"Activo":               RegistroPlanAdquisicionActividadFuente[index]["Activo"],
				"FuenteFinanciamiento": RegistroPlanAdquisicionActividadFuente[index]["FuenteFinanciamientoId"],
				"FechaCreacion":        RegistroPlanAdquisicionActividadFuente[index]["FechaCreacion"],
				"FechaModificacion":    RegistroPlanAdquisicionActividadFuente[index]["FechaModificacion"],
				"Nombre":               Fuente["Nombre"],
			}
			fuentesFinanciamiento = append(fuentesFinanciamiento, fuenteFinanciamiento)

			registro = map[string]interface{}{
				"ActividadId":                 ActividadID,
				"Nombre":                      NombreActividad,
				"RegistroPlanAdquisicionesId": idStr,
				"Valor":                       ValorActividad,
				"Activo":                      ActivoActividad,
				"RegistroActividadId":         RegistroActividadID,
				"FuentesFinanciamiento":       fuentesFinanciamiento,
				"FechaCreacion":               FechaCreacionActividad,
				"FechaModificacion":           FechaModificacionActividad,
			}

		}
		registros = append(registros[1:], registro)
		return registros, nil
	}

}

//RegistroFuenteModificado valida si algun campo de la fuente de financiamiento fue modificado
func RegistroFuenteModificado(registroFuente map[string]interface{}, RegistroPlanAdquisicionActividadFuente map[string]interface{}) (validacion bool) {
	registroFuenteActual := make(map[string]interface{})

	registroFuenteActual = map[string]interface{}{
		"ValorAsignado":          RegistroPlanAdquisicionActividadFuente["ValorAsignado"],
		"FechaCreacion":          RegistroPlanAdquisicionActividadFuente["FechaCreacion"],
		"FuenteFinanciamientoId": RegistroPlanAdquisicionActividadFuente["FuenteFinanciamientoId"],
		"Activo":                 RegistroPlanAdquisicionActividadFuente["Activo"],
	}

	if reflect.DeepEqual(registroFuenteActual["ValorAsignado"], registroFuente["ValorAsignado"]) && reflect.DeepEqual(registroFuenteActual["Activo"], registroFuente["Activo"]) && reflect.DeepEqual(registroFuenteActual["FuenteFinanciamientoId"], registroFuente["FuenteFinanciamientoId"]) {
		return true
	} else {
		return false
	}

}

//SumaFuenteFinanciamiento regresa la suma de todas las fuentes de financimiento Antes de realizar un POST o PUT de un renglon
func SumaFuenteFinanciamiento(PlanAdquisicionActividades []interface{}, IDRubro string, Vigencia string, AreaFuncional string) (outputError interface{}) {
	var RubroMongo map[string]interface{}
	var valor float64
	error := request.GetJson(beego.AppConfig.String("plan_cuentas_mongo_crud_url")+"arbol_rubro_apropiacion/"+IDRubro+"/"+Vigencia+"/"+AreaFuncional+"/", &RubroMongo)
	if error != nil {
		return error
	} else {
		m := RubroMongo["Body"].(interface{})
		ValorActualRubro := m.(map[string]interface{})["ValorActual"].(float64)
		for Index := range PlanAdquisicionActividades {
			PlanAdquisicionActividad := PlanAdquisicionActividades[Index].(map[string]interface{})
			FuentesFinanciamiento := PlanAdquisicionActividad["FuentesFinanciamiento"].([]interface{})
			if PlanAdquisicionActividad["Activo"].(bool) {
				for fuenteIndex := range FuentesFinanciamiento {
					FuenteFinanciamiento := FuentesFinanciamiento[fuenteIndex].(map[string]interface{})
					if FuenteFinanciamiento["Activo"].(bool) {
						valor = valor + FuenteFinanciamiento["ValorAsignado"].(float64)
					}
				}
			}
		}
		if valor > ValorActualRubro {
			errorValorRubro := "La suma de las fuentes de financiamiento supera el valor actual del rubro"
			return errorValorRubro
		}
		return nil
	}
}

//ObtenerFuenteFinanciamientoByCodigo Trae campos de fuente financiamiento segun codigo
func ObtenerFuenteFinanciamientoByCodigo(Codigo string, Vigencia string, UnidadEjecutora string) (fuentefinanciamiento map[string]interface{}, outputError interface{}) {
	var FuenteFinanciamientoMongo map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_cuentas_mongo_crud_url")+"fuente_financiamiento/"+Codigo+"/"+Vigencia+"/"+UnidadEjecutora, &FuenteFinanciamientoMongo)
	if error != nil {
		return nil, error
	} else {
		if FuenteFinanciamientoMongo["Body"] == nil {
			error := "No se encontro fuente de financiamiento"
			return nil, error
		}
		m := FuenteFinanciamientoMongo["Body"].(interface{})
		FuenteFinanciamiento := m.(map[string]interface{})
		return FuenteFinanciamiento, nil
	}
}
