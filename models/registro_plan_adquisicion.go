package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerRegistroPlanAdquisicion regresa todos los registros de plan de adquisicion
func ObtenerRegistroPlanAdquisicion() (registroPlanAdquisicion []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicion []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/", &RegistroPlanAdquisicion)

	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicion, nil
	}

}

//ObtenerRegistroPlanAdquisicionByIDplan regresa un registro del plan de adquisicion segun su ID
func ObtenerRegistroPlanAdquisicionByIDplan(planAdquisicionID string) (registroPlanAdquisicion map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicion []map[string]interface{}
	var rubro []map[string]interface{}
	var unicos []string
	FuentesRegistroPlanAdquisicion := make(map[string]interface{})
	query := "PlanAdquisicionesId:" + planAdquisicionID + "&sortby=RubroId&order=asc"
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/?query="+query, &RegistroPlanAdquisicion)

	if error != nil {
		return nil, error
	} else {

		for rubroindex := range RegistroPlanAdquisicion {
			delete(RegistroPlanAdquisicion[rubroindex], "PlanAdquisicionesId")
			fuentes, errFuente := SeparaFuentes(RegistroPlanAdquisicion[rubroindex]["RubroId"])
			if errFuente != nil {
				return nil, errFuente
			}
			newfuente := stringInSlice(fuentes, unicos)
			if !newfuente {
				unicos = append(unicos, fuentes)
				rubro = make([]map[string]interface{}, 0)
			}
			rubro = append(rubro, RegistroPlanAdquisicion[rubroindex])
			FuentesRegistroPlanAdquisicion["Rubro: "+fuentes] = rubro
		}
		return FuentesRegistroPlanAdquisicion, nil
	}
}

//IngresoPlanAdquisicion crea un registro de plan de adquisicion
func IngresoPlanAdquisicion(registroPlanAdquisicion map[string]interface{}) (registroPlanAdquisicionRespuesta []map[string]interface{}, outputError interface{}) {
	registroPlanAdquisicionIngresado := make(map[string]interface{})
	registroPlanAdquisicionPost := make(map[string]interface{})
	result := []map[string]interface{}{}

	registroPlanAdquisicionIngresado = map[string]interface{}{
		"AreaFuncional":       registroPlanAdquisicion["AreaFuncional"],
		"CentroGestor":        registroPlanAdquisicion["CentroGestor"],
		"ResponsableId":       registroPlanAdquisicion["ResponsableId"],
		"MetaId":              registroPlanAdquisicion["MetaId"],
		"ProductoId":          registroPlanAdquisicion["ProductoId"],
		"RubroId":             registroPlanAdquisicion["RubroId"],
		"FechaEstimadaInicio": registroPlanAdquisicion["FechaEstimadaInicio"],
		"FechaEstimadaFin":    registroPlanAdquisicion["FechaEstimadaFin"],
		"Activo":              registroPlanAdquisicion["Activo"],
		"PlanAdquisicionesId": map[string]interface{}{"Id": registroPlanAdquisicion["PlanAdquisicionesId"]},
	}
	ModalidadSeleccion := registroPlanAdquisicion["ModalidadSeleccion"].([]interface{})
	CodigoArka := registroPlanAdquisicion["CodigoArka"].([]interface{})
	PlanAdquisicionActividad := registroPlanAdquisicion["RegistroPlanAdquisicionActividad"].([]interface{})

	errorSuma := SumaFuenteFinanciamiento(PlanAdquisicionActividad, registroPlanAdquisicion["RubroId"].(string))
	if errorSuma != nil {
		return nil, errorSuma
	}

	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/", "POST", &registroPlanAdquisicionPost, registroPlanAdquisicionIngresado)
	if error != nil {
		return nil, nil
	} else {
		result = append(result, registroPlanAdquisicionPost)
		DatosModalidadSeleccion, errorModalidad := GuardarModalidadSeleccion(ModalidadSeleccion, registroPlanAdquisicionPost["Id"])
		if errorModalidad != nil {
			return nil, errorModalidad
		} else {
			result = append(result, DatosModalidadSeleccion...)
			DatosCodigoArka, errorCodigoArka := GuardarCodigoArka(CodigoArka, registroPlanAdquisicionPost["Id"])
			if errorCodigoArka != nil {
				return nil, errorCodigoArka
			} else {
				result = append(result, DatosCodigoArka...)
				DatosPlanAdquisicionActividad, errorActividad := GuardarPlanAdquisicionActividad(PlanAdquisicionActividad, registroPlanAdquisicionPost["Id"])
				if errorActividad != nil {
					return nil, errorActividad
				} else {
					result = append(result, DatosPlanAdquisicionActividad...)
				}

			}
		}

		return result, nil
	}

}

