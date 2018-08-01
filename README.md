## Minimalistic Single Page Application for GOPHERJS

* Provides routing
* Provides templating

## Installation

```
go get -u github.com/gopherjs/gopherjs
go get -u github.com/Kotlang/gospars/gospars
```

## Usage

### app.go
```
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/Kotlang/gospars/gospars"
	"landing"
	"profile"
)


func ConfigRouter(component *js.Object) {
	router := gospars.NewRouter(func(err error) {
		if err.Error() == "NO_ROUTE_FOUND" {
			component.Set("innerHTML", "404 - No path")
		} else {
			component.Set("innerHTML", "Check Internet")
		}
	})
	router.On("/landing", landing.LandingContoller{component})
	// Both path params and query params are supported
	// All path params and query params will be included in params called to handler of controller
	router.On("/profile/:user", profile.ProfileContoller{component})
	router.On("/search", search.SearchController{component})
	router.Init("/landing")
}

func main() {
	component := js.Global.Get("document").Call("getElementById", "root")
	ConfigRouter(component)
}

```

### index.html
```
<html>
<head>
    <title>Sample gospars App</title>
</head>

<body>

    <div class="container">
        <div id="root"></div>
    </div>
    <script src="build/app.js"></script>
</body>
</html>
```

### SearchController

```
package search

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/Kotlang/gospars/gospars"
)

type SearchController struct {
	Component *js.Object
}

type ExpertProfile struct {
	Name      string
	Expertise string
}

func(l SearchController) Handle(templateBody gospars.TemplateBody, params map[string]string) {
	experts := []ExpertProfile {
		{Name: "Sai NS", Expertise: "Machine Learning"},
		{Name: "Vasu", Expertise: "Machine Learning"},
		{Name: "Siddhanth", Expertise: "Machine Learning"} }


	l.Component.Set("innerHTML", templateBody.Render(experts))
}

func (l SearchController) GetTemplatePath() string  {
	return "build/search/search.html"
}

```

### search.html (template)

```
<!-- templating as provided by goland template/html -->
<h1>Results</h1>
<ul>
    {{range .}}
        <li>{{.Name}}</li>
        <li>{{.Expertise}}</li>
    {{end}}
</ul>
```