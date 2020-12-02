package models

import (
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

//ObtenerRegistroPlanInversionActividadFuente ...
func ObtenerRegistroPlanInversionActividadFuente() (registroPlanAdquisicionActividadFuente []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividadFuente []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/", &RegistroPlanAdquisicionActividadFuente)

	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividadFuente, nil
	}

}

//IngresoRegistroActividad ...
func IngresoRegistroPlanInversionActividadFuente(registroActividadFuente map[string]interface{}) (registroActividadFuenteRespuesta map[string]interface{}, outputError interface{}) {
	registroActividadFuenteIngresado := make(map[string]interface{})
	registroActividadFuentePost := make(map[string]interface{})

	registroActividadFuenteIngresado = map[string]interface{}{
		"RegistroPlanAdquisicionesActividadId": map[string]interface{}{"Id": registroActividadFuente["ActividadId"]},
		"ValorAsignado":                        registroActividadFuente["ValorAsignado"],
		"FuenteFinanciamientoId":               registroActividadFuente["FuenteFinanciamientoId"],
		"Activo":                               registroActividadFuente["Activo"],
	}
	error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/", "POST", &registroActividadFuentePost, registroActividadFuenteIngresado)
	if error != nil {
		return nil, error
	} else {
		return registroActividadFuentePost, nil
	}

}

//ActualizarRegistroActividad ...
func ActualizarRegistroActividadFuente(registroActividadFuente map[string]interface{}, idStr string, idStrActividad string) (registroActividadFuenteRespuesta map[string]interface{}, outputError interface{}) {
	fuenteActividadPut := make(map[string]interface{})
	fuenteActividadActualizar := make(map[string]interface{})
	RegistroPlanAdquisicionActividadFuente, error := ObtenerRegistroPlanAdquisicionActividadFuenteByID(idStr)
	if error != nil {
		return nil, error
	} else {
		if RegistroPlanAdquisicionActividadFuente["Status"] != nil {
			//fmt.Println("No existe fuente toca crearla")
			idint, _ := strconv.Atoi(idStrActividad)
			registroActividadFuente["ActividadId"] = idint
			RegistroActividadFuente, errRegistroActividadFuente := IngresoRegistroPlanInversionActividadFuente(registroActividadFuente)
			if errRegistroActividadFuente != nil {
				return nil, errRegistroActividadFuente
			} else {
				return RegistroActividadFuente, nil
			}

		} else {
			validacion := RegistroFuenteModificado(registroActividadFuente, RegistroPlanAdquisicionActividadFuente)
			if validacion {
				//fmt.Println("existe fuente No toca modificar")
			} else {
				//fmt.Println("existe fuente toca modificar")
				idint, _ := strconv.Atoi(idStrActividad)
				registroActividadFuente["ActividadId"] = idint
				fuenteActividadActualizar = map[string]interface{}{
					"RegistroPlanAdquisicionesActividadId": map[string]interface{}{"Id": registroActividadFuente["ActividadId"]},
					"ValorAsignado":                        registroActividadFuente["ValorAsignado"],
					"FechaCreacion":                        RegistroPlanAdquisicionActividadFuente["FechaCreacion"],
					"FuenteFinanciamientoId":               registroActividadFuente["FuenteFinanciamientoId"],
					"Activo":                               registroActividadFuente["Activo"],
				}
				error := request.SendJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/"+idStr, "PUT", &fuenteActividadPut, fuenteActividadActualizar)
				if error != nil {
					return nil, error
				} else {
					return fuenteActividadPut, nil
				}
			}
		}

		return RegistroPlanAdquisicionActividadFuente, nil
	}

}

//ObtenerRegistroPlanAdquisicionActividadFuenteByID ...
func ObtenerRegistroPlanAdquisicionActividadFuenteByID(idStr string) (registroPlanAdquisicionActividadFuente map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividadFuente map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/"+idStr, &RegistroPlanAdquisicionActividadFuente)
	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividadFuente, nil
	}

}

//ObtenerRegistroPlanAdquisicionActividadFuenteByActividadID ...
func ObtenerRegistroPlanAdquisicionActividadFuenteByActividadID(idStr string) (registroPlanAdquisicionActividadFuente []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividadFuente []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/?query=RegistroPlanAdquisicionesActividadId.Id%3A"+idStr, &RegistroPlanAdquisicionActividadFuente)
	if error != nil {
		return nil, error
	} else {
		return RegistroPlanAdquisicionActividadFuente, nil
	}

}

func RegistroFuenteModificado(registroFuente map[string]interface{}, RegistroPlanAdquisicionActividadFuente map[string]interface{}) (validacion bool) {
	registroFuenteActual := make(map[string]interface{})

	registroFuenteActual = map[string]interface{}{
		"ValorAsignado":          RegistroPlanAdquisicionActividadFuente["ValorAsignado"],
		"FechaCreacion":          RegistroPlanAdquisicionActividadFuente["FechaCreacion"],
		"FuenteFinanciamientoId": RegistroPlanAdquisicionActividadFuente["FuenteFinanciamientoId"],
		"Activo":                 RegistroPlanAdquisicionActividadFuente["Activo"],
	}

	if reflect.DeepEqual(registroFuenteActual["ValorAsignado"], registroFuente["ValorAsignado"]) && reflect.DeepEqual(registroFuenteActual["Activo"], registroFuente["Activo"]) && reflect.DeepEqual(registroFuenteActual["FuenteFinanciamientoId"], registroFuente["FuenteFinanciamientoId"]) {
		return true
	} else {
		return false
	}

}
