package main

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"unicode"
	"text/template"
)

var currentDirectory string

func init() {
	var err error
	currentDirectory, err = os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func main()  {
	createComponent := &cobra.Command{
		Use: "component [component name] [output path]",
		Short: "Generates controller and template for a route",
		Long: `Generates controller and template for a route.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			data := struct {
				Component string
				Controller string
				HtmlFile string
			}{
				args[0],
				ucFirst(args[0]) + "Controller",
				args[0] + ".html" }

			dest := "."
			if len(args) == 2 {
				dest = args[1]
			}

			componentFolder := dest + "/" + data.Component
			htmlFile := componentFolder + "/" + data.HtmlFile
			controllerFile := componentFolder + "/" + data.Controller + ".go"

			os.Mkdir(componentFolder, 0777)
			os.Create(htmlFile)
			controllerFilePtr, _ := os.Create(controllerFile)

			tmpl := template.Must(template.New("Controller").Parse(controllerTemplate))
			tmpl.Execute(controllerFilePtr, data)
			controllerFilePtr.Close()
		}}

	createCommand := &cobra.Command{Use: "create"}
	createCommand.AddCommand(createComponent)

	rootCmd := &cobra.Command{Use: "gospars"}
	rootCmd.AddCommand(createCommand)
	rootCmd.Execute()
}

const controllerTemplate = `
package {{.Component}}

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/Kotlang/gospars/gospars"
)

type {{.Controller}} struct {
	Component *js.Object
}

func(l {{.Controller}})Handle(templateBody gospars.TemplateBody, params map[string]string) {
	l.Component.Set("innerHTML", templateBody.Render(nil))
}

func (l {{.Controller}}) GetTemplatePath() string  {
	return "build/{{.Component}}/{{.HtmlFile}}"
}

`