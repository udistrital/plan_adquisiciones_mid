package helpers

import (
	"strconv"

	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/utils"
	"github.com/udistrital/utils_oas/errorctrl"
)

// ObtenerRegistroMovimientoInversion obtiene la estructura de registro múltiple para rubros de inversión
// * Función repetida en models, se pone en helpers para mejorar la estructura del proyecto y no tener importaciones recursivas
// * Pendiente revisar para eliminar alguna repetida
func ObtenerRegistroMovimientoInversion(
	registroPlanAdquisiciones map[string]interface{},
	idMovimientoExterno int,
) (
	registroMovimientosInversionRespuesta []models_movimientosCrud.CuentasMovimientoProcesoExterno,
	outputError interface{},
) {
	defer errorctrl.ErrorControlFunction("ObtenerRegistroMovimientoInversion - Unhandled Error!", "500")

	registroPlanAdquisicionesActividades := registroPlanAdquisiciones["RegistroPlanAdquisicionActividad"].([]interface{})

	for _, actividad := range registroPlanAdquisicionesActividades {
		for _, fuente := range actividad.(map[string]interface{})["FuentesFinanciamiento"].([]interface{}) {

			registroMovimientoExternoId := strconv.Itoa(idMovimientoExterno)

			detalle, _ := utils.Serializar(map[string]interface{}{
				"RubroId":                registroPlanAdquisiciones["RubroId"].(string),
				"ActividadId":            strconv.Itoa(int(actividad.(map[string]interface{})["ActividadId"].(float64))),
				"FuenteFinanciamientoId": fuente.(map[string]interface{})["FuenteFinanciamientoId"].(string),
			})

			saldoValor := fuente.(map[string]interface{})["ValorAsignado"].(float64)

			registroTemporal := models_movimientosCrud.CuentasMovimientoProcesoExterno{
				Cuen_Pre:     detalle,
				Mov_Proc_Ext: registroMovimientoExternoId,
				Saldo:        saldoValor,
			}

			registroMovimientosInversionRespuesta = append(registroMovimientosInversionRespuesta, registroTemporal)
		}
	}

	return registroMovimientosInversionRespuesta, nil
}