//ObtenerRenglonRegistroPlanAdquisicionByID regresa un renglon segun el id del registro de plan de adquisicion
func ObtenerRenglonRegistroPlanAdquisicionByID(idStr string) (renglonRegistroPlanAdquisicion []map[string]interface{}, outputError interface{}) {
	var RenglonRegistroPlanAdquisicion []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/?query=Id%3A"+idStr, &RenglonRegistroPlanAdquisicion)
	if error != nil {
		return nil, error
	} else {
		if len(RenglonRegistroPlanAdquisicion) == 1 && len(RenglonRegistroPlanAdquisicion[0]) == 0 {
			error := "No existe Registro Plan Adquisicion"
			return nil, error
		}
		CodigoArka, error := ObtenerRegistroCodigoArkaByIDPlanAdquisicion(idStr)
		if error != nil {
			return nil, error
		} else {
			ModalidadSeleccion, error := ObtenerRegistroModalidadSeleccionByIDPlanAdquisicion(idStr)
			if error != nil {
				return nil, error
			} else {
				RegistroPlanAdquisicionActividad, error := ObtenerRegistroTablaActividades(idStr)
				if error != nil {
					return nil, error
				} else {
					EliminarCampos(CodigoArka, "RegistroPlanAdquisicionesId")
					EliminarCampos(ModalidadSeleccion, "RegistroPlanAdquisicionesId")
					RenglonRegistroPlanAdquisicion[0]["CodigoArka"] = CodigoArka
					RenglonRegistroPlanAdquisicion[0]["ModalidadSeleccion"] = ModalidadSeleccion
					RenglonRegistroPlanAdquisicion[0]["RegistroPlanAdquisicionActividad"] = RegistroPlanAdquisicionActividad
				}
			}
		}
		return RenglonRegistroPlanAdquisicion, nil
	}

}

//ActualizarRegistroPlanAdquisicion verifica y actualiza los campos de un renglon segun el ID de un registro plan de adquisicion
func ActualizarRegistroPlanAdquisicion(registroPlanAdquisicion map[string]interface{}, idStr string) (registroActividadRespuesta map[string]interface{}, outputError interface{}) {
	registroPlanAdquisicionPut := make(map[string]interface{})
	registroPlanAdquisicionActualizar := make(map[string]interface{})
	RegistroPlanAdquisicionAntiguo, error := ObtenerRenglonRegistroPlanAdquisicionByID(idStr)
	if error != nil && error != "No existe Registro Plan Adquisicion" {
		return nil, error
	} else {
		if error == "No existe Registro Plan Adquisicion" {
			//fmt.Println("No existe Registro Plan Adquisicion, toca crearlo")
			_, errRegistroPlanAdquisicion := IngresoPlanAdquisicion(registroPlanAdquisicion)
			if errRegistroPlanAdquisicion != nil {
				return nil, errRegistroPlanAdquisicion
			} else {
				return registroPlanAdquisicion, nil
			}

		} else {
			validacion := RegistroPlanAdquisicionModificado(registroPlanAdquisicion, RegistroPlanAdquisicionAntiguo[0], idStr)
			if validacion {
				//fmt.Println("existe registro Plan Adquisicion y no toca modificarlo")
				PlanAdquisicionActividad := registroPlanAdquisicion["RegistroPlanAdquisicionActividad"].([]interface{})
				errorSuma := SumaFuenteFinanciamiento(PlanAdquisicionActividad, registroPlanAdquisicion["RubroId"].(string))
				if errorSuma != nil {
					return nil, errorSuma
				}

				error := CodigoArkaModificado(registroPlanAdquisicion, idStr)
				if error != nil {
					return nil, error
				} else {
					error := ModalidadSeleccionModificado(registroPlanAdquisicion, idStr)
					if error != nil {
						return nil, error
					} else {
						error := RegistroActividadModificado(registroPlanAdquisicion, idStr)
						if error != nil {
							return nil, error
						} else {
							return registroPlanAdquisicion, nil
						}
					}
				}
			} else {
				//fmt.Println("existe registro y  toca modificarlo")
				PlanAdquisicionActividad := registroPlanAdquisicion["RegistroPlanAdquisicionActividad"].([]interface{})
				errorSuma := SumaFuenteFinanciamiento(PlanAdquisicionActividad, registroPlanAdquisicion["RubroId"].(string))
				if errorSuma != nil {
					return nil, errorSuma
				}

				registroPlanAdquisicionActualizar = map[string]interface{}{
					"AreaFuncional":       registroPlanAdquisicion["AreaFuncional"],
					"CentroGestor":        registroPlanAdquisicion["CentroGestor"],
					"ResponsableId":       registroPlanAdquisicion["ResponsableId"],
					"MetaId":              registroPlanAdquisicion["MetaId"],
					"ProductoId":          registroPlanAdquisicion["ProductoId"],
					"RubroId":             registroPlanAdquisicion["RubroId"],
					"FechaCreacion":       RegistroPlanAdquisicionAntiguo[0]["FechaCreacion"],
					"FechaEstimadaInicio": registroPlanAdquisicion["FechaEstimadaInicio"],
					"FechaEstimadaFin":    registroPlanAdquisicion["FechaEstimadaFin"],
					"Activo":              registroPlanAdquisicion["Activo"],
					"PlanAdquisicionesId": map[string]interface{}{"Id": registroPlanAdquisicion["PlanAdquisicionesId"]},
				}

				error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/"+idStr, "PUT", &registroPlanAdquisicionPut, registroPlanAdquisicionActualizar)
				if error != nil {
					return nil, error
				} else {
					error := CodigoArkaModificado(registroPlanAdquisicion, idStr)
					if error != nil {
						return nil, error
					} else {
						error := ModalidadSeleccionModificado(registroPlanAdquisicion, idStr)
						if error != nil {
							return nil, error
						} else {
							error := RegistroActividadModificado(registroPlanAdquisicion, idStr)
							if error != nil {
								return nil, error
							} else {
								return registroPlanAdquisicion, nil
							}
						}
					}
				}

			}
		}
	}
}

