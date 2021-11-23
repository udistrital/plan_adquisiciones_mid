package models

import "time"

type RegistroPlanAdquisicionRecibida struct {
	AreaFuncional                    int
	CentroGestor                     int
	ResponsableId                    int
	Activo                           bool
	RubroId                          string
	FuenteFinanciamientoId           string
	FechaEstimadaInicio              time.Time
	FechaEstimadaFin                 time.Time
	PlanAdquisicionesId              int
	ModalidadSeleccion               []ModalidadSeleccion
	CodigoArka                       []CodigoArka
	RegistroPlanAdquisicionActividad []RegistroPlanAdquisicionActividad
	MetasAsociadas                   []MetaAsociada
	ProductosAsociados               []ProductoAsociado
	Type                             string `json:"type"`
}

type ModalidadSeleccion struct {
	Id                          int
	IdModalidadSeleccion        string
	Activo                      bool
	RegistroPlanAdquisicionesId interface{}
}

type CodigoArka struct {
	Id         int
	CodigoArka string
	Activo     bool
}

type RegistroPlanAdquisicionActividad struct {
	Id                    int
	RegistroActividadId   int
	ActividadId           int
	Valor                 int
	Activo                bool
	FuentesFinanciamiento []FuenteFinanciamiento
}

type FuenteFinanciamiento struct {
	Id                     int
	FuenteFinanciamientoId string
	Activo                 bool
	ValorAsignado          int
}

type MetaAsociada struct {
	Id     int
	Activo bool
	MetaId int
}

type ProductoAsociado struct {
	Id                     int
	Activo                 bool
	ProductoAsociadoId     string
	PorcentajeDistribucion int
}
