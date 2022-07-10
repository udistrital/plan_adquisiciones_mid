package movimientosCrud

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	models_movimientosCrud "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/plan_adquisiciones_mid/models"
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

	// logs.Info("ml: ", ml)

	return
}

func CrearMovimientoProcesoExterno(detalle []byte) (movimientoProcesoExternoRespuesta models_movimientosCrud.MovimientoProcesoExterno, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("CrearMovimientoProcesoExterno - Unhandled error!", "500")
	tipoMovimientoId, err := beego.AppConfig.Int("tipoMovimientoIdAfectacionCuenPre")
	if err != nil {
		outputError = errorctrl.Error("CrearMovimientoProcesoExterno - beego.AppConfig.Int(\"tipoMovimientoIdAfectacionCuenPre\")", err, "500")
		return models_movimientosCrud.MovimientoProcesoExterno{}, outputError
	}

	tipoMovimiento := models_movimientosCrud.TipoMovimiento{
		Id: tipoMovimientoId,
	}

	// var movimientoEstructura models.MovimientoProcesoExterno
	nuevoMovimiento := models_movimientosCrud.MovimientoProcesoExterno{
		TipoMovimientoId: &tipoMovimiento,
		Activo:           true,
		Detalle:          string(detalle),
	}

	// logs.Debug("nuevoMovimiento: ", nuevoMovimiento)
	urlPublicar := beego.AppConfig.String("movimientos_api_crud_url") +
		"movimiento_proceso_externo/"

	if err := request.SendJson(urlPublicar, "POST", &movimientoProcesoExternoRespuesta, nuevoMovimiento); err != nil {
		outputError = errorctrl.Error("CrearMovimientoProcesoExterno - request.SendJson(urlPublicar, \"POST\", &movimientoProcesoExternoRespuesta, insertarMovimiento)", err, "500")
		return models_movimientosCrud.MovimientoProcesoExterno{}, outputError
	}

	return movimientoProcesoExternoRespuesta, nil
}

func AñadirDatosMovimientosDetalle(idPlanAdquisicionesMongo string, idMovimientoProcesoExterno int) (movimientosDetalle []models.MovimientosDetalle, err error) {
	urlConsultarId := beego.AppConfig.String("plan_adquicisiones_crud_url") +
		"Plan_adquisiciones_mongo/diferencia/" + idPlanAdquisicionesMongo

	if err := request.GetJson(urlConsultarId, &movimientosDetalle); err != nil {
		logs.Error(err)
		return nil, err
	}

	for keyMovimiento := range movimientosDetalle {
		movimientosDetalle[keyMovimiento].MovimientoProcesoExternoId = idMovimientoProcesoExterno
		movimientosDetalle[keyMovimiento].Activo = true
	}

	return
}

func CrearMovimientosDetalle(insertarMovimientos []models.MovimientosDetalle) (movimientosDetalleRespuesta []models_movimientosCrud.MovimientoDetalle, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("CrearMovimientosDetalle - Unhandled error!", "500")

	urlCrearMovimientos := beego.AppConfig.String("movimientos_api_crud_url") +
		"movimiento_detalle/crearMovimientosDetalle"

	if err := request.SendJson(urlCrearMovimientos, "POST", &movimientosDetalleRespuesta, insertarMovimientos); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("CrearMovimientosDetalle - request.SendJson(urlPublicar, \"POST\", &movimientosDetalleRespuesta, insertarMovimientos)", err, "500")
		return nil, outputError
	} else {
		// logs.Debug(movimientosDetalleRespuesta)
	}

	return movimientosDetalleRespuesta, nil
}
