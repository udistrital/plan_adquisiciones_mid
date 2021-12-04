package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
	"github.com/udistrital/utils_oas/request"
)

// INICIO Movimientos Procesos Externos
// Construir la estructura para registrar el respectivo Movimiento Proceso Externo
func ObtenerMovimientoProcesoExterno(registroPlanAdquisicion map[string]interface{}) (registroMovimientoProcesoExternoRespuesta MovimientoProcesoExternoId, outputError interface{}) {
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

	planAdquisicionesID := registroPlanAdquisicion["PlanAdquisicionesId"].(float64)
	detalle := "{\"PlanAdquisicionesId\":" + fmt.Sprintf("%.0f", planAdquisicionesID) + ", \"Estado\":\"Registrado\"}"
	procesoExterno, err := strconv.Atoi(beego.AppConfig.String("procesoExternoPlanPublicado"))
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "ObtenerMovimientoProcesoExterno - strconv.Atoi(beego.AppConfig.String(\"procesoExternoPlanPublicado\"))",
			"err":     err,
			"status":  "500",
		}
		return MovimientoProcesoExternoId{}, outputError
	}
	tipoMovimientoId, err := strconv.Atoi(beego.AppConfig.String("tipoMovimientoIdAfectacionCuenPre"))
	if err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "ObtenerMovimientoProcesoExterno - strconv.Atoi(beego.AppConfig.String(\"tipoMovimientoIdAfectacionCuenPre\"))",
			"err":     err,
			"status":  "500",
		}
		return MovimientoProcesoExternoId{}, outputError
	} else {
		registroTipoMovimientoId := TipoMovimientoId{Id: tipoMovimientoId}
		registroMovimientoProcesoExterno := MovimientoProcesoExternoId{
			Activo:                   true,
			Detalle:                  detalle,
			MovimientoProcesoExterno: 0,
			ProcesoExterno:           procesoExterno,
			TipoMovimientoId:         registroTipoMovimientoId,
		}
		resultado, err := RegistrarMovimientoProcesoExterno(registroMovimientoProcesoExterno)
		if err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{
				"funcion": "ObtenerMovimientoProcesoExterno - RegistrarMovimientoProcesoExterno(registroMovimientoProcesoExterno)",
				"err":     err,
				"status":  "500",
			}
		} else {
			// logs.Debug(resultado)
		}

		return resultado, outputError
	}
}

// Registrar el Movi	miento Proceso Extorno según la estructura obtenida antes
func RegistrarMovimientoProcesoExterno(registroMovimientoProcesoExterno MovimientoProcesoExternoId) (registroMovimientoProcesoExternoRespuesta MovimientoProcesoExternoId, outputError interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "RegistrarMovimientoProcesoExterno - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	registroMovimientoProcesoExternoIngresado := registroMovimientoProcesoExterno
	var registroMovimientoProcesoExternoPost MovimientoProcesoExternoId
	urlMovimientoProcesoExternoCREATE := beego.AppConfig.String("movimientos_api_crud_url") + "movimiento_proceso_externo/"
	if err := request.SendJson(urlMovimientoProcesoExternoCREATE, "POST", &registroMovimientoProcesoExternoPost, registroMovimientoProcesoExternoIngresado); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "ObtenerMovimientoProcesoExterno - request.SendJson(urlMovimientoProcesoExternoCRUD, \"POST\", &registroMovimientoProcesoExternoPost, registroMovimientoProcesoExternoIngresado)",
			"err":     err,
			"status":  "502",
		}
		return MovimientoProcesoExternoId{}, outputError
	} else {
		return registroMovimientoProcesoExternoPost, outputError
	}
}

// FIN Movimientos Procesos Externos

