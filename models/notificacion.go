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

type ProduccionAcademica struct{
	Fecha                   string
	NombreDocente           string
	llaveConsulta           string
	ContenidoProduccion     template.HTML
	UrlCreacionCuentaLogin  template.HTML
	UrlRechazarEvaluacion   template.HTML
}