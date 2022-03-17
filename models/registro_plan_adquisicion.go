package models

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_adquisiciones_mid/helpers"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/movimientosCrud"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/utils"
	"github.com/udistrital/utils_oas/errorctrl"
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

//ObtenerRegistroPlanAdquisicionByIDplan regresa un registro del plan de adquisicion segun ID planADquisicion seprados por fuentes de financiamiento
func ObtenerRegistroPlanAdquisicionByIDplan(planAdquisicionID string) (PlanAdquisicion []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicion []map[string]interface{}
	registros := make([]map[string]interface{}, 0)
	// TODO: Dejar únicamente los campos necesarios para la consulta
	query := beego.AppConfig.String("plan_adquicisiones_crud_url") + "Registro_plan_adquisiciones?query=" +
		"PlanAdquisicionesId:" + planAdquisicionID + "&sortby=FechaCreacion,RubroId&order=asc"
	error := request.GetJson(query, &RegistroPlanAdquisicion)
	if error != nil {
		return nil, error
	} else {
		if len(RegistroPlanAdquisicion) == 0 {
			return RegistroPlanAdquisicion, nil
		} else {
			for index := range RegistroPlanAdquisicion {
				id := fmt.Sprintf("%.0f", RegistroPlanAdquisicion[index]["Id"].(float64))
				registro, error := ObtenerRenglonRegistroPlanAdquisicionByID(id)
				// logs.Debug(error)
				if error != nil {
					logs.Warning(error)
				} else {
					registros = append(registros, registro[0])
				}
			}
			// fmt.Println("registros: ", registros)
			FuentesRegistroPlanAdquisicion, error := SepararRegistrosPorFuente(registros)
			if error != nil {
				logs.Warning(error)
			} else {
				return FuentesRegistroPlanAdquisicion, nil
			}

			return
		}
	}
}

//SepararRegistrosPorFuente separa los registros del plan adquisicion por rubro
func SepararRegistrosPorFuente(RegistroPlanAdquisicion []map[string]interface{}) (registroSeparados []map[string]interface{}, outputError interface{}) {
	var fuente map[string]interface{}
	var unicos []string
	FuentesRegistroPlanAdquisicion := make([]map[string]interface{}, 0)
	rubrosSeparados, errorRubro := SepararRegistrosPorRubro(RegistroPlanAdquisicion)
	if errorRubro != nil {
		return nil, errorRubro
	}
	for rubroindex := range rubrosSeparados {
		fuentes, errFuente := SeparaFuentes(rubrosSeparados[rubroindex]["Rubro"])
		if errFuente != nil {
			return RegistroPlanAdquisicion, nil
		}
		newfuente := stringInSlice(fuentes, unicos)
		if !newfuente {
			FuenteData, error := ObtenerFuenteReducidaByID(fuentes)
			if error != nil {
				return nil, error
			} else {
				unicos = append(unicos, fuentes)
				fuente = map[string]interface{}{
					"Fuente":     fuentes,
					"FuenteData": FuenteData,
					"datos":      make([]map[string]interface{}, 0),
				}
				fuente["datos"] = append(fuente["datos"].([]map[string]interface{}), rubrosSeparados[rubroindex])
				FuentesRegistroPlanAdquisicion = append(FuentesRegistroPlanAdquisicion, fuente)
			}
		} else {
			index := BuscarIndexPorCampo(FuentesRegistroPlanAdquisicion, fuentes, "Fuente")
			if index != -1 {
				FuentesRegistroPlanAdquisicion[index]["datos"] = append(FuentesRegistroPlanAdquisicion[index]["datos"].([]map[string]interface{}), rubrosSeparados[rubroindex])
			}
		}
	}
	return FuentesRegistroPlanAdquisicion, nil
}

