package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/utils"
)

// ObtenerMovimientoProcesoExterno construye la estructura para registrar el respectivo Movimiento Proceso Externo
// * Función repetida en models, se pone en helpers para mejorar la estructura del proyecto y no tener importaciones recursivas
// * Pendiente revisar para eliminar alguna repetida
func ObtenerMovimientoProcesoExterno(idPlanAdqusiciones int) (registroMovimientoProcesoExternoRespuesta models_movimientosCrud.MovimientoProcesoExterno, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "ObtenerMovimientoProcesoExterno - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	// * Estado por defecto preliminar
	// ? Parametrizable o administrable
	detalle, _ := utils.Serializar(map[string]interface{}{
		"PlanAdquisicionesId": strconv.Itoa(idPlanAdqusiciones),
		"Estado":              "Preliminar",
	})

	procesoExterno, err := strconv.ParseInt(beego.AppConfig.String("procesoExternoPlanPublicado"), 10, 64)
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "ObtenerMovimientoProcesoExterno - strconv.Atoi(beego.AppConfig.String(\"procesoExternoPlanPublicado\"))",
			"err":     err,
			"status":  "500",
		}
		return models_movimientosCrud.MovimientoProcesoExterno{}, outputError
	}

	tipoMovimientoId, err := strconv.Atoi(beego.AppConfig.String("tipoMovimientoIdAfectacionCuenPre"))
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "ObtenerMovimientoProcesoExterno - strconv.Atoi(beego.AppConfig.String(\"tipoMovimientoIdAfectacionCuenPre\"))",
			"err":     err,
			"status":  "500",
		}
		return models_movimientosCrud.MovimientoProcesoExterno{}, outputError

	} else {
		registroTipoMovimientoId := models_movimientosCrud.TipoMovimiento{Id: tipoMovimientoId}
		registroMovimientoProcesoExterno := models_movimientosCrud.MovimientoProcesoExterno{
			Activo:                   true,
			Detalle:                  detalle,
			MovimientoProcesoExterno: 0,
			ProcesoExterno:           procesoExterno,
			TipoMovimientoId:         &registroTipoMovimientoId,
		}

		return registroMovimientoProcesoExterno, nil
	}
}
