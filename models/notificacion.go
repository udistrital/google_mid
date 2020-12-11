package models

type Notificacion struct {
	Trom         string
	To           []string
	CC           []string
	BCC          []string
	Subject      string
	TemplateName string
	TemplateData interface{}
}