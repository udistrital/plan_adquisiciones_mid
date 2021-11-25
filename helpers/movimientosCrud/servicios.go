package movimientosCrud

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func GetMovimientoProcesoExterno(
	query string,
	fields string,
	sortby string,
	order string,
	offset string,
	limit string,
) (ml []interface{}, err error) {

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
		"movimiento_proceso_externo?" + params.Encode()

	if err := request.GetJson(urlQuery, &ml); err != nil {
		return nil, err
	}

	return
}
