package models

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerProductoByID regresa los elementos de la tabla productos
func ObtenerProductoByID(idStr string) (InfoProducto map[string]interface{}, outputError interface{}) {
	var producto map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_cuentas_mongo_crud_url")+"producto/"+idStr, &producto)

	if error != nil {
		return nil, error
	} else {
		if producto["Body"] == nil {
			error := "No se encontro Producto "
			return nil, error
		}
		m := producto["Body"].(interface{})
		Producto := m.(map[string]interface{})
		return Producto, nil
	}

}
