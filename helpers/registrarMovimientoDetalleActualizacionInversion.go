package helpers

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/movimientosCrud"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/utils"
	"github.com/udistrital/utils_oas/errorctrl"
)

func init() {
	// ? Parametrizable
	rFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
}

func RegistrarMovimientoDetalleActualizacionInversion(registroPlanAdquisicion map[string]interface{}) (outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("RegistrarMovimientoDetalleActualizacionInversion - Unhandled error!", "500")

	var idPlanAdquisiciones int
	var err error
	var filtroJsonB string
	var movimientoPublicado interface{}
	var movimientoPreliminar interface{}
	var query string
	// Se sugiere ordenar por fecha de modificación
	sortby := "FechaModificacion"
	// El orden descendente velará por traer el último registro modificado
	order := "desc"
	// Para traer el último
	limit := "1"
	var movimientoPublicadoObtenido []interface{}
	var movimientoPreliminarObtenido []interface{}
	var movimientoExternoID int

	switch registroPlanAdquisicion["PlanAdquisicionesId"].(type) {
	case float64:
		idPlanAdquisiciones = int(registroPlanAdquisicion["PlanAdquisicionesId"].(float64))
	case nil:
		err := "El tipo de dato es nil"
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - registroPlanAdquisicion[\"PlanAdquisicionesId\"].(type)", err, "400")
		logs.Error(err)
		return outputError
	default:
		err := "No se reconoce el tipo de dato"
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - registroPlanAdquisicion[\"PlanAdquisicionesId\"].(type)", err, "400")
		logs.Error(err)
		return outputError
	}

	// * Estado por defecto es "Publicado"
	// ? Parametrizable o administrable
	if filtroJsonB, err = utils.Serializar(map[string]interface{}{
		"Estado":              "Publicado",
		"PlanAdquisicionesId": strconv.Itoa(idPlanAdquisiciones),
	}); err != nil {
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - utils.Serializar(map[string]interface{}", err, "500")
		logs.Error(err)
		return outputError
	}

	query = filtroJsonB

	if movimientoPublicado, outputError = movimientosCrud.GetMovimientoProcesoExterno(query, "", sortby, order, "", limit); err != nil {
		return outputError
	}

	// * Estado por defecto es "Preliminar"
	// ? Parametrizable o administrable
	if filtroJsonB, err = utils.Serializar(map[string]interface{}{
		"Estado":              "Preliminar",
		"PlanAdquisicionesId": strconv.Itoa(idPlanAdquisiciones),
	}); err != nil {
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - utils.Serializar(map[string]interface{}", err, "500")
		logs.Error(err)
		return outputError
	}

	query = filtroJsonB

	if movimientoPreliminar, outputError = movimientosCrud.GetMovimientoProcesoExterno(query, "", sortby, order, "", limit); err != nil {
		return outputError
	}

	movimientoPublicadoObtenido, outputError = KeysMovimientoProcesoExterno(movimientoPublicado)
	if outputError != nil {
		return outputError
	}

	movimientoPreliminarObtenido, outputError = KeysMovimientoProcesoExterno(movimientoPreliminar)
	if outputError != nil {
		return outputError
	}

	if len(movimientoPublicadoObtenido) > 0 && len(movimientoPreliminarObtenido) > 0 {
		// movimientoObtenido[0].(map[string]interface{})["Id"].(float64)
		// logs.Debug(reflect.TypeOf(movimientoPublicadoObtenido[0].(map[string]interface{})["FechaCreacion"]))
		// logs.Debug(reflect.TypeOf(movimientoPreliminarObtenido[0].(map[string]interface{})["FechaCreacion"]))
		tPreliminar, err := time.Parse(rFC3339Nano, movimientoPreliminarObtenido[0].(map[string]interface{})["FechaCreacion"].(string))
		if err != nil {
			outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - time.Parse(RFC3339Nano, movimientoPreliminarObtenido[\"FechaCreacion\"].(string))", err, "500")
			logs.Error(err)
			return outputError
		}
		// logs.Debug("tPreliminar: ", tPreliminar)

		tPublicado, err := time.Parse(rFC3339Nano, movimientoPublicadoObtenido[0].(map[string]interface{})["FechaCreacion"].(string))
		if err != nil {
			outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - time.Parse(RFC3339Nano, movimientoPublicadoObtenido[\"FechaCreacion\"].(string))", err, "500")
			logs.Error(err)
			return outputError
		}
		// logs.Debug("tPublicado: ", tPublicado)

		if tPreliminar.After(tPublicado) {
			// logs.Debug("Es un preliminar después de publicar")
			switch movimientoPreliminarObtenido[0].(map[string]interface{})["Id"].(type) {
			case float64:
				movimientoExternoID = int(movimientoPreliminarObtenido[0].(map[string]interface{})["Id"].(float64))
			case nil:
				err := "El tipo de dato es nil"
				outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientoPreliminarObtenido[\"Id\"].(type)", err, "500")
				logs.Error(err)
				return outputError
			default:
				err := "No se reconoce el tipo de dato"
				outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientoPreliminarObtenido[\"Id\"].(type)", err, "500")
				logs.Error(err)
				return outputError
			}
		} else if tPreliminar.Before(tPublicado) {
			// logs.Debug("Es un preliminar anterior a la publicación")
			if movimientoInsertar, err := ObtenerMovimientoProcesoExterno(idPlanAdquisiciones); err != nil {
				logs.Error(err)
				return err
			} else {
				// logs.Debug(movimientoInsertar)
				if idMovimientoInsertado, outputError := movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar); err != nil {
					return outputError
				} else {
					// logs.Debug("ID OBTENIDO: ", idMovimientoInsertado)
					switch idMovimientoInsertado.(map[string]interface{})["Id"].(type) {
					case float64:
						movimientoExternoID = int(idMovimientoInsertado.(map[string]interface{})["Id"].(float64))
					case nil:
						err := "El tipo de dato es nil"
						outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - idMovimientoInsertado.(map[string]interface{})[\"Id\"].(type)", err, "500")
						logs.Error(err)
						return outputError
					default:
						err := "No se reconoce el tipo de dato"
						outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - idMovimientoInsertado.(map[string]interface{})[\"Id\"].(type)", err, "500")
						logs.Error(err)
						return outputError
					}
				}
			}
		}
	} else if len(movimientoPreliminarObtenido) > 0 {
		// logs.Debug("No hay planes publicados")
		switch movimientoPreliminarObtenido[0].(map[string]interface{})["Id"].(type) {
		case float64:
			movimientoExternoID = int(movimientoPreliminarObtenido[0].(map[string]interface{})["Id"].(float64))
		case nil:
			err := "El tipo de dato es nil"
			outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientoPreliminarObtenido[\"Id\"].(type)", err, "500")
			logs.Error(err)
			return outputError
		default:
			err := "No se reconoce el tipo de dato"
			outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientoPreliminarObtenido[\"Id\"].(type)", err, "500")
			logs.Error(err)
			return outputError
		}
	} else if len(movimientoPublicadoObtenido) > 0 {
		// logs.Debug("No hay planes preliminares")
		if movimientoInsertar, err := ObtenerMovimientoProcesoExterno(idPlanAdquisiciones); err != nil {
			return err
		} else {
			// logs.Debug(movimientoInsertar)
			if idMovimientoInsertado, outputError := movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar); err != nil {
				logs.Error(outputError["err"])
				return outputError
			} else {
				// logs.Debug("ID OBTENIDO: ", idMovimientoInsertado)
				switch idMovimientoInsertado.(map[string]interface{})["Id"].(type) {
				case float64:
					movimientoExternoID = int(idMovimientoInsertado.(map[string]interface{})["Id"].(float64))
				case nil:
					err := "El tipo de dato es nil"
					outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - idMovimientoInsertado.(map[string]interface{})[\"Id\"].(type)", err, "500")
					logs.Error(err)
					return outputError
				default:
					err := "No se reconoce el tipo de dato"
					outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - idMovimientoInsertado.(map[string]interface{})[\"Id\"].(type)", err, "500")
					logs.Error(err)
					return outputError
				}
			}
		}
	}

	// logs.Debug("movimientoPublicado: ", movimientoPublicado)
	// logs.Debug("movimientoPreliminar: ", movimientoPreliminar)

	if registroMovimientoInversion, err := ObtenerRegistroMovimientoInversion(registroPlanAdquisicion, movimientoExternoID); err != nil {
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - ObtenerRegistroMovimientoInversion(registroPlanAdquisicion, movimientoExternoID)", err, "502")
		logs.Error(err)
		return outputError
	} else {
		if len(registroMovimientoInversion) > 0 {
			if _, outputError := movimientosCrud.CrearMovimientosDetalle(registroMovimientoInversion); err != nil {
				logs.Error(outputError["err"])
				return outputError
			}
		}
	}

	return nil
}