func SepararRegistrosPorRubro(RegistroPlanAdquisicion []map[string]interface{}) (registrosSeparados []map[string]interface{}, outputError interface{}) {
	var rubro map[string]interface{}
	var unicos []string
	RubrosRegistroPlanAdquisicion := make([]map[string]interface{}, 0)
	for rubroindex := range RegistroPlanAdquisicion {
		RubroPorAgregar := RegistroPlanAdquisicion[rubroindex]["RubroId"].(string)
		newRubro := stringInSlice(RubroPorAgregar, unicos)
		if !newRubro {
			idStr := fmt.Sprintf("%.0f", RegistroPlanAdquisicion[rubroindex]["Id"].(float64))
			Vigencia, AreaFuncional, errorVigencia := VigenciaYAreaFuncional(idStr)
			if errorVigencia != nil {
				return nil, errorVigencia
			}
			delete(RegistroPlanAdquisicion[rubroindex], "PlanAdquisicionesId")
			unicos = append(unicos, RubroPorAgregar)
			Rubro, error := ObtenerRubroByID(RubroPorAgregar, Vigencia, AreaFuncional)
			if error != nil {
				return nil, error
			} else {
				rubro = map[string]interface{}{
					"Rubro":     RubroPorAgregar,
					"RubroInfo": Rubro,
					"datos":     make([]map[string]interface{}, 0),
				}
				rubro["datos"] = append(rubro["datos"].([]map[string]interface{}), RegistroPlanAdquisicion[rubroindex])
				RubrosRegistroPlanAdquisicion = append(RubrosRegistroPlanAdquisicion, rubro)
			}

		} else {
			delete(RegistroPlanAdquisicion[rubroindex], "PlanAdquisicionesId")
			index := BuscarIndexPorCampo(RubrosRegistroPlanAdquisicion, RubroPorAgregar, "Rubro")
			if index != -1 {
				RubrosRegistroPlanAdquisicion[index]["datos"] = append(RubrosRegistroPlanAdquisicion[index]["datos"].([]map[string]interface{}), RegistroPlanAdquisicion[rubroindex])
			}
		}
	}
	return RubrosRegistroPlanAdquisicion, nil
}
func BuscarIndexPorCampo(RegistroPlanAdquisicion []map[string]interface{}, Rubro string, Campo string) (index int) {

	for i := range RegistroPlanAdquisicion {
		if RegistroPlanAdquisicion[i][Campo] == Rubro {
			return i
		}
	}
	return -1
}

