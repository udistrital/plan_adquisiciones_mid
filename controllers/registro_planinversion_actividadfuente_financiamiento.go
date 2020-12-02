package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_adquisiciones_mid/models"
)

// Registro_PlanInversion_ActividadFuente_financiamientoController operations for Registro_PlanInversion_ActividadFuente_financiamiento
type Registro_PlanInversion_ActividadFuente_financiamientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *Registro_PlanInversion_ActividadFuente_financiamientoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)

}

// Post ...
// @Title Create
// @Description create Registro_PlanInversion_ActividadFuente_financiamiento
// @Param	body		body 	models.Registro_PlanInversion_ActividadFuente_financiamiento	true		"body for Registro_PlanInversion_ActividadFuente_financiamiento content"
// @Success 201 {object} models.Registro_PlanInversion_ActividadFuente_financiamiento
// @Failure 403 body is empty
// @router / [post]
func (c *Registro_PlanInversion_ActividadFuente_financiamientoController) Post() {

	var registroActividadFuenteRecibida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroActividadFuenteRecibida); err == nil {

		registroActividadFuenteRespuesta, errRegistroActividadFuente := models.IngresoRegistroPlanInversionActividadFuente(registroActividadFuenteRecibida)

		if registroActividadFuenteRespuesta != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = registroActividadFuenteRespuesta
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, errRegistroActividadFuente)
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get Registro_PlanInversion_ActividadFuente_financiamiento
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Registro_PlanInversion_ActividadFuente_financiamiento
// @Failure 403
// @router / [get]
func (c *Registro_PlanInversion_ActividadFuente_financiamientoController) GetAll() {

	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	RegistroPlanInversionActividadFuente, errRegistroPlanInversionActividadFuente := models.ObtenerRegistroPlanInversionActividadFuente()

	if RegistroPlanInversionActividadFuente != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = RegistroPlanInversionActividadFuente
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, errRegistroPlanInversionActividadFuente)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Registro_PlanInversion_ActividadFuente_financiamiento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Registro_PlanInversion_ActividadFuente_financiamiento	true		"body for Registro_PlanInversion_ActividadFuente_financiamiento content"
// @Success 200 {object} models.Registro_PlanInversion_ActividadFuente_financiamiento
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Registro_PlanInversion_ActividadFuente_financiamientoController) Put() {

	idStr := c.Ctx.Input.Param(":id")
	var registroActividadFuenteRecibida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroActividadFuenteRecibida); err == nil {

		registroActividadFuenteRespuesta, errRegistroActividadFuente := models.ActualizarRegistroActividadFuente(registroActividadFuenteRecibida, idStr, idStr)

		if registroActividadFuenteRespuesta != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = registroActividadFuenteRespuesta
			//alertErr.Body = models.CrearSuccess("Registro de actividad ingresado con exito")
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, errRegistroActividadFuente)
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()

}
