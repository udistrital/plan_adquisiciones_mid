package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/plan_adquisiciones_mid/models"
)

// Registro_plan_adquisicionController operations for Registro_plan_adquisicion
type Registro_plan_adquisicionController struct {
	beego.Controller
}

// URLMapping ...
func (c *Registro_plan_adquisicionController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll Regresa todos Registros_plan_adquisicion ordenados por rubro
// @Title GetAll
// @Description get Registro_plan_adquisicion
// @Success 200 {object} models.Registro_plan_adquisicion
// @Failure 404
// @router / [get]
func (c *Registro_plan_adquisicionController) GetAll() {

	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	RegistroPlanAdquisicion, errRegistroPlanAdquisicion := models.ObtenerRegistroPlanAdquisicion()

	if RegistroPlanAdquisicion != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = RegistroPlanAdquisicion
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, errRegistroPlanAdquisicion)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()

}
