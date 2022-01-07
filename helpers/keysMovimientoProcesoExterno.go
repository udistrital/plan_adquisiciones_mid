package helpers

import "github.com/udistrital/utils_oas/errorctrl"

func KeysMovimientoProcesoExterno(movimientoResultado interface{}) (keysObtenidas []string, movimientoObtenido map[string]interface{}, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("KeysMovimientoProcesoExterno - Unhandled Error!", "500")

	switch movimientoResultado.(type) {
	case map[string]interface{}:
		movimientoObtenido = movimientoResultado.(map[string]interface{})
	case []interface{}:
		// logs.Debug("Tipo: ", reflect.TypeOf(resultado))
		movimientoObtenido = movimientoResultado.([]interface{})[0].(map[string]interface{})
	case nil:
		// logs.Debug("Tipo: ", reflect.TypeOf(resultado))
	default:
		// logs.Debug("Tipo: ", reflect.TypeOf(resultado))
	}

	keys := make([]string, 0, len(movimientoObtenido))
	for k := range movimientoObtenido {
		keys = append(keys, k)
	}

	keysObtenidas = keys
	return keysObtenidas, movimientoObtenido, nil
}
