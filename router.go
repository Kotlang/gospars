package gospars

import (
	"github.com/gopherjs/gopherjs/js"
	"errors"
)

const NO_ROUTE_FOUND  = "NO_ROUTE_FOUND"

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

func (r *Router) onHashChangeEvent(changeEvent *js.Object)  {
	r.fireEvent()
}

/**
 * Returns appropriate ViewController and a map of path params and query params
 * For e.g. /profile/:user will match with /profile/sai.satchidanand
 * and ProfileController, map[string]string { ":user": "sai.satchidanand" } will be returned
 * Returns nil if no match found
 */
func (r *Router) getHandler(path string) (ViewController, map[string]string) {
	for _, actionHandle := range r.actions {
		if isMatch, pathParamsMap := MatchPathAndGetPathParams(actionHandle.path, path); isMatch {
			return actionHandle.handler, pathParamsMap
		}
	}
	return nil, nil
}

func (r *Router) fireEvent()  {
	pathParams := r.dloc.Get("hash").String()
	pathParams = GetHashPath(pathParams)

	v, params := r.getHandler(pathParams)
	if v == nil {
		r.errCallback(errors.New(NO_ROUTE_FOUND))
		return
	}
	queryParameterString := r.dloc.Get("search").String()
	queryParams := GetQueryParams(queryParameterString)
	params = MergeMaps(params, queryParams)

	getTemplate(v.GetTemplatePath(), func(err error, templateBody TemplateBody) {
		if err != nil {
			r.errCallback(err)
			return
		}
		v.Handle(templateBody, params)
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

func (r *Router) Init(landingHash string) {
	js.Global.Get("window").Set("onhashchange", r.onHashChangeEvent)
	currentHash := r.dloc.Get("hash").String()

	if currentHash == "" || currentHash =="#" {
		r.dloc.Set("hash", landingHash)
	} else {
		r.fireEvent()
	}
}
