package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	helpersCrud "github.com/udistrital/plan_adquisiciones_mid/helpers/movimientosCrud"
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
	c.Mapping("GetOne", c.GetOne)

}

// Post Funcion para ingresar una version de un plan de adquisicion a mongo si su campo publicado es true
// @Title Create
// @Description create Plan_adquisicion
// @Param	id		path 	string	true		"Id del registro de plan de adquisicion"
// @Success 201 {object} models.Plan_adquisicion
// @Failure 403 body is empty
// @router /:id [post]
func (c *Plan_adquisicionController) Post() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "Plan_adquisicionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("500")
			}
		}
	}()

	idStr := c.Ctx.Input.Param(":id")
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	PlanAdquisicionRespuesta, errPlanAdquisicion := models.ObtenerPlanAdquisicionMongo(idStr)

	// formatdata.JsonPrint(PlanAdquisicionRespuesta)

	if PlanAdquisicionRespuesta != nil {
		planAdquisiconesId := int(PlanAdquisicionRespuesta.(map[string]interface{})["PlanAdquisiconesMongo"].(map[string]interface{})["id"].(float64))
		planAdquisiconesIdMongo := PlanAdquisicionRespuesta.(map[string]interface{})["PlanAdquisiconesMongo"].(map[string]interface{})["IdMongo"].(string)

		MovimientosRespuesta, _ := helpersCrud.CrearMovimientosDetallePlan(planAdquisiconesId, planAdquisiconesIdMongo)
		if len(MovimientosRespuesta) != 0 {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = map[string]interface{}{"PlanAdquisiciones": PlanAdquisicionRespuesta, "Movimientos": MovimientosRespuesta}
		}
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

		// formatdata.JsonPrint(PlanAdquisicionRespuesta)

		if PlanAdquisicionRespuesta != nil || alertErr.Type == "error" {
			planAdquisiconesId := int(PlanAdquisicionRespuesta.(map[string]interface{})["PlanAdquisiconesMongo"].(map[string]interface{})["id"].(float64))
			planAdquisiconesIdMongo := PlanAdquisicionRespuesta.(map[string]interface{})["IdMongo"].(string)

			MovimientosRespuesta, _ := helpersCrud.CrearMovimientosDetallePlan(planAdquisiconesId, planAdquisiconesIdMongo)
			if len(MovimientosRespuesta) != 0 {
				alertErr.Type = "OK"
				alertErr.Code = "200"
				alertErr.Body = map[string]interface{}{"PlanAdquisiciones": PlanAdquisicionRespuesta, "Movimientos": MovimientosRespuesta}
			}
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

// GetOne Función para obterner la version del plan de adquision almacenadas en mongo
// @Title GetOne
// @Description get Plan_adquisicionController by id
// @Param	id		path 	string	true		"Id de un  plan de adquisicion"
// @Success 200 {object} models.Plan_adquisicion
// @Failure 403 :id is empty
// @router /versiones/:id [get]
func (c *Plan_adquisicionController) GetOne() {

	planAdquisicionID := c.Ctx.Input.Param(":id")
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	RegistroPlanAdquisicion, errRegistroPlanAdquisicion := models.ObtenerVersionesMongoByID(planAdquisicionID)

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