// INICIO Registro Múltiple Rubros de Inversión
// Obtener la estructura de registro múltiple para rubros de inversión
func ObtenerRegistroMovimientoInversion(registroPlanAdquisiciones map[string]interface{}, idMovimientoExterno int) (registroMovimientosInversionRespuesta []models_movimientosCrud.CuentasMovimientoProcesoExterno, outputError interface{}) {
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

	for i := range registroPlanAdquisicionesActividades {
		for j := range registroPlanAdquisicionesActividades[i].(map[string]interface{})["FuentesFinanciamiento"].([]interface{}) {

			registroMovimientoExternoId := strconv.Itoa(idMovimientoExterno)

			detalleRaw := map[string]interface{}{
				"RubroId":                registroPlanAdquisiciones["RubroId"].(string),
				"ActividadId":            strconv.Itoa(int(registroPlanAdquisicionesActividades[i].(map[string]interface{})["ActividadId"].(float64))),
				"FuenteFinanciamientoId": registroPlanAdquisicionesActividades[i].(map[string]interface{})["FuentesFinanciamiento"].([]interface{})[j].(map[string]interface{})["FuenteFinanciamientoId"].(string),
			}

			detalleCast, _ := json.Marshal(detalleRaw)

			saldoValor := registroPlanAdquisicionesActividades[i].(map[string]interface{})["FuentesFinanciamiento"].([]interface{})[j].(map[string]interface{})["ValorAsignado"].(float64)

			registroTemporal := models_movimientosCrud.CuentasMovimientoProcesoExterno{
				Cuen_Pre:     string(detalleCast),
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
// Obtener la estructura de registro múltiple para rubros de funcionamiento
func ObtenerRegistroMovimientoFuncionamiento(registroPlanAdquisiciones map[string]interface{}, idMovimientoExterno int) (registroMovimientosFuncionamientoRespuesta []models_movimientosCrud.CuentasMovimientoProcesoExterno, outputError interface{}) {
	defer errorctrl.ErrorControlFunction("ObtenerRegistroMovimientoInversion - Unhandled Error!", "500")

	registroMovimientoExternoId := strconv.Itoa(idMovimientoExterno)

	detalleRaw := map[string]interface{}{
		"RubroId":                registroPlanAdquisiciones["RubroId"].(string),
		"FuenteFinanciamientoId": registroPlanAdquisiciones["FuenteFinanciamientoId"].(string),
	}

	detalleCast, _ := json.Marshal(detalleRaw)

	saldoValor := registroPlanAdquisiciones["ValorActividad"].(float64)

	registroTemporal := models_movimientosCrud.CuentasMovimientoProcesoExterno{
		Cuen_Pre:     string(detalleCast),
		Mov_Proc_Ext: registroMovimientoExternoId,
		Valor:        saldoValor,
	}

	registroMovimientosFuncionamientoRespuesta = append(registroMovimientosFuncionamientoRespuesta, registroTemporal)

	/*logs.Debug(registroMovimientosFuncionamientoRespuesta)
	if resultado, err := RegistrarMultiplesMovimientos(registroMovimientosFuncionamientoRespuesta); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "ObtenerRegistroMovimientoInversion - Unhandled Error!",
			"err":     err,
			"status":  "500",
		}
		return nil, err
	} else {
		logs.Debug(resultado)
	}*/

	return registroMovimientosFuncionamientoRespuesta, nil
}

// FIN Registro Múltiple Rubros de Funcionamiento

// INICIO Registrar Múltiples Movimientos
// Insertar dentro de la base de datos los movimientos definidos en funciones anteriores
func RegistrarMultiplesMovimientos(registrosMultiples []RegistrosMultiplesMovimientos) (registrosMultiplesRespuesta RegistrosMultiplesMovimientos, outputError interface{}) {
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

	urlRegistrosMultiplesCREATE := beego.AppConfig.String("movimientos_api_crud_url") + "movimiento_detalle/"

	for i := range registrosMultiples {
		// PlanAdquisicionPost := make(map[string]interface{})
		if err := request.SendJson(urlRegistrosMultiplesCREATE, "POST", &registrosMultiplesRespuesta, registrosMultiples[i]); err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{
				"funcion": "RegistrarMovimientosMasivosPlan - request.SendJson(urlRegistrosMultiplesCREATE, \"POST\", &registrosMultiplesRespuesta, registrosMultiples)",
				"err":     err,
				"status":  "502",
			}
			return RegistrosMultiplesMovimientos{}, outputError
		} else {
			return registrosMultiplesRespuesta, nil
		}
	}

	return registrosMultiplesRespuesta, nil
}

// FIN Registrar Múltiples Movimientos
