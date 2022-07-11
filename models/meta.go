package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerMetaByID regresa los elementos de la tabla meta
func ObtenerMetaByID(idstr string) (InfoMeta map[string]interface{}, outputError interface{}) {
	var meta map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Meta/"+idstr, &meta)
	if error != nil {
		return nil, error
	} else {
		return meta, nil
	}

}

//ObtenerLineamiento regresa los elementos de la tabla lineamiento
func ObtenerLineamiento(idMeta string) (InfoLineamiento []map[string]interface{}, outputError interface{}) {
	var Lineamiento []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Meta/?query=Id:"+idMeta+"&fields=LineamientoId", &Lineamiento)
	if error != nil {
		return nil, error
	} else {
		return Lineamiento, nil
	}

}

func ObtenerActividadById(idActividad interface{}) (Actividad map[string]interface{}, outputError interface{}) {
	var ActividadAsociada []map[string]interface{}
	s := fmt.Sprintf("%.0f", idActividad.(float64))
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Actividad?query=Id:"+s+"&fields=Nombre,Numero", &ActividadAsociada)
	// logs.Debug("ActividadAsociada: ", ActividadAsociada)
	if error != nil {
		return nil, error
	} else {
		if len(ActividadAsociada) > 0 {
			Actividad = ActividadAsociada[0]
		}
		return Actividad, nil
	}

}