//IngresoPlanAdquisicion crea un registro de plan de adquisicion
func IngresoPlanAdquisicion(registroPlanAdquisicion map[string]interface{}) (registroPlanAdquisicionRespuesta []map[string]interface{}, outputError interface{}) {
	defer errorctrl.ErrorControlFunction("IngresoPlanAdquisicion - Unhandled Error!", "500")
	var movimientoExternoID int
	// logs.Debug("registroPlanAdquisicion: ", registroPlanAdquisicion)
	if registroPlanAdquisicion["FuenteFinanciamientoId"] == "" {
		resultadoPlan, err := IngresoRenglonPlanInversion(registroPlanAdquisicion)
		if err != nil {
			logs.Error(err)
			return nil, err
		} else {
			idPlanAdquisiciones := int(registroPlanAdquisicion["PlanAdquisicionesId"].(float64))

			filtroJsonB, _ := utils.Serializar(map[string]interface{}{
				"Estado":              "Preliminar",
				"PlanAdquisicionesId": strconv.Itoa(idPlanAdquisiciones),
			})

			query := filtroJsonB

			// Se sugiere ordenar por fecha de modificación
			sortby := "FechaModificacion"

			// El orden descendente velará por traer el último registro modificado
			order := "desc"

			// Para traer el último
			limit := "1"

			if resultado, err := movimientosCrud.GetMovimientoProcesoExterno(query, "", sortby, order, "", limit); err != nil {
				outputError = errorctrl.Error("IngresoPlanAdquisicion -  movimientosCrud.GetMovimientoProcesoExterno(query, \"\", sortby, order, \"\", limit)", err, "502")
				return nil, outputError
			} else {
				var movimientoObtenido []interface{}
				switch resultado.(type) {
				case []interface{}:
					if len(resultado.([]interface{})) > 0 {
						// logs.Debug("Traje información")
						movimientoObtenido = resultado.([]interface{})
					} else {
						// logs.Debug("No encontré información")
						movimientoObtenido = resultado.([]interface{})
					}
				case nil:
					err := "No se encontraron resultados"
					logs.Error(err)
					outputError = errorctrl.Error("IngresoPlanAdquisicion - resultado.(type)", err, "404")
					return nil, outputError
				default:
					err := "La variable resultado no tiene un tipo de dato coherente"
					logs.Error(err)
					outputError = errorctrl.Error("IngresoPlanAdquisicion - resultado.(type)", err, "409")
					return nil, outputError
				}

				if len(movimientoObtenido) > 0 {
					movimientoExternoID = int(movimientoObtenido[0].(map[string]interface{})["Id"].(float64))
				} else {
					if movimientoInsertar, err := ObtenerMovimientoProcesoExterno(idPlanAdquisiciones); err != nil {
						logs.Error(err)
						return nil, err
					} else {
						// logs.Debug(movimientoInsertar)
						if idMovimientoInsertado, err := movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar); err != nil {
							logs.Error(err)
							outputError = errorctrl.Error("IngresoPlanAdquisicion -  movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar)", err, "502")
							return nil, outputError
						} else {
							// logs.Debug("ID OBTENIDO: ", idMovimientoInsertado)
							movimientoExternoID = int(idMovimientoInsertado.(map[string]interface{})["Id"].(float64))
						}
					}
				}

			}

			if resultado, err := ObtenerRegistroMovimientoInversion(registroPlanAdquisicion, movimientoExternoID); err != nil {
				logs.Error(err)
				outputError = map[string]interface{}{
					"funcion": "IngresoPlanAdquisicion - ObtenerRegistroMovimientoInversion(registroPlanAdquisicion, movimientoExternoID)",
					"err":     err,
					"status":  "502",
				}
				return nil, outputError
			} else {
				if len(resultado) > 0 {
					if _, err := movimientosCrud.CrearMovimientosDetalle(resultado); err != nil {
						logs.Error(err)
						outputError = errorctrl.Error("IngresoPlanAdquisicion -  movimientosCrud.CrearMovimientosDetalle(resultado)", err, "502")
						return nil, outputError
					}
				}
			}
			return resultadoPlan, nil
		}
	} else {
		resultadoPlan, err := IngresoRenglonPlanFuncionamiento(registroPlanAdquisicion)
		if err != nil {
			logs.Error(err)
			outputError := errorctrl.Error("IngresoPlanAdquisicion -  IngresoRenglonPlanFuncionamiento(registroPlanAdquisicion)", err, "502")
			return nil, outputError
		} else {
			idPlanAdquisiciones := int(registroPlanAdquisicion["PlanAdquisicionesId"].(float64))

			filtroJsonB, _ := utils.Serializar(map[string]interface{}{
				"Estado":              "Preliminar",
				"PlanAdquisicionesId": strconv.Itoa(idPlanAdquisiciones),
			})

			query := filtroJsonB

			// Se sugiere ordenar por fecha de modificación
			sortby := "FechaModificacion"

			// El orden descendente velará por traer el último registro modificado
			order := "desc"

			// Para traer el último
			limit := "1"

			if resultado, err := movimientosCrud.GetMovimientoProcesoExterno(query, "", sortby, order, "", limit); err != nil {
				outputError = map[string]interface{}{
					"funcion": "IngresoPlanAdquisicion -  movimientosCrud.GetMovimientoProcesoExterno(query, \"\", sortby, order, \"\", limit)",
					"err":     err,
					"status":  "502",
				}
				return nil, outputError
			} else {
				var movimientoObtenido []interface{}
				switch resultado.(type) {
				case []interface{}:
					if len(resultado.([]interface{})) > 0 {
						// logs.Debug("Traje información")
						movimientoObtenido = resultado.([]interface{})
					} else {
						// logs.Debug("No encontré información")
						movimientoObtenido = resultado.([]interface{})
					}
				case nil:
					err := "No se encontraron resultados"
					logs.Error(err)
					outputError = errorctrl.Error("IngresoPlanAdquisicion - resultado.(type)", err, "404")
					return nil, outputError
				default:
					err := "La variable resultado no tiene un tipo de dato coherente"
					logs.Error(err)
					outputError = errorctrl.Error("IngresoPlanAdquisicion - resultado.(type)", err, "409")
					return nil, outputError
				}

				if len(movimientoObtenido) > 0 {
					movimientoExternoID = int(movimientoObtenido[0].(map[string]interface{})["Id"].(float64))
				} else {
					if movimientoInsertar, err := ObtenerMovimientoProcesoExterno(idPlanAdquisiciones); err != nil {
						logs.Error(err)
						return nil, err
					} else {
						// logs.Debug(movimientoInsertar)
						if idMovimientoInsertado, err := movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar); err != nil {
							logs.Error(err)
							outputError = errorctrl.Error("IngresoPlanAdquisicion -  movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar)", err, "502")
							return nil, outputError
						} else {
							// logs.Debug("ID OBTENIDO: ", idMovimientoInsertado)
							movimientoExternoID = int(idMovimientoInsertado.(map[string]interface{})["Id"].(float64))
						}
					}
				}

			}

			if resultado, err := ObtenerRegistroMovimientoFuncionamiento(registroPlanAdquisicion, movimientoExternoID); err != nil {
				logs.Error(err)
				outputError = map[string]interface{}{
					"funcion": "IngresoPlanAdquisicion -  ObtenerRegistroMovimientoFuncionamiento(registroPlanAdquisicion, movimientoExternoID)",
					"err":     err,
					"status":  "502",
				}
				return nil, outputError
			} else {
				// logs.Debug("Cuentas insertar: ", resultado)
				if len(resultado) > 0 {
					if _, err := movimientosCrud.CrearMovimientosDetalle(resultado); err != nil {
						logs.Error(err)
						outputError = errorctrl.Error("IngresoPlanAdquisicion -  movimientosCrud.CrearMovimientosDetalle(resultado)", err, "502")
						return nil, outputError
					}
				}
			}
			return resultadoPlan, nil
		}
	}
}

