package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Plan_adquisicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Plan_adquisicionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Plan_adquisicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Plan_adquisicionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Plan_adquisicion_por_fuentesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Plan_adquisicion_por_fuentesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/:planAdquisicionID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanAdquisiciones_ActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanAdquisiciones_ActividadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanAdquisiciones_ActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanAdquisiciones_ActividadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanAdquisiciones_ActividadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanAdquisiciones_ActividadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanInversion_ActividadFuente_financiamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanInversion_ActividadFuente_financiamientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanInversion_ActividadFuente_financiamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanInversion_ActividadFuente_financiamientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanInversion_ActividadFuente_financiamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_PlanInversion_ActividadFuente_financiamientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/:idPlanAdquisicion",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_plan_adquisicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_plan_adquisicionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_plan_adquisicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_plan_adquisicionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_plan_adquisicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_plan_adquisicionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_plan_adquisicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Registro_plan_adquisicionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
