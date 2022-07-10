package movimientosCrud

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_adquisiciones_mid/models"
	"github.com/udistrital/utils_oas/errorctrl"
	"github.com/udistrital/utils_oas/formatdata"
)

func CrearMovimientosDetallePlan(planAdquisiconesId int, planAdquisicionesMongoId string) (res map[string]interface{}, outputError map[string]interface{}) {

	logs.Debug("planAdquisiconesId: ", planAdquisiconesId)

	detalle, err := json.Marshal(models.DetalleMovimientoProcesoExterno{
		Estado:              "Publicado",
		PlanAdquisicionesId: planAdquisiconesId,
	})
	if err != nil {
		logs.Error(err)
		outputError = errorctrl.Error("CrearMovimientosDetallePlan - json.Marshal(models.DetalleMovimientoProcesoExterno...)", err, "500")
		return nil, outputError
	}

	movimientoProcesoExternoRespuesta, outputError := CrearMovimientoProcesoExterno(detalle)
	if outputError != nil {
		return nil, outputError
	}

	movimientosDetalle, err := AñadirDatosMovimientosDetalle(planAdquisicionesMongoId, movimientoProcesoExternoRespuesta.Id)
	if err != nil {
		logs.Error(err)
		outputError = errorctrl.Error("CrearMovimientosDetallePlan - AñadirDatosMovimientosDetalle()", err, "500")
		return nil, outputError
	}

	logs.Debug("movimientosDetalle: ")
	formatdata.JsonPrint(movimientosDetalle)

	respuesta, outputError := CrearMovimientosDetalle(movimientosDetalle)
	if outputError != nil {
		return nil, outputError
	}

	logs.Debug("respuesta: ")
	formatdata.JsonPrint(respuesta)

	return
}