func IngresoRenglonPlanInversion(registroPlanAdquisicion map[string]interface{}) (registroPlanAdquisicionRespuesta []map[string]interface{}, outputError interface{}) {
	registroPlanAdquisicionIngresado := make(map[string]interface{})
	registroPlanAdquisicionPost := make(map[string]interface{})
	result := []map[string]interface{}{}

	registroPlanAdquisicionIngresado = map[string]interface{}{
		"AreaFuncional":          registroPlanAdquisicion["AreaFuncional"],
		"CentroGestor":           registroPlanAdquisicion["CentroGestor"],
		"ResponsableId":          registroPlanAdquisicion["ResponsableId"],
		"RubroId":                registroPlanAdquisicion["RubroId"],
		"MetaId":                 nil,
		"ProductoId":             nil,
		"FuenteFinanciamientoId": "",
		"ActividadId":            0,
		"ValorActividad":         0,
		"FechaEstimadaInicio":    registroPlanAdquisicion["FechaEstimadaInicio"],
		"FechaEstimadaFin":       registroPlanAdquisicion["FechaEstimadaFin"],
		"Activo":                 registroPlanAdquisicion["Activo"],
		"PlanAdquisicionesId":    map[string]interface{}{"Id": registroPlanAdquisicion["PlanAdquisicionesId"]},
	}
	ModalidadSeleccion := registroPlanAdquisicion["ModalidadSeleccion"].([]interface{})
	CodigoArka := registroPlanAdquisicion["CodigoArka"].([]interface{})
	PlanAdquisicionActividad := registroPlanAdquisicion["RegistroPlanAdquisicionActividad"].([]interface{})
	MetasAsociadas := registroPlanAdquisicion["MetasAsociadas"].([]interface{})
	ProductosAsociados := registroPlanAdquisicion["ProductosAsociados"].([]interface{})

	// Ojo comprobacion de los valores de los rubros con las fuentes !!!!!!!
	// PlanAdquisicionesID := fmt.Sprintf("%.0f", registroPlanAdquisicion["PlanAdquisicionesId"].(float64))
	// AreaFuncional := fmt.Sprintf("%.0f", registroPlanAdquisicion["AreaFuncional"].(float64))
	// Vigencia, errorVigencia := VigenciaYCentroGestorByPlanID(PlanAdquisicionesID)
	// if errorVigencia != nil {
	// 	return nil, errorVigencia
	// }
	// errorSuma := SumaFuenteFinanciamiento(PlanAdquisicionActividad, registroPlanAdquisicion["RubroId"].(string), Vigencia, AreaFuncional)
	// if errorSuma != nil {
	// 	return nil, errorSuma
	// }
	// !!!!!!!!!

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
					DatosMetasAsociadas, errorMetasAsociadas := GuardarMetasAsociadas(MetasAsociadas, registroPlanAdquisicionPost["Id"])
					if errorMetasAsociadas != nil {
						return nil, errorMetasAsociadas
					} else {
						result = append(result, DatosMetasAsociadas...)
						DatosProductosAsociados, errorProductosAsociados := GuardarProductosAsociados(ProductosAsociados, registroPlanAdquisicionPost["Id"])
						if errorProductosAsociados != nil {
							return nil, errorProductosAsociados
						} else {
							result = append(result, DatosProductosAsociados...)
						}
					}
				}
			}
		}
		return result, nil
	}
}

