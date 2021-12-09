package models

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/plan_adquisiciones_mid/helpers/utils"
	"github.com/udistrital/utils_oas/errorctrl"
)

// INICIO Movimientos Procesos Externos
// ObtenerMovimientoProcesoExterno construye la estructura para registrar el respectivo Movimiento Proceso Externo
func ObtenerMovimientoProcesoExterno(idPlanAdqusiciones int) (registroMovimientoProcesoExternoRespuesta models_movimientosCrud.MovimientoProcesoExterno, outputError interface{}) {
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

// FIN Movimientos Procesos Externos

// INICIO Registro Múltiple Rubros de Inversión
// ObtenerRegistroMovimientoInversion obtiene la estructura de registro múltiple para rubros de inversión
func ObtenerRegistroMovimientoInversion(
	registroPlanAdquisiciones map[string]interface{},
	idMovimientoExterno int,
) (
	registroMovimientosInversionRespuesta []models_movimientosCrud.CuentasMovimientoProcesoExterno,
	outputError interface{},
) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "ObtenerRegistroMovimientoInversion - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

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
				Valor:        saldoValor,
			}

			registroMovimientosInversionRespuesta = append(registroMovimientosInversionRespuesta, registroTemporal)
		}
	}

	return registroMovimientosInversionRespuesta, nil
}

// FIN Registro Múltiple Rubros de Inversión

// INICIO Registro Múltiple Rubros de Funcionamiento
// ObtenerRegistroMovimientoFuncionamiento obtiene la estructura de registro múltiple para rubros de funcionamiento
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
		Valor:        saldoValor,
	}

	registroMovimientosFuncionamientoRespuesta = append(registroMovimientosFuncionamientoRespuesta, registroTemporal)

	return registroMovimientosFuncionamientoRespuesta, nil
}

// FIN Registro Múltiple Rubros de Funcionamiento
