package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_adquisiciones_mid/models"
)

// Registro_PlanAdquisiciones_ActividadController operations for Registro_PlanAdquisiciones_Actividad
type Registro_PlanAdquisiciones_ActividadController struct {
	beego.Controller
}

// URLMapping ...
func (c *Registro_PlanAdquisiciones_ActividadController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post Funcion para ingresar un nuevo registro de actividad a la tabla registro_plan_adquisición_actividad
// @Title Create
// @Description create Registro_PlanAdquisiciones_Actividad
// @Param	body		body 	models.Registro_PlanAdquisiciones_Actividad	true		"body for Registro_PlanAdquisiciones_Actividad content"
// @Success 201 {object} models.Registro_PlanAdquisiciones_Actividad
// @Failure 403 body is empty
// @router / [post]
func (c *Registro_PlanAdquisiciones_ActividadController) Post() {

	var registroActividadRecibida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroActividadRecibida); err == nil {

		registroActividadRespuesta, errRegistroActividad := models.IngresoRegistroActividad(registroActividadRecibida)

		if registroActividadRespuesta != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = registroActividadRespuesta
			//alertErr.Body = models.CrearSuccess("Registro de actividad ingresado con exito")
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, errRegistroActividad)
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

// GetAll Funcion para regresar todos los elementos de la tabla registro_plan_adquisición_actividad
// @Title GetAll
// @Description get Registro_PlanAdquisiciones_Actividad
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Registro_PlanAdquisiciones_Actividad
// @Failure 403
// @router / [get]
func (c *Registro_PlanAdquisiciones_ActividadController) GetAll() {

	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	RegistroPlanAdquisicionActividad, errRegistroPlanAdquisicionActividad := models.ObtenerRegistroPlanAdquisicionActividad()

	if RegistroPlanAdquisicionActividad != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = RegistroPlanAdquisicionActividad
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, errRegistroPlanAdquisicionActividad)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()

}

// Put Función para actualizar un elemento de la tabla registro_plan_adquisición_actividad
// @Title Put
// @Description update the Registro_PlanAdquisiciones_Actividad
// @Param	id		path 	string	true		"Id del registro_plan_adquisición_actividad"
// @Param	body		body 	models.Registro_PlanAdquisiciones_Actividad	true		"body for Registro_PlanAdquisiciones_Actividad content"
// @Success 200 {object} models.Registro_PlanAdquisiciones_Actividad
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Registro_PlanAdquisiciones_ActividadController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	var registroActividadRecibida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroActividadRecibida); err == nil {

		registroActividadRespuesta, errRegistroActividad := models.ActualizarRegistroActividad(registroActividadRecibida, idStr)

		if registroActividadRespuesta != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = registroActividadRespuesta
			//alertErr.Body = models.CrearSuccess("Registro de actividad ingresado con exito")
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, errRegistroActividad)
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
