package gospars

type ViewController interface {
	Handle(templateContent string, params map[string]string)
	GetTemplatePath() string
}
