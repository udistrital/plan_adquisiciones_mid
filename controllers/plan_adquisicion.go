package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_adquisiciones_mid/models"
)

// Plan_adquisicionController operations for Plan_adquisicion
type Plan_adquisicionController struct {
	beego.Controller
}

// URLMapping ...
func (c *Plan_adquisicionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)

}

// Post Funcion para ingresar una version de un plan de adquisicion a mongo si su campo publicado es true
// @Title Create
// @Description create Plan_adquisicion
// @Param	id		path 	string	true		"Id del registro de plan de adquisicion"
// @Success 201 {object} models.Plan_adquisicion
// @Failure 403 body is empty
// @router /:id [post]
func (c *Plan_adquisicionController) Post() {
	idStr := c.Ctx.Input.Param(":id")
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	PlanAdquisicionRespuesta, errPlanAdquisicion := models.ObtenerPlanAdquisicionMongo(idStr)

	if PlanAdquisicionRespuesta != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = PlanAdquisicionRespuesta
		//alertErr.Body = models.CrearSuccess("Registro de actividad ingresado con exito")
	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, errPlanAdquisicion)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()

}

// Put Funcion para actualizar campo publicado del plan de adquisicion e ingresarlo a mongo en caso de que el campo sea true
// @Title Put
// @Description update the Plan_adquisicion
// @Param	id		path 	string	true		"ID del plan de adquisicion a actualizar"
// @Param	body		body 	models.Plan_adquisicion	true		"body for Plan_adquisicion content"
// @Success 200 {object} models.Plan_adquisicion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Plan_adquisicionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	var PlanAdquisicionRecibida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &PlanAdquisicionRecibida); err == nil {

		PlanAdquisicionRespuesta, errPlanAdquisicion := models.ActualizarPlanAdquisicion(PlanAdquisicionRecibida, idStr)

		if PlanAdquisicionRespuesta != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = PlanAdquisicionRespuesta
			//alertErr.Body = models.CrearSuccess("Registro de actividad ingresado con exito")
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, errPlanAdquisicion)
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
