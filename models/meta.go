package models

import (
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
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Meta/?query=Id%3A"+idMeta+"&fields=LineamientoId", &Lineamiento)
	if error != nil {
		return nil, error
	} else {
		return Lineamiento, nil
	}

}