//RegistroPlanAdquisicionModificado Valida si existen actualizaciones en los campos del registro de plan de adquisicion
func RegistroPlanAdquisicionModificado(registroPlanAdquisicion map[string]interface{}, RegistroPlanAdquisicionAntiguo map[string]interface{}, idStr string) (validacion bool) {
	registroPlanAdquisicionActual := make(map[string]interface{})

	registroPlanAdquisicionActual = map[string]interface{}{
		"AreaFuncional":       RegistroPlanAdquisicionAntiguo["AreaFuncional"],
		"CentroGestor":        RegistroPlanAdquisicionAntiguo["CentroGestor"],
		"ResponsableId":       RegistroPlanAdquisicionAntiguo["ResponsableId"],
		"MetaId":              RegistroPlanAdquisicionAntiguo["MetaId"],
		"ProductoId":          RegistroPlanAdquisicionAntiguo["ProductoId"],
		"RubroId":             RegistroPlanAdquisicionAntiguo["RubroId"],
		"FechaCreacion":       RegistroPlanAdquisicionAntiguo["FechaCreacion"],
		"FechaEstimadaInicio": RegistroPlanAdquisicionAntiguo["FechaEstimadaInicio"],
		"FechaEstimadaFin":    RegistroPlanAdquisicionAntiguo["FechaEstimadaFin"],
		"Activo":              RegistroPlanAdquisicionAntiguo["Activo"],
		"PlanAdquisicionesId": map[string]interface{}{"Id": RegistroPlanAdquisicionAntiguo["PlanAdquisicionesId"]},
	}

	id := registroPlanAdquisicionActual["PlanAdquisicionesId"].(map[string]interface{})
	idRegistroPlanAdquisicion := id["Id"].(map[string]interface{})

	if reflect.DeepEqual(idRegistroPlanAdquisicion["Id"], registroPlanAdquisicion["PlanAdquisicionesId"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["AreaFuncional"], registroPlanAdquisicion["AreaFuncional"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["CentroGestor"], registroPlanAdquisicion["CentroGestor"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["ResponsableId"], registroPlanAdquisicion["ResponsableId"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["MetaId"], registroPlanAdquisicion["MetaId"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["ProductoId"], registroPlanAdquisicion["ProductoId"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["RubroId"], registroPlanAdquisicion["RubroId"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["FechaEstimadaInicio"], registroPlanAdquisicion["FechaEstimadaInicio"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["FechaEstimadaFin"], registroPlanAdquisicion["FechaEstimadaFin"]) &&
		reflect.DeepEqual(registroPlanAdquisicionActual["Activo"], registroPlanAdquisicion["Activo"]) {
		return true
	} else {
		return false
	}

}

//EliminarCampos eliminar campos que no se quieran ver en el JSON
func EliminarCampos(mapa []map[string]interface{}, campo string) {
	for index := range mapa {
		delete(mapa[index], campo)
	}

}

//SeparaFuentes separa el id del rubro y su fuente
func SeparaFuentes(RubroRegistroPlanAdquisicion interface{}) (string, interface{}) {
	str := MapToString(RubroRegistroPlanAdquisicion)
	fuente := strings.Split(str, "-")
	if len(fuente) < 2 {
		error := "No existe Plan de adquisicion"
		return "", error
	}
	fuentes := fuente[0] + fuente[1]
	return fuentes, nil
}

//MapToString convierte un MAP a string
func MapToString(RubroRegistroPlanAdquisicion interface{}) string {
	str := fmt.Sprintf("%v", RubroRegistroPlanAdquisicion)
	return str
}

// stringInSlice regresa true/false si se repite un elemento
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
