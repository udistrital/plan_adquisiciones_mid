package movimientosCrud

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/plan_adquisiciones_mid/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

func CrearMovimientosDetallePlan(planAdquisiconesId int, planAdquisicionesMongoId string) (movimientosDetalleInsertados []models_movimientosCrud.MovimientoDetalle, outputError map[string]interface{}) {
	// logs.Debug("planAdquisiconesId: ", planAdquisiconesId)

	detalle, err := json.Marshal(models.DetalleMovimientoProcesoExterno{
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

	// logs.Debug("movimientosDetalle: ")
	// formatdata.JsonPrint(movimientosDetalle)

	movimientosDetalleInsertados, outputError = CrearMovimientosDetalle(movimientosDetalle)
	if outputError != nil {
		// logs.Debug(fmt.Sprintf("outputError: %+v", outputError))
		return nil, outputError
	}

	// logs.Debug("movimientosDetalleInsertados: ")
	// formatdata.JsonPrint(movimientosDetalleInsertados)

	return
}
