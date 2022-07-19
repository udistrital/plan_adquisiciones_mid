package movimientosCrud

import (
	"github.com/astaxie/beego/logs"
	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

func CrearMovimientosDetallePlan(versionPlan map[string]interface{}) (movimientosDetalleInsertados []models_movimientosCrud.MovimientoDetalle, outputError map[string]interface{}) {
	// logs.Debug("planAdquisiconesId: ", planAdquisiconesId)

	// logs.Debug("versionPlan: ")
	// formatdata.JsonPrint(versionPlan)

	movimientosDetalle, err := AñadirDatosMovimientosDetalle(versionPlan)
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
