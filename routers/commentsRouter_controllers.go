package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Plan_adquisicion_por_fuentesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_adquisiciones_mid/controllers:Plan_adquisicion_por_fuentesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/:planAdquisicionID",
            AllowHTTPMethods: []string{"get"},
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

}
