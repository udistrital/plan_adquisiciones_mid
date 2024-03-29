package helpers

import (
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/errorctrl"
)

// KeysMovimientoProcesoExterno se encarga de verificar el tipo de dato recibido de la consulta de Movimiento Proceso Externo
func KeysMovimientoProcesoExterno(movimientoResultado interface{}) (movimientoObtenido []interface{}, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("KeysMovimientoProcesoExterno - Unhandled Error!", "500")

	switch movimientoResultado.(type) {
	case []interface{}:
		// logs.Debug("Tipo: ", reflect.TypeOf(resultado))
		if len(movimientoResultado.([]interface{})) > 0 {
			// logs.Debug("Traje información")
			movimientoObtenido = movimientoResultado.([]interface{})
		} else {
			// logs.Debug("No encontré información")
			movimientoObtenido = movimientoResultado.([]interface{})
		}
	case nil:
		// logs.Debug("Tipo: ", reflect.TypeOf(resultado))
		err := "No se encontraron resultados"
		logs.Error(err)
		outputError = errorctrl.Error("IngresoPlanAdquisicion - resultado.(type)", err, "404")
		return nil, outputError
	default:
		// logs.Debug("Tipo: ", reflect.TypeOf(resultado))
		err := "La variable resultado no tiene un tipo de dato coherente"
		logs.Error(err)
		outputError = errorctrl.Error("IngresoPlanAdquisicion - resultado.(type)", err, "409")
		return nil, outputError
	}

	return movimientoObtenido, nil
}