func IngresoRenglonPlanFuncionamiento(registroPlanAdquisicion map[string]interface{}) (registroPlanAdquisicionRespuesta []map[string]interface{}, outputError interface{}) {
	// logs.Debug("registroPlanAdquisicion: ", registroPlanAdquisicion)
	registroPlanAdquisicionIngresado := make(map[string]interface{})
	registroPlanAdquisicionPost := make(map[string]interface{})
	result := []map[string]interface{}{}

	registroPlanAdquisicionIngresado = map[string]interface{}{
		"AreaFuncional":          registroPlanAdquisicion["AreaFuncional"],
		"CentroGestor":           registroPlanAdquisicion["CentroGestor"],
		"ResponsableId":          registroPlanAdquisicion["ResponsableId"],
		"RubroId":                registroPlanAdquisicion["RubroId"],
		"MetaId":                 nil,
		"ProductoId":             nil,
		"ActividadId":            registroPlanAdquisicion["ActividadId"],
		"ValorActividad":         registroPlanAdquisicion["ValorActividad"],
		"FuenteFinanciamientoId": registroPlanAdquisicion["FuenteFinanciamientoId"],
		"FechaEstimadaInicio":    registroPlanAdquisicion["FechaEstimadaInicio"],
		"FechaEstimadaFin":       registroPlanAdquisicion["FechaEstimadaFin"],
		"Activo":                 registroPlanAdquisicion["Activo"],
		"PlanAdquisicionesId":    map[string]interface{}{"Id": registroPlanAdquisicion["PlanAdquisicionesId"]},
	}
	ModalidadSeleccion := registroPlanAdquisicion["ModalidadSeleccion"].([]interface{})
	CodigoArka := registroPlanAdquisicion["CodigoArka"].([]interface{})

	// ! WARNING Ojo, se debe comprobar la suma de los rubros con los datos !!!!
	// PlanAdquisicionesID := fmt.Sprintf("%.0f", registroPlanAdquisicion["PlanAdquisicionesId"].(float64))
	// AreaFuncional := fmt.Sprintf("%.0f", registroPlanAdquisicion["AreaFuncional"].(float64))
	// Vigencia, errorVigencia := VigenciaYCentroGestorByPlanID(PlanAdquisicionesID)
	// if errorVigencia != nil {
	// 	return nil, errorVigencia
	// }
	// errorSuma := SumaFuenteFinanciamientoFuncionamiento(registroPlanAdquisicion["ValorActividad"], registroPlanAdquisicion["RubroId"].(string), Vigencia, AreaFuncional)
	// if errorSuma != nil {
	// 	return nil, errorSuma
	// }
	// !!!!!!

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
			}
		}
		return result, nil
	}
}

//ObtenerRenglonRegistroPlanAdquisicionByID regresa un renglon segun el id del registro de plan de adquisicion
func ObtenerRenglonRegistroPlanAdquisicionByID(idStr string) (renglonRegistroPlanAdquisicion []map[string]interface{}, outputError interface{}) {
	var RenglonRegistroPlanAdquisicion []map[string]interface{}
	// TODO: Se podría reemplazar la petición dependiendo los datos necesarios
	query := beego.AppConfig.String("plan_adquicisiones_crud_url") + "Registro_plan_adquisiciones?query=Id:" + idStr
	// logs.Debug("query:", query)
	error := request.GetJson(query, &RenglonRegistroPlanAdquisicion)
	if error != nil {
		return nil, error
	} else {
		// fmt.Println(RenglonRegistroPlanAdquisicion)
		if len(RenglonRegistroPlanAdquisicion) == 0 {
			error := "No existe Registro Plan Adquisicion"
			return nil, error
		} else {
			if RenglonRegistroPlanAdquisicion[0]["FuenteFinanciamientoId"] == "" {
				RenglonRegistro, errorDatos := ObtenerRenglonInversion(RenglonRegistroPlanAdquisicion[0], idStr)
				// logs.Debug("errorDatos:", errorDatos)
				if errorDatos != nil {
					logs.Warning(errorDatos)
					return nil, errorDatos
				} else {
					return RenglonRegistro, nil
				}
			} else {
				RenglonRegistro, errorDatos := ObtenerRenglonFuncionamiento(RenglonRegistroPlanAdquisicion[0], idStr)
				if errorDatos != nil {
					return nil, errorDatos
				} else {
					return RenglonRegistro, nil
				}
			}
		}

	}

}

