package models

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerRubroByID regresa los elementos del rubro
func ObtenerRubroByID(idstr string, Vigencia string, UnidadEjecutora string) (InfoRubro map[string]interface{}, outputError interface{}) {
	var rubro map[string]interface{}
	// error := request.GetJson(beego.AppConfig.String("plan_cuentas_mongo_crud_url")+"arbol_rubro_apropiacion/"+idstr+"/"+Vigencia+"/"+UnidadEjecutora+"/", &rubro)
	error := request.GetJson(beego.AppConfig.String("plan_cuentas_mongo_crud_url")+"arbol_rubro_apropiacion/arbol_apropiacion_valores/"+UnidadEjecutora+"/"+Vigencia+"/"+idstr+"?nivel=0", &rubro)
	if error != nil {
		return nil, error
	} else {
		if rubro["Body"] == nil {
			error := "No se encontro Rubro"
			return nil, error
		}
		m := rubro["Body"].([]interface{})
		Rubro := m[0].(map[string]interface{})
		return Rubro["data"].(map[string]interface{}), nil
	}

}

//ObtenerRubroByID regresa los elementos del rubro
func ObtenerFuenteReducidaByID(CodigoFuente string) (InfoFuente map[string]interface{}, outputError interface{}) {
	var fuente []map[string]interface{}
	// error := request.GetJson(beego.AppConfig.String("plan_cuentas_mongo_crud_url")+"arbol_rubro_apropiacion/"+idstr+"/"+Vigencia+"/"+UnidadEjecutora+"/", &rubro)
	error := request.GetJson(beego.AppConfig.String("plan_cuentas_mongo_crud_url")+"arbol_rubro/arbol_reducido/"+CodigoFuente+"?nivel=0", &fuente)
	if error != nil {
		return nil, error
	} else {
		FuenteData := fuente[0]["data"].(map[string]interface{})
		return FuenteData, nil
	}

}

//ObtenerFuenteRecursoByIDRubro regresa los elementos de la fuente de recursos
func ObtenerFuenteRecursoByIDRubro(idstr string, Vigencia string, UnidadEjecutora string) (InfoFuenterecuerso map[string]interface{}, outputError interface{}) {
	var fuenteRecurso []map[string]interface{}
	fuentes, errFuente := SeparaFuentes(idstr)
	if errFuente != nil {
		return nil, errFuente
	}
	error := request.GetJson(beego.AppConfig.String("plan_cuentas_mongo_crud_url")+"/arbol_rubro/arbol/"+fuentes, &fuenteRecurso)

	if error != nil {
		return nil, error
	} else {
		if fuenteRecurso[0]["data"] == nil {
			error := "No se encontro Fuente de Recurso"
			return nil, error
		}
		m := fuenteRecurso[0]["data"].(interface{})
		FuenteRecurso := m.(map[string]interface{})
		return FuenteRecurso, nil
	}

}
