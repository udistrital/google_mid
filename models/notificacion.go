package models

import (
	"html/template"
)

type Notificacion struct {
	Trom         string
	To           []string
	CC           []string
	BCC          []string
	Subject      string
	TemplateName string
	TemplateData *ProduccionAcademica
}

type ProduccionAcademica struct {
	Fecha                  string
	NombreDocente          string
	NombreEstudiante       string
	llaveConsulta          string
	EspacioAcademico       string
	FechaInicio            string
	FechaFin               string
	ContenidoProduccion    template.HTML
	UrlCreacionCuentaLogin template.HTML
	UrlRechazarEvaluacion  template.HTML
}

// RespNotificacionPost es el tipo de dato que retorna
// POST notificacion/
//
// TODO: Completar para que se actualice el swagger
type RespNotificacionPost struct {
}