func ObtenerRenglonInversion(RenglonRegistro map[string]interface{}, idStr string) (renglonRegistroPlanAdquisicion []map[string]interface{}, outputError interface{}) {

	var RenglonRegistroPlanAdquisicion []map[string]interface{}
	var Responsable []map[string]interface{}
	query := beego.AppConfig.String("plan_adquicisiones_crud_url") + "Registro_plan_adquisiciones/?query=Id:" + idStr
	error := request.GetJson(query, &RenglonRegistroPlanAdquisicion)
	if error != nil {
		return nil, error
	} else {
		if len(RenglonRegistroPlanAdquisicion) == 0 {
			error := "No existe Registro Plan Adquisicion"
			return nil, error
		}
		CodigoArka, error := ObtenerRegistroCodigoArkaByIDPlanAdquisicion(idStr)
		// logs.Debug("error: ", error)
		if error != nil {
			return nil, error
		} else {
			ModalidadSeleccion, error := ObtenerRegistroModalidadSeleccionByIDPlanAdquisicion(idStr)
			// logs.Debug("error: ", error)
			if error != nil {
				return nil, error
			} else {
				Metas, error := ObtenerRegistroMetasAsociadasByIDPlanAdquisicion(idStr)
				// logs.Debug("error: ", error)
				if error != nil {
					return nil, error
				} else {
					Productos, error := ObtenerRegistroProductosAsociadosByIDPlanAdquisicion(idStr)
					// logs.Debug("error: ", error)
					if error != nil {
						return nil, error
					} else {
						RegistroPlanAdquisicionActividad, error := ObtenerRegistroTablaActividades(idStr)
						logs.Debug("error: ", error)
						// logs.Debug("len(RegistroPlanAdquisicionActividad): ", len(RegistroPlanAdquisicionActividad))
						if error != nil {
							return nil, error
						} else {
							s := fmt.Sprintf("%.0f", RenglonRegistroPlanAdquisicion[0]["ResponsableId"].(float64))
							error := request.GetJson(beego.AppConfig.String("oikos_api_url")+"dependencia/?query=Id:"+s, &Responsable)
							if error != nil {
								return nil, error
							} else {
								valorTotalActividad := SumaActividades(RegistroPlanAdquisicionActividad)
								EliminarCampos(CodigoArka, "RegistroPlanAdquisicionesId")
								EliminarCampos(ModalidadSeleccion, "RegistroPlanAdquisicionesId")
								RenglonRegistroPlanAdquisicion[0]["registro_plan_adquisiciones-codigo_arka"] = CodigoArka
								RenglonRegistroPlanAdquisicion[0]["registro_funcionamiento-modalidad_seleccion"] = ModalidadSeleccion
								RenglonRegistroPlanAdquisicion[0]["registro_funcionamiento-metas_asociadas"] = Metas
								RenglonRegistroPlanAdquisicion[0]["registro_funcionamiento-productos_asociados"] = Productos
								RenglonRegistroPlanAdquisicion[0]["registro_plan_adquisiciones-actividad"] = RegistroPlanAdquisicionActividad
								RenglonRegistroPlanAdquisicion[0]["ResponsableNombre"] = Responsable[0]["Nombre"]
								RenglonRegistroPlanAdquisicion[0]["ValorTotalActividades"] = valorTotalActividad
							}
						}
					}
				}
			}
		}
		return RenglonRegistroPlanAdquisicion, nil
	}

}

func ObtenerRenglonFuncionamiento(RenglonRegistro map[string]interface{}, idStr string) (renglonRegistroPlanAdquisicion []map[string]interface{}, outputError interface{}) {

	var RenglonRegistroPlanAdquisicion []map[string]interface{}
	var Responsable []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/?query=Id:"+idStr, &RenglonRegistroPlanAdquisicion)
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
				Vigencia, AreaFuncional, errorVigenciaYAreaFuncional := VigenciaYAreaFuncional(idStr)
				if error != nil && errorVigenciaYAreaFuncional != nil {
					return nil, error
				} else {
					FuenteFinanciamiento, error := ObtenerFuenteFinanciamientoByCodigo(MapToString(RenglonRegistroPlanAdquisicion[0]["FuenteFinanciamientoId"]), Vigencia, AreaFuncional)
					if error != nil {
						return nil, error
					} else {
						ActividadData, error := ObtenerActividadById(RenglonRegistroPlanAdquisicion[0]["ActividadId"])
						if error != nil {
							return nil, error
						} else {
							s := fmt.Sprintf("%.0f", RenglonRegistroPlanAdquisicion[0]["ResponsableId"].(float64))
							error := request.GetJson(beego.AppConfig.String("oikos_api_url")+"dependencia/?query=Id:"+s, &Responsable)
							if error != nil {
								return nil, error
							} else {
								EliminarCampos(CodigoArka, "RegistroPlanAdquisicionesId")
								EliminarCampos(ModalidadSeleccion, "RegistroPlanAdquisicionesId")
								RenglonRegistroPlanAdquisicion[0]["registro_plan_adquisiciones-codigo_arka"] = CodigoArka
								RenglonRegistroPlanAdquisicion[0]["registro_funcionamiento-modalidad_seleccion"] = ModalidadSeleccion
								RenglonRegistroPlanAdquisicion[0]["ActividadData"] = ActividadData
								RenglonRegistroPlanAdquisicion[0]["FuenteFinanciamientoData"] = FuenteFinanciamiento
								RenglonRegistroPlanAdquisicion[0]["ResponsableNombre"] = Responsable[0]["Nombre"]
							}
						}
					}
				}
			}
		}
		return RenglonRegistroPlanAdquisicion, nil
	}

}

