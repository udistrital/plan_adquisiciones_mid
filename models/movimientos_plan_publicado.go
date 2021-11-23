package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/formatdata"
)

// ObtenerRegistrosMasivosPlan crea la estructura de registración masiva al publicar un plan de adquisciones
func ObtenerRegistrosMasivosPlan(idStr string) (respuestaMovimientosPlanPublicado interface{}, outputError interface{}) {
	MovimientosMasivosPlan := make(map[string]interface{})
	registros := make([]map[string]interface{}, 0)
	PlanAdquisicion, error := ObtenerPlanAdquisicionByID(idStr)
	if error != nil || !PlanAdquisicion["Publicado"].(bool) {
		return MovimientosMasivosPlan, error
	} else {
		RegistrosID, error := ObtenerIDRegistrosPlanAdquisicion(idStr)
		if error != nil {
			return nil, error
		} else {
			for _, index := range RegistrosID {
				id := fmt.Sprintf("%.0f", index["Id"].(float64))
				RegistroPlanAdquisicion, _ := ObtenerRenglonRegistroPlanAdquisicionByID(id)
				logs.Debug(fmt.Sprintf("RegistroPlanAdquisicion: %v+", RegistroPlanAdquisicion))
				if RegistroPlanAdquisicion[0]["FuenteFinanciamientoId"] == "" {
					for i := range RegistroPlanAdquisicion[0]["registro_plan_adquisiciones-actividad"].([]map[string]interface{}) {
						idActividad := fmt.Sprintf("%.0f", RegistroPlanAdquisicion[0]["registro_plan_adquisiciones-actividad"].([]map[string]interface{})[i]["ActividadId"].(float64))
						InfoActividad, _ := ObtenerActividadbyID(idActividad)
						RegistroPlanAdquisicion[0]["registro_plan_adquisiciones-actividad"].([]map[string]interface{})[i]["actividad"] = InfoActividad[0]
					}
					EliminarCampos(RegistroPlanAdquisicion[0]["registro_plan_adquisiciones-actividad"].([]map[string]interface{}), "ActividadId")
					EliminarCampos(RegistroPlanAdquisicion, "PlanAdquisicionesId")
					registros = append(registros, RegistroPlanAdquisicion[0])
				} else {
					registros = append(registros, RegistroPlanAdquisicion[0])
				}
			}
			movimientosPlan, errorMovimientos := RegistrarMovimientosMasivosPlan(MovimientosMasivosPlan)
			if errorMovimientos != nil {
				return nil, errorMovimientos
			}
			return movimientosPlan, nil
		}
	}
}

func RegistrarMovimientosMasivosPlan(registroPlanAdquisicion map[string]interface{}) (PlanAdquisicionRespuesta interface{}, outputError interface{}) {
	// PlanAdquisicionPost := make(map[string]interface{})
	// error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Plan_adquisiciones_mongo/", "POST", &PlanAdquisicionPost, registroPlanAdquisicion)
	return "Movimientos Registrados", nil
	/*if error != nil {
		return nil, error
	} else {
		return "Copia Generada", nil
	}*/
}

func ObtenerRegistroMovimientoInversion(registroMovimientoInversion map[string]interface{}) (registroMovimientoInversionRespuesta []map[string]interface{}, outputError interface{}) {
	/*registroMovimientoInversionPost := make(map[string]interface{})
	result := []map[string]interface{}{}

	registroMovimientoInversionPost = map[string]interface{

	}*/
	// logs.Debug(fmt.Sprintf("registroMovimientoInversion: %v", registroMovimientoInversion))
	formatdata.JsonPrint(registroMovimientoInversion["RegistroPlanAdquisicionActividad"].([]interface{})[0].(map[string]interface{})["FuentesFinanciamiento"].([]interface{})[0].(map[string]interface{})["ValorAsignado"].(float64))
	// logs.Debug(fmt.Sprintf("registroMovimientoInversion[\"RegistroPlanAdquisicionesActividad\"].([]map[string]interface{}): %v", registroMovimientoInversion["RegistroPlanAdquisicionActividad"].([]map[string]interface{})))|
	// prueba := registroMovimientoInversion["RegistroPlanAdquisicionesActividad"].(map[string]interface{})
	/*for i := range registroMovimientoInversion["RegistroPlanAdquisicionActividad"].([]map[string]interface{}) {
		for j := range registroMovimientoInversion["RegistroPlanAdquisicionActividad"].([]map[string]interface{}) {
			logs.Info(i, j)
		}
	}*/

	return registroMovimientoInversionRespuesta, nil
}
