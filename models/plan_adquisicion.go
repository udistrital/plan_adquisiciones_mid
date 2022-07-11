package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerPlanAdquisicionByID regresa un de plan de adquisicion segun ID
func ObtenerPlanAdquisicionByID(idstr string) (respuestaPlanAdquisicion map[string]interface{}, outputError interface{}) {
	var PlanAdquisicion map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Plan_adquisiciones/"+idstr, &PlanAdquisicion)
	if error != nil {
		return nil, error
	} else {
		if PlanAdquisicion["Message"] == "Not found resource" {
			error := "No existe plan adquisicion"
			return nil, error
		}
		return PlanAdquisicion, nil
	}
}

//ObtenerVersionesMongoByID regresa Versiones almacenadas en Mongo segun ID plan adquisicion
func ObtenerVersionesMongoByID(idstr string) (respuestaVersionesMongo []map[string]interface{}, outputError interface{}) {
	var versionesMongo []map[string]interface{}

	mongoIDs := make([]map[string]interface{}, 0)
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+`Plan_adquisiciones_mongo/?query={"id":`+idstr+`}`, &versionesMongo)
	if error != nil {
		return nil, error
	} else {
		if versionesMongo[0]["Message"] == "Not found resource" {
			error := "No existen versiones del plan adquisicion"
			return nil, error
		}
		for index := range versionesMongo {
			mongoID := make(map[string]interface{})
			mongoID["_id"] = versionesMongo[index]["_id"]
			mongoID["id"] = versionesMongo[index]["id"]
			mongoID["fechacreacion"] = versionesMongo[index]["fechacreacion"]
			mongoIDs = append(mongoIDs, mongoID)
		}
		return mongoIDs, nil
	}
}

//ActualizarPlanAdquisicion actualizar los campo Publicado de la tabla plan de adquisicion
func ActualizarPlanAdquisicion(PlanAdquisicion map[string]interface{}, idStr string) (PlanAdquisionRespuesta interface{}, outputError interface{}) {
	PlanAdquisicionPut := make(map[string]interface{})
	PlanAdquisicionActualizar := make(map[string]interface{})
	PlanAdquisicionAntiguo, error := ObtenerPlanAdquisicionByID(idStr)
	if error != nil {
		return PlanAdquisicionActualizar, nil
	} else {
		PlanAdquisicionActualizar = map[string]interface{}{
			"Descripcion":   PlanAdquisicionAntiguo["Descripcion"],
			"Vigencia":      PlanAdquisicionAntiguo["Vigencia"],
			"FechaCreacion": PlanAdquisicionAntiguo["FechaCreacion"],
			"Activo":        PlanAdquisicionAntiguo["Activo"],
			"Publicado":     PlanAdquisicion["Publicado"],
		}
		error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Plan_adquisiciones/"+idStr, "PUT", &PlanAdquisicionPut, PlanAdquisicionActualizar)
		// logs.Debug("Plan PUT", PlanAdquisicionPut)
		if error != nil {
			return nil, error
		} else {
			PlanAdquisicionMongo, error := ObtenerPlanAdquisicionMongo(idStr)
			if error != nil {
				return PlanAdquisicionMongo, error
			} else {
				// logs.Debug("Plan Mongo", PlanAdquisicionMongo)
				return PlanAdquisicionMongo, nil
			}

		}

	}

}

//ObtenerFichaEBMGAByIDPlan regresa Ficha_EB_IMGA segun ID plan adquisicion
func ObtenerFichaEBMGAByIDPlan(idstr string) (respuestaFichaEBMGA []map[string]interface{}, outputError interface{}) {
	var FichaEBMGA []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Ficha_EB_IMGA/?query=PlanAdquisicionesId.id:"+idstr, &FichaEBMGA)
	if error != nil {
		return nil, error
	} else {
		//EliminarCampos(FichaEBMGA, "PlanAdquisicionesId")
		return FichaEBMGA, nil
	}
}

//ObtenerIDRegistrosPlanAdquisicion regresa los Id de los registros_plan_adquisicion asociados a un ID plan adquisicion
func ObtenerIDRegistrosPlanAdquisicion(idstr string) (respuestaRegistroID []map[string]interface{}, outputError interface{}) {
	var RegistroID []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_plan_adquisiciones/?query=PlanAdquisicionesId.id:"+idstr+"&fields=Id&sortby=RubroId&order=asc", &RegistroID)
	if error != nil {
		return nil, error
	} else {
		return RegistroID, nil
	}
}

//ObtenerActividadbyID regresa la informacion relacionada a la tabla de actividad, meta, lineamiento
func ObtenerActividadbyID(idstr string) (respuestaActividad []map[string]interface{}, outputError interface{}) {
	var Actividad []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Actividad/?query=Id:"+idstr, &Actividad)
	if error != nil {
		return nil, error
	} else {
		return Actividad, nil
	}
}

//ObtenerPlanAdquisicionMongo construye un de plan de adquisicion segun ID con el formato Json plan_adquisiciones_mongo
func ObtenerPlanAdquisicionMongo(idStr string) (respuestaPlanAdquisicionMongo interface{}, outputError interface{}) {
	PlanAdquisicionMongo := make(map[string]interface{})
	registros := make([]map[string]interface{}, 0)
	PlanAdquisicion, error := ObtenerPlanAdquisicionByID(idStr)
	if error != nil || !PlanAdquisicion["Publicado"].(bool) {
		return PlanAdquisicionMongo, error
	} else {
		PlanAdquisicionMongo = PlanAdquisicion
		FichaEBMGA, error := ObtenerFichaEBMGAByIDPlan(idStr)
		if error != nil {
			return nil, error
		} else {
			PlanAdquisicionMongo["ficha_eb_imga"] = FichaEBMGA
			RegistrosID, error := ObtenerIDRegistrosPlanAdquisicion(idStr)
			if error != nil {
				return nil, error
			} else {
				for _, index := range RegistrosID {
					id := fmt.Sprintf("%.0f", index["Id"].(float64))
					RegistroPlanAdquisicion, _ := ObtenerRenglonRegistroPlanAdquisicionByID(id)
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
				registrosSperados, _ := SepararRegistrosPorFuente(registros)
				PlanAdquisicionMongo["registro_plan_adquisiciones"] = registrosSperados
				planMongo, erroMOngo := IngresoPlanAdquisicionMongo(PlanAdquisicionMongo)
				if erroMOngo != nil {
					return nil, erroMOngo
				}
				return planMongo, nil

			}

		}

	}
}

//IngresoPlanAdquisicionMongo crea una copia del plan de adquisicion en mongo
func IngresoPlanAdquisicionMongo(registroPlanAdquisicion map[string]interface{}) (PlanAdquisicionRespuesta interface{}, outputError interface{}) {
	PlanAdquisicionPost := make(map[string]interface{})
	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Plan_adquisiciones_mongo/", "POST", &PlanAdquisicionPost, registroPlanAdquisicion)
	if error != nil {
		return nil, error
	} else {
		// logs.Debug(fmt.Sprintf("PlanAdquisicionPost: %+v", PlanAdquisicionPost))
		return PlanAdquisicionPost, nil
	}

}