//ActualizarRegistroPlanAdquisicion verifica y actualiza los campos de un renglon segun el ID de un registro plan de adquisicion
func ActualizarRegistroPlanAdquisicion(registroPlanAdquisicion map[string]interface{}, idStr string) (registroActividadRespuesta map[string]interface{}, outputError interface{}) {
	// logs.Debug("registroPlanAdquisicion: ")
	// formatdata.JsonPrint(registroPlanAdquisicion)
	if registroPlanAdquisicion["FuenteFinanciamientoId"] == "" {
		registro, error := ActualizarRegistroInversion(registroPlanAdquisicion, idStr)
		if error != nil {
			return nil, error
		} else {
			if err := helpers.RegistrarMovimientoDetalleActualizacionInversion(registroPlanAdquisicion); err != nil {
				return nil, err
			}
			return registro, nil
		}
	} else {
		registro, error := ActualizarRegistroFuncionamiento(registroPlanAdquisicion, idStr)
		if error != nil {
			return nil, error
		} else {
			if err := helpers.RegistrarMovimientoDetalleActualizacionFuncionamiento(registroPlanAdquisicion); err != nil {
				return nil, err
			}
			return registro, nil
		}
	}
}

func ActualizarRegistroFuncionamiento(registroPlanAdquisicion map[string]interface{}, idStr string) (registroActividadRespuesta map[string]interface{}, outputError interface{}) {
	registroPlanAdquisicionPut := make(map[string]interface{})
	registroPlanAdquisicionActualizar := make(map[string]interface{})
	RegistroPlanAdquisicionAntiguo, error := ObtenerRenglonRegistroPlanAdquisicionByID(idStr)

	// Ojo, se debe comprobar la suma de los rubros con los datos !!!!!!
	//fmt.Println("existe registro y  toca modificarlo")
	// Vigencia, AreaFuncional, errorVigencia := VigenciaYAreaFuncional(idStr)
	// if errorVigencia != nil {
	// 	return nil, errorVigencia
	// }
	// PlanAdquisicionActividad := registroPlanAdquisicion["RegistroPlanAdquisicionActividad"].([]interface{})
	// errorSuma := SumaFuenteFinanciamientoFuncionamiento(registroPlanAdquisicion["ValorActividad"], registroPlanAdquisicion["RubroId"].(string), Vigencia, AreaFuncional)
	// if errorSuma != nil {
	// 	return nil, errorSuma
	// }
	// !!!!!!!

	registroPlanAdquisicionActualizar = map[string]interface{}{
		"AreaFuncional":          registroPlanAdquisicion["AreaFuncional"],
		"CentroGestor":           registroPlanAdquisicion["CentroGestor"],
		"ResponsableId":          registroPlanAdquisicion["ResponsableId"],
		"MetaId":                 nil,
		"ProductoId":             nil,
		"ActividadId":            registroPlanAdquisicion["ActividadId"],
		"ValorActividad":         registroPlanAdquisicion["ValorActividad"],
		"RubroId":                registroPlanAdquisicion["RubroId"],
		"FuenteFinanciamientoId": registroPlanAdquisicion["FuenteFinanciamientoId"],
		"FechaCreacion":          RegistroPlanAdquisicionAntiguo[0]["FechaCreacion"],
		"FechaEstimadaInicio":    registroPlanAdquisicion["FechaEstimadaInicio"],
		"FechaEstimadaFin":       registroPlanAdquisicion["FechaEstimadaFin"],
		"Activo":                 registroPlanAdquisicion["Activo"],
		"PlanAdquisicionesId":    map[string]interface{}{"Id": registroPlanAdquisicion["PlanAdquisicionesId"]},
	}
	error2 := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/"+idStr, "PUT", &registroPlanAdquisicionPut, registroPlanAdquisicionActualizar)
	if error2 != nil {
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
				return registroPlanAdquisicion, nil
			}
		}
	}
}

