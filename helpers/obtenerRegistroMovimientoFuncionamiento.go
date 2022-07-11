package helpers

import (
	"strconv"

	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/utils"
	"github.com/udistrital/utils_oas/errorctrl"
)

// ObtenerRegistroMovimientoFuncionamiento obtiene la estructura de registro múltiple para rubros de funcionamiento
// * Función repetida en models, se pone en helpers para mejorar la estructura del proyecto y no tener importaciones recursivas
// * Pendiente revisar para eliminar alguna repetida
func ObtenerRegistroMovimientoFuncionamiento(
	registroPlanAdquisiciones map[string]interface{},
	idMovimientoExterno int,
) (
	registroMovimientosFuncionamientoRespuesta []models_movimientosCrud.CuentasMovimientoProcesoExterno,
	outputError interface{},
) {
	defer errorctrl.ErrorControlFunction("ObtenerRegistroMovimientoInversion - Unhandled Error!", "500")

	registroMovimientoExternoId := strconv.Itoa(idMovimientoExterno)

	detalle, _ := utils.Serializar(map[string]interface{}{
		"RubroId":                registroPlanAdquisiciones["RubroId"].(string),
		"FuenteFinanciamientoId": registroPlanAdquisiciones["FuenteFinanciamientoId"].(string),
	})

	saldoValor := registroPlanAdquisiciones["ValorActividad"].(float64)

	registroTemporal := models_movimientosCrud.CuentasMovimientoProcesoExterno{
		Cuen_Pre:     detalle,
		Mov_Proc_Ext: registroMovimientoExternoId,
		Saldo:        saldoValor,
	}

	registroMovimientosFuncionamientoRespuesta = append(registroMovimientosFuncionamientoRespuesta, registroTemporal)

	return registroMovimientosFuncionamientoRespuesta, nil
}

// FIN Registro Múltiple Rubros de Funcionamiento
