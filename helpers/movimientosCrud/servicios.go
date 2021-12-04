package movimientosCrud

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/request"
)

func GetMovimientoProcesoExterno(query string, fields string, sortby string, order string, offset string, limit string) (ml interface{}, err error) {

	params := url.Values{}

	if len(query) > 0 {
		params.Add("query", query)
	}

	if len(fields) > 0 {
		params.Add("fields", fields)
	}

	if len(sortby) > 0 {
		params.Add("sortby", sortby)

	}

	if len(order) > 0 {
		params.Add("order", order)

	}

	if len(offset) > 0 {
		params.Add("offset", offset)

	}

	if len(limit) > 0 {
		params.Add("limit", limit)

	}

	urlQuery := beego.AppConfig.String("movimientos_api_crud_url") +
		"movimiento_proceso_externo/movimientoFiltroJsonB?" + params.Encode()

	// logs.Debug(urlQuery)

	if err := request.GetJson(urlQuery, &ml); err != nil {
		logs.Error(err)
		return nil, err
	}

	return ml.(map[string]interface{})["Body"], err
}

func CrearMovimientosPublicacion(idMovProcExt int) (movimientosDetalleRespuesta []models_movimientosCrud.MovimientoDetalle, err error) {
	urlPublicar := beego.AppConfig.String("movimientos_api_crud_url") +
		"movimiento_detalle/publicarMovimientosDetalle"

	if err := request.SendJson(urlPublicar, "POST", &movimientosDetalleRespuesta, idMovProcExt); err != nil {
		logs.Error(err)
		return nil, err
	}

	return
}

func CrearMovimientosDetalle(insertarMovimientos []models_movimientosCrud.CuentasMovimientoProcesoExterno) (movimientosDetalleRespuesta interface{}, err error) {
	urlPublicar := beego.AppConfig.String("movimientos_api_crud_url") +
		"movimiento_detalle/crearMovimientosDetalle"

	if err := request.SendJson(urlPublicar, "POST", &movimientosDetalleRespuesta, insertarMovimientos); err != nil {
		logs.Error(err)
		return nil, err
	} else {
		logs.Debug(movimientosDetalleRespuesta)
	}

	return
}
