package models

import (
	"fmt"
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

//IngresoRegistroPlanInversionActividadFuente ...
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

//ActualizarRegistroActividadFuente ...
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

//ObtenerRegistroTablaActividades ...
func ObtenerRegistroTablaActividades(idStr string) (registroPlanAdquisicionActividadFuente []map[string]interface{}, outputError interface{}) {
	var RegistroPlanAdquisicionActividadFuente []map[string]interface{}
	var unicos []string
	registro := make(map[string]interface{})
	registros := make([]map[string]interface{}, 0)
	fuentesFinanciamiento := make([]map[string]interface{}, 0)
	query := "?query=RegistroPlanAdquisicionesActividadId.RegistroPlanAdquisicionesId.Id%3A" + idStr + "%2CRegistroPlanAdquisicionesActividadId.Activo%3Atrue%2CActivo%3Atrue&sortby=RegistroPlanAdquisicionesActividadId__Id&order=asc"
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/"+query, &RegistroPlanAdquisicionActividadFuente)

	if error != nil {
		return nil, error
	} else {
		if len(RegistroPlanAdquisicionActividadFuente) == 1 {
			if len(RegistroPlanAdquisicionActividadFuente[0]) == 0 {
				error := "No existe plan adquisicion"
				return nil, error
			}
		}
		for index := range RegistroPlanAdquisicionActividadFuente {
			ActividadID := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["ActividadId"].(map[string]interface{})["Id"]
			ValorActividad := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["Valor"]
			RegistroActividadID := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["Id"]
			ActivoActividad := RegistroPlanAdquisicionActividadFuente[index]["RegistroPlanAdquisicionesActividadId"].(map[string]interface{})["Activo"]
			newdata := stringInSlice(fmt.Sprintf("%.0f", RegistroActividadID.(float64)), unicos)
			if !newdata {
				unicos = append(unicos, fmt.Sprintf("%.0f", RegistroActividadID.(float64)))
				registros = append(registros, registro)
				fuentesFinanciamiento = make([]map[string]interface{}, 0)
			}

			fuenteFinanciamiento := map[string]interface{}{
				"Id":                   RegistroPlanAdquisicionActividadFuente[index]["Id"],
				"ValorAsignado":        RegistroPlanAdquisicionActividadFuente[index]["ValorAsignado"],
				"Activo":               RegistroPlanAdquisicionActividadFuente[index]["Activo"],
				"FuenteFinanciamiento": RegistroPlanAdquisicionActividadFuente[index]["FuenteFinanciamientoId"],
			}
			fuentesFinanciamiento = append(fuentesFinanciamiento, fuenteFinanciamiento)

			registro = map[string]interface{}{
				"ActividadId":                 ActividadID,
				"RegistroPlanAdquisicionesId": idStr,
				"Valor":                       ValorActividad,
				"Activo":                      ActivoActividad,
				"RegistroActividadId":         RegistroActividadID,
				"FuentesFinanciamiento":       fuentesFinanciamiento,
			}

		}
		// valor := SumaFuenteFinanciamiento(idStr)
		// fmt.Println(valor)
		registros = append(registros[1:], registro)
		return registros, nil
	}

}

//RegistroFuenteModificado ...
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

//SumaFuenteFinanciamiento ...
func SumaFuenteFinanciamiento(idStr string) (total interface{}) {
	var RegistroPlanAdquisicionActividadFuente []map[string]interface{}
	var valor float64
	query := "?query=RegistroPlanAdquisicionesActividadId.RegistroPlanAdquisicionesId.Id%3A" + idStr + "&sortby=RegistroPlanAdquisicionesActividadId__Id&order=asc"
	error := request.GetJson(beego.AppConfig.String("plan_adquicisiones_crud_url")+"Registro_inversion_actividad-Fuente_financiamiento/"+query, &RegistroPlanAdquisicionActividadFuente)

	if error != nil {
		return error
	} else {
		for index := range RegistroPlanAdquisicionActividadFuente {
			valor = valor + RegistroPlanAdquisicionActividadFuente[index]["ValorAsignado"].(float64)
		}
		return valor
	}
}
