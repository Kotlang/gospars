## Minimalistic Single Page Application for GOPHERJS

* Provides routing
* Provides templating

## Installation

```
go get -u github.com/gopherjs/gopherjs
go get -u github.com/Kotlang/gospars
```

## Usage

### app.js
```
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/Kotlang/gospars"
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
	router.On("/profile/:user", profile.ProfileContoller{component})
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

### ProfileController

```
package profile

import (
	"github.com/gopherjs/gopherjs/js"
)

type ProfileContoller struct {
	Component *js.Object
}

// template string will contain complete html of "build/profile/profile.html"
func(p ProfileContoller) Handle(template string, params map[string]string) {
	p.Component.Set("innerHTML", "Welcome " + params[":user"])
}

func (p ProfileContoller) GetTemplatePath() string  {
	return "build/profile/profile.html"
}
```