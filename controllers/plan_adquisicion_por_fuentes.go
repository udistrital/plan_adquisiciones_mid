package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/plan_adquisiciones_mid/models"
)

// Plan_adquisicion_por_fuentesController operations for Plan_adquisicion_por_fuentes
type Plan_adquisicion_por_fuentesController struct {
	beego.Controller
}

// URLMapping ...
func (c *Plan_adquisicion_por_fuentesController) URLMapping() {

	c.Mapping("GetAll", c.GetAll)

}

// GetAll Regresa los Registros_plan_adquisicion separados por ID_fuente y ordenados por ID_rubro
// @Title GetAll
// @Description Obtiene todos los planes de adquisición separados por fuente de recurso y divididos por rubros segun el id dado
// @Param	planAdquisicionID	path 	string	"Id del plan_de_adquisicion"
// @Success 200 {object} models.Plan_adquisicion_por_fuentes
// @Failure 403
// @router /:planAdquisicionID [get]
func (c *Plan_adquisicion_por_fuentesController) GetAll() {

	planAdquisicionID := c.Ctx.Input.Param(":planAdquisicionID")
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	RegistroPlanAdquisicion, errRegistroPlanAdquisicion := models.ObtenerRegistroPlanAdquisicionByIDplan(planAdquisicionID)

	if RegistroPlanAdquisicion != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = RegistroPlanAdquisicion
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, errRegistroPlanAdquisicion)
		alertErr.Body = alertas
		// c.Ctx.Output.SetStatus(404)
	}
	// logs.Debug(alertErr.Body)
	c.Data["json"] = alertErr
	c.ServeJSON()

}
