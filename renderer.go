package gospars

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"html/template"
)

type TemplateBody struct {
	templateHtml string
}

func (t TemplateBody) Render(data interface{}) string {
	if data == nil {
		return t.templateHtml
	}
	tmpl := template.Must(template.New("Search").Parse(t.templateHtml))
	var resultHtml bytes.Buffer
	tmpl.Execute(&resultHtml, data)
	return resultHtml.String()
}

func getTemplate(path string, callback func(error, TemplateBody)) {
	go func() {
		resp, err := http.Get(path)
		if err != nil {
			callback(err, TemplateBody{""})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		callback(nil, TemplateBody{string(body[:])})
	}()
}
