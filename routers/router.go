// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/plan_adquisiciones_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/RegistrosPlanAdquisicion",
			beego.NSInclude(
				&controllers.Registro_plan_adquisicionController{},
			),
		),
		beego.NSNamespace("/RegistrosOrdenadoPorRubro",
			beego.NSInclude(
				&controllers.Plan_adquisicion_por_fuentesController{},
			),
		),
		beego.NSNamespace("/RegistrosPlanAdquisicionActividad",
			beego.NSInclude(
				&controllers.Registro_PlanAdquisiciones_ActividadController{},
			),
		),
		beego.NSNamespace("/RegistrosPlanInversionActividadFuente",
			beego.NSInclude(
				&controllers.Registro_PlanInversion_ActividadFuente_financiamientoController{},
			),
		),
		beego.NSNamespace("/PlanAdquisicion",
			beego.NSInclude(
				&controllers.Plan_adquisicionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
