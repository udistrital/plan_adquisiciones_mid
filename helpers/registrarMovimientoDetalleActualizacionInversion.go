package helpers

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/movimientosCrud"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/utils"
	"github.com/udistrital/utils_oas/errorctrl"
)

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
	var movimientoPublicadoObtenido map[string]interface{}
	var movimientoPreliminarObtenido map[string]interface{}
	// ? Parametrizable
	RFC3339Nano := "2006-01-02T15:04:05.999999999Z07:00"
	var movimientoExternoID int

	switch registroPlanAdquisicion["PlanAdquisicionesId"].(type) {
	case float64:
		idPlanAdquisiciones = int(registroPlanAdquisicion["PlanAdquisicionesId"].(float64))
	case nil:
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - registroPlanAdquisicion[\"PlanAdquisicionesId\"].(type)", "El tipo de dato es nil", "400")
		logs.Error(outputError)
		return outputError
	default:
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - registroPlanAdquisicion[\"PlanAdquisicionesId\"].(type)", "No se reconoce el tipo de dato", "400")
		logs.Error(outputError)
		return outputError
	}

	// * Estado por defecto es "Publicado"
	// ? Parametrizable o administrable
	if filtroJsonB, err = utils.Serializar(map[string]interface{}{
		"Estado":              "Publicado",
		"PlanAdquisicionesId": strconv.Itoa(idPlanAdquisiciones),
	}); err != nil {
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - utils.Serializar(map[string]interface{}", err, "500")
		logs.Error(outputError)
		return outputError
	}

	query = filtroJsonB

	if movimientoPublicado, outputError = movimientosCrud.GetMovimientoProcesoExterno(query, "", sortby, order, "", limit); err != nil {
		logs.Error(outputError)
		return outputError
	}

	// * Estado por defecto es "Preliminar"
	// ? Parametrizable o administrable
	if filtroJsonB, err = utils.Serializar(map[string]interface{}{
		"Estado":              "Preliminar",
		"PlanAdquisicionesId": strconv.Itoa(idPlanAdquisiciones),
	}); err != nil {
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - utils.Serializar(map[string]interface{}", err, "500")
		logs.Error(outputError)
		return outputError
	}

	query = filtroJsonB

	if movimientoPreliminar, outputError = movimientosCrud.GetMovimientoProcesoExterno(query, "", sortby, order, "", limit); err != nil {
		logs.Error(outputError)
		return outputError
	}

	keysPublicado, movimientoPublicadoObtenido, outputError := KeysMovimientoProcesoExterno(movimientoPublicado)
	if outputError != nil {
		logs.Error(outputError)
		return outputError
	}

	keysPreliminar, movimientoPreliminarObtenido, outputError := KeysMovimientoProcesoExterno(movimientoPreliminar)
	if outputError != nil {
		logs.Error(outputError)
		return outputError
	}

	if len(keysPublicado) > 0 && len(keysPreliminar) > 0 {
		// logs.Debug(reflect.TypeOf(movimientoPublicadoObtenido["FechaCreacion"]))
		// logs.Debug(reflect.TypeOf(movimientoPreliminarObtenido["FechaCreacion"]))
		tPreliminar, err := time.Parse(RFC3339Nano, movimientoPreliminarObtenido["FechaCreacion"].(string))
		if err != nil {
			outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - time.Parse(RFC3339Nano, movimientoPreliminarObtenido[\"FechaCreacion\"].(string))", err, "500")
			logs.Error(outputError)
			return outputError
		}
		// logs.Debug("tPreliminar: ", tPreliminar)

		tPublicado, err := time.Parse(RFC3339Nano, movimientoPublicadoObtenido["FechaCreacion"].(string))
		if err != nil {
			outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - time.Parse(RFC3339Nano, movimientoPublicadoObtenido[\"FechaCreacion\"].(string))", err, "500")
			logs.Error(outputError)
			return outputError
		}
		// logs.Debug("tPublicado: ", tPublicado)

		if tPreliminar.After(tPublicado) {
			// logs.Debug("Es un preliminar después de publicar")
			switch movimientoPreliminarObtenido["Id"].(type) {
			case float64:
				movimientoExternoID = int(movimientoPreliminarObtenido["Id"].(float64))
			case nil:
				outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientoPreliminarObtenido[\"Id\"].(type)", "El tipo de dato es nil", "500")
				logs.Error(outputError)
				return outputError
			default:
				outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientoPreliminarObtenido[\"Id\"].(type)", "No se reconoce el tipo de dato", "500")
				logs.Error(outputError)
				return outputError
			}
		} else if tPreliminar.Before(tPublicado) {
			// logs.Debug("Es un preliminar anterior a la publicación")
			if movimientoInsertar, err := ObtenerMovimientoProcesoExterno(idPlanAdquisiciones); err != nil {
				logs.Error(err)
				return err
			} else {
				// logs.Debug(movimientoInsertar)
				if idMovimientoInsertado, err := movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar); err != nil {
					outputError = errorctrl.Error("IngresoPlanAdquisicion -  movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar)", err, "502")
					logs.Error(outputError)
					return outputError
				} else {
					// logs.Debug("ID OBTENIDO: ", idMovimientoInsertado)
					switch idMovimientoInsertado.(map[string]interface{})["Body"].(map[string]interface{})["Id"].(type) {
					case float64:
						movimientoExternoID = int(idMovimientoInsertado.(map[string]interface{})["Body"].(map[string]interface{})["Id"].(float64))
					case nil:
						outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - idMovimientoInsertado.(map[string]interface{})[\"Body\"].(map[string]interface{})[\"Id\"].(type)", "El tipo de dato es nil", "500")
						logs.Error(outputError)
						return outputError
					default:
						outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - idMovimientoInsertado.(map[string]interface{})[\"Body\"].(map[string]interface{})[\"Id\"].(type)", "No se reconoce el tipo de dato", "500")
						logs.Error(outputError)
						return outputError
					}
				}
			}
		}
	} else if len(keysPreliminar) > 0 {
		// logs.Debug("No hay planes publicados")
		switch movimientoPreliminarObtenido["Id"].(type) {
		case float64:
			movimientoExternoID = int(movimientoPreliminarObtenido["Id"].(float64))
		case nil:
			outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientoPreliminarObtenido[\"Id\"].(type)", "El tipo de dato es nil", "500")
			logs.Error(outputError)
			return outputError
		default:
			outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientoPreliminarObtenido[\"Id\"].(type)", "No se reconoce el tipo de dato", "500")
			logs.Error(outputError)
			return outputError
		}
	} else if len(keysPublicado) > 0 {
		// logs.Debug("No hay planes preliminares")
		if movimientoInsertar, err := ObtenerMovimientoProcesoExterno(idPlanAdquisiciones); err != nil {
			logs.Error(err)
			return err
		} else {
			// logs.Debug(movimientoInsertar)
			if idMovimientoInsertado, err := movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar); err != nil {
				outputError = errorctrl.Error("IngresoPlanAdquisicion -  movimientosCrud.CrearMovimientoProcesoExterno(movimientoInsertar)", err, "502")
				logs.Error(outputError)
				return outputError
			} else {
				// logs.Debug("ID OBTENIDO: ", idMovimientoInsertado)
				switch idMovimientoInsertado.(map[string]interface{})["Body"].(map[string]interface{})["Id"].(type) {
				case float64:
					movimientoExternoID = int(idMovimientoInsertado.(map[string]interface{})["Body"].(map[string]interface{})["Id"].(float64))
				case nil:
					outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - idMovimientoInsertado.(map[string]interface{})[\"Body\"].(map[string]interface{})[\"Id\"].(type)", "El tipo de dato es nil", "500")
					logs.Error(outputError)
					return outputError
				default:
					outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - idMovimientoInsertado.(map[string]interface{})[\"Body\"].(map[string]interface{})[\"Id\"].(type)", "No se reconoce el tipo de dato", "500")
					logs.Error(outputError)
					return outputError
				}
			}
		}
	}

	// logs.Debug("movimientoPublicado: ", movimientoPublicado)
	// logs.Debug("movimientoPreliminar: ", movimientoPreliminar)

	if registroMovimientoInversion, err := ObtenerRegistroMovimientoInversion(registroPlanAdquisicion, movimientoExternoID); err != nil {
		outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - ObtenerRegistroMovimientoInversion(registroPlanAdquisicion, movimientoExternoID)", err, "502")
		logs.Error(outputError)
		return outputError
	} else {
		if len(registroMovimientoInversion) > 0 {
			if _, err := movimientosCrud.CrearMovimientosDetalle(registroMovimientoInversion); err != nil {
				outputError = errorctrl.Error("RegistrarMovimientoDetalleActualizacionInversion - movimientosCrud.CrearMovimientosDetalle(registroMovimientoInversion)", err, "502")
				logs.Error(outputError)
				return outputError
			}
		}
	}

	return nil
}
