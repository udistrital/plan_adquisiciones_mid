package models_test

import (
	"testing"

	"github.com/udistrital/plan_adquisiciones_mid/models"
)

func TestObtenerRegistroPlanAdquisicion(t *testing.T){
	valor, err := models.ObtenerRegistroPlanAdquisicion()
	if err != nil{
		t.Error("No se pudo obtener el registro del plan de adquision")
		t.Fail()
	}else{
		t.Log(valor)
		t.Log("TestObtenerRegistroPlanAdquisicion Finalizado Correctamente (OK)")
	}
}