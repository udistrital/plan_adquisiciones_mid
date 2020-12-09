package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_adquisiciones_mid/models"
)

// Registro_plan_adquisicionController operations for Registro_plan_adquisicion
type Registro_plan_adquisicionController struct {
	beego.Controller
}

// URLMapping ...
func (c *Registro_plan_adquisicionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Registro_plan_adquisicionController
// @Param	body		body 	models.Registro_plan_adquisicion	true		"body for Registro_plan_adquisicion content"
// @Success 201 {object} models.Registro_plan_adquisicion
// @Failure 403 body is empty
// @router / [post]
func (c *Registro_plan_adquisicionController) Post() {
	var registroPlanAdquisicionRecibida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroPlanAdquisicionRecibida); err == nil {

		registroPlanAdquisicion, errPlanAdquisicion := models.IngresoPlanAdquisicion(registroPlanAdquisicionRecibida)

		if registroPlanAdquisicion != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = registroPlanAdquisicion
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

// GetOne ...
// @Title GetOne
// @Description get Registro_plan_adquisicionController by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Registro_plan_adquisicion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Registro_plan_adquisicionController) GetOne() {

	planAdquisicionID := c.Ctx.Input.Param(":id")
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	RegistroPlanAdquisicion, errRegistroPlanAdquisicion := models.ObtenerRenglonRegistroPlanAdquisicionByID(planAdquisicionID)

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

// GetAll Regresa todos Registros_plan_adquisicion ordenados por rubro
// @Title GetAll
// @Description Obtiene todos los registros de planes de adquisicion
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

// Put ...
// @Title Put
// @Description update the Registro_plan_adquisicionController
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Registro_plan_adquisicion	true		"body for Registro_plan_adquisicion content"
// @Success 200 {object} models.Registro_plan_adquisicion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Registro_plan_adquisicionController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Registro_plan_adquisicionController
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Registro_plan_adquisicionController) Delete() {

}