func ActualizarRegistroInversion(registroPlanAdquisicion map[string]interface{}, idStr string) (registroActividadRespuesta map[string]interface{}, outputError interface{}) {
	registroPlanAdquisicionPut := make(map[string]interface{})
	registroPlanAdquisicionActualizar := make(map[string]interface{})
	RegistroPlanAdquisicionAntiguo, error := ObtenerRenglonRegistroPlanAdquisicionByID(idStr)

	// Ojo funcion para comprobar valores de rubros !!!!!
	//fmt.Println("existe registro y  toca modificarlo")
	// Vigencia, AreaFuncional, errorVigencia := VigenciaYAreaFuncional(idStr)
	// if errorVigencia != nil {
	// 	return nil, errorVigencia
	// }
	// PlanAdquisicionActividad := registroPlanAdquisicion["RegistroPlanAdquisicionActividad"].([]interface{})
	// errorSuma := SumaFuenteFinanciamiento(PlanAdquisicionActividad, registroPlanAdquisicion["RubroId"].(string), Vigencia, AreaFuncional)
	// if errorSuma != nil {
	// 	return nil, errorSuma
	// }
	// !!!!!!!

	registroPlanAdquisicionActualizar = map[string]interface{}{
		"AreaFuncional":          registroPlanAdquisicion["AreaFuncional"],
		"CentroGestor":           registroPlanAdquisicion["CentroGestor"],
		"ResponsableId":          registroPlanAdquisicion["ResponsableId"],
		"MetaId":                 nil,
		"ProductoId":             nil,
		"FuenteFinanciamientoId": "",
		"ActividadId":            0,
		"ValorActividad":         0,
		"RubroId":                registroPlanAdquisicion["RubroId"],
		"FechaCreacion":          RegistroPlanAdquisicionAntiguo[0]["FechaCreacion"],
		"FechaEstimadaInicio":    registroPlanAdquisicion["FechaEstimadaInicio"],
		"FechaEstimadaFin":       registroPlanAdquisicion["FechaEstimadaFin"],
		"Activo":                 registroPlanAdquisicion["Activo"],
		"PlanAdquisicionesId":    map[string]interface{}{"Id": registroPlanAdquisicion["PlanAdquisicionesId"]},
	}
	error2 := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/"+idStr, "PUT", &registroPlanAdquisicionPut, registroPlanAdquisicionActualizar)
	if error2 != nil {
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
					error := MetasAsociadasModificado(registroPlanAdquisicion, idStr)
					if error != nil {
						return nil, error
					} else {
						error := ProductosAsociadosModificado(registroPlanAdquisicion, idStr)
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
	fuentes := fuente[0] + "-" + fuente[1]
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

// VigenciaYAreaFuncional regresa vigencia y AreaFuncional
func VigenciaYAreaFuncional(RegistroplanAdquisicionID string) (Vigencia string, AreaFuncional string, outputError interface{}) {
	var RegistroPlanAdquisicion []map[string]interface{}

	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/?query=Id:"+RegistroplanAdquisicionID+"&fields=AreaFuncional,PlanAdquisicionesId", &RegistroPlanAdquisicion)
	if error != nil {
		return "", "", error
	} else {
		AreaFuncional := fmt.Sprintf("%.0f", RegistroPlanAdquisicion[0]["AreaFuncional"].(float64))
		vigencia := fmt.Sprintf("%.0f", RegistroPlanAdquisicion[0]["PlanAdquisicionesId"].(map[string]interface{})["Vigencia"].(float64))
		return vigencia, AreaFuncional, nil
	}
}

// VigenciaYCentroGestorByPlanID regresa vigencia si no se tiene un Id del registro de plan de adquisicion
func VigenciaYCentroGestorByPlanID(PlanID string) (Vigencia string, outputError interface{}) {

	var Planadquisicion map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Plan_adquisiciones/"+PlanID, &Planadquisicion)
	if error != nil {
		return "", error
	} else {
		vigencia := fmt.Sprintf("%.0f", Planadquisicion["Vigencia"].(float64))
		return vigencia, nil
	}

}
