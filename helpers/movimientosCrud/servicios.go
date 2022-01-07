package movimientosCrud

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
	"github.com/udistrital/utils_oas/request"
)

func GetMovimientoProcesoExterno(query string, fields string, sortby string, order string, offset string, limit string) (ml interface{}, err map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("GetMovimientoProcesoExterno - Unhandled error!", "500")

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
		outputError := errorctrl.Error("GetMovimientoProcesoExterno - request.GetJson(urlQuery, &ml)", err, "500")
		return nil, outputError
	}

	// logs.Info(ml)

	if ml == nil {
		return []interface{}{}, nil
	} else {
		return ml.(map[string]interface{})["Body"], err
	}
}

func CrearMovimientosPublicacion(idMovProcExt int) (movimientosDetalleRespuesta []models.MovimientoDetalle, err map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("CrearMovimientosPublicacion - Unhandled error!", "500")

	urlPublicar := beego.AppConfig.String("movimientos_api_crud_url") +
		"movimiento_detalle/publicarMovimientosDetalle"

	if err := request.SendJson(urlPublicar, "POST", &movimientosDetalleRespuesta, idMovProcExt); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("CrearMovimientosPublicacion - request.SendJson(urlPublicar, \"POST\", &movimientosDetalleRespuesta, idMovProcExt)", err, "500")
		return nil, outputError
	}

	return movimientosDetalleRespuesta, nil
}

func CrearMovimientosDetalle(insertarMovimientos []models.CuentasMovimientoProcesoExterno) (movimientosDetalleRespuesta interface{}, err map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("CrearMovimientosDetalle - Unhandled error!", "500")

	urlPublicar := beego.AppConfig.String("movimientos_api_crud_url") +
		"movimiento_detalle/crearMovimientosDetalle"

	if err := request.SendJson(urlPublicar, "POST", &movimientosDetalleRespuesta, insertarMovimientos); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("CrearMovimientosDetalle - request.SendJson(urlPublicar, \"POST\", &movimientosDetalleRespuesta, insertarMovimientos)", err, "500")
		return nil, outputError
	} else {
		// logs.Debug(movimientosDetalleRespuesta)
	}

	return movimientosDetalleRespuesta, nil
}

func CrearMovimientoProcesoExterno(insertarMovimiento models.MovimientoProcesoExterno) (movimientoProcesoExternoRespuesta interface{}, err map[string]interface{}) {
	// logs.Debug(insertarMovimiento)
	defer errorctrl.ErrorControlFunction("CrearMovimientoProcesoExterno - Unhandled error!", "500")

	urlPublicar := beego.AppConfig.String("movimientos_api_crud_url") +
		"movimiento_proceso_externo"

	if err := request.SendJson(urlPublicar, "POST", &movimientoProcesoExternoRespuesta, insertarMovimiento); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("CrearMovimientoProcesoExterno - request.SendJson(urlPublicar, \"POST\", &movimientoProcesoExternoRespuesta, insertarMovimiento)", err, "500")
		return nil, outputError
	} else {
		// logs.Debug(movimientoProcesoExternoRespuesta)
	}

	return movimientoProcesoExternoRespuesta, nil
}
