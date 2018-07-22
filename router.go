package gospars

import (
	"github.com/gopherjs/gopherjs/js"
	"regexp"
	"strings"
	"errors"
)

type action struct {
	path string
	handler ViewController
}

type Router struct {
	// document.location reference
	dloc *js.Object
	actions []action
	errCallback func(err error)
}

func matchAllPathParams(configPath string, locationPath string) (bool, map[string]string) {
	configPathParams := strings.Split(configPath, "/")
	locationPathParams := strings.Split(locationPath, "/")

	if len(configPathParams) != len(locationPathParams) {
		return false, map[string]string {}
	}

	pathParamsMap := map[string]string {}
	for i, configPathParam := range configPathParams {
		if strings.HasPrefix(configPathParam, ":") {
			pathParamsMap[configPathParam] = locationPathParams[i]
			continue
		} else if configPathParam != locationPathParams[i] {
			return false, map[string]string {}
		}
	}
	return true, pathParamsMap
}

func (r *Router) onHashChangeEvent(changeEvent *js.Object)  {
	hash := r.dloc.Get("hash").String()
	r.fireEvent(hash)
}

/**
 * Returns appropriate ViewController and a map of path params
 * For e.g. /profile/:user will match with /profile/sai.satchidanand
 * and ProfileController, map[string]string { ":user": "sai.satchidanand" } will be returned
 * Returns nil if no match found
 */
func (r *Router) getHandler(path string) (ViewController, map[string]string) {
	for _, actionHandle := range r.actions {
		if isMatch, pathParamsMap := matchAllPathParams(actionHandle.path, path); isMatch {
			return actionHandle.handler, pathParamsMap
		}
	}
	return nil, nil
}

func (r *Router) fireEvent(path string)  {
	// remove #/ from path
	cleanPath, _ := regexp.Compile("#")
	path = cleanPath.ReplaceAllString(path, "")

	v, params := r.getHandler(path)
	if v == nil {
		r.errCallback(errors.New("NO_ROUTE_FOUND"))
		return
	}

	getTemplate(v.GetTemplatePath(), func(err error, templateContent string) {
		if err != nil {
			r.errCallback(err)
			return
		}
		v.Handle(templateContent, params)
	})
}

func NewRouter(errCallback func(error)) *Router {
	return &Router{ dloc: js.Global.Get("document").Get("location"),
					actions: []action{},
					errCallback: errCallback }
}

func (r *Router) On(path string, v ViewController) {
	r.actions = append(r.actions, action{ path: path, handler:v })
}

func (r *Router) Init(path string) {
	js.Global.Get("window").Set("onhashchange", r.onHashChangeEvent)
	currentHash := r.dloc.Get("hash").String()

	if currentHash == "" || currentHash =="#" {
		r.dloc.Set("hash", path)
	} else {
		r.fireEvent(currentHash)
	}
}
