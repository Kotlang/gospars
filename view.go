package gospars

type ViewController interface {
	Handle(templateBody TemplateBody, params map[string]string)
	GetTemplatePath() string
}
