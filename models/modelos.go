package models

import "time"

type RegistroPlanAdquisiciones struct {
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

type RegistrosMultiplesMovimientos struct {
	Activo                     bool
	Descripcion                string
	Detalle                    string
	FechaCreacion              time.Time
	FechaModificacion          time.Time
	Id                         int
	MovimientoProcesoExternoId MovimientoProcesoExternoId
	Saldo                      int
	Valor                      int
}

type MovimientoProcesoExternoId struct {
	Activo                   bool
	Detalle                  string
	FechaCreacion            time.Time
	FechaModificacion        time.Time
	Id                       int
	MovimientoProcesoExterno int
	ProcesoExterno           int
	TipoMovimientoId         TipoMovimientoId
}

type TipoMovimientoId struct {
	Acronimo          string
	Activo            bool
	Descripcion       string
	FechaCreacion     time.Time
	FechaModificacion time.Time
	Id                int
	Nombre            string
	Parametros        string
}

type MovimientoProcesoExterno struct {
	TipoMovimientoId         int
	ProcesoExterno           int
	MovimientoProcesoExterno int
	Activo                   bool
	Detalle                  string
}

type DetalleMovimientoProcesoExterno struct {
	Estado              string
	PlanAdquisicionesId int
}

type MovimientosDetalle struct {
	MovimientoProcesoExternoId int
	Valor                      float64
	Descripcion                string
	Activo                     bool
	Saldo                      float64
	Detalle                    string
}
