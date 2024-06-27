package htte

type RouterConfig struct {
}

type Router struct {
	config RouterConfig
	routes []Route
}

type Route struct {
	Method  string
	Path    string
	Handler func(Request) string
}

type HandlerFunc func(Request) string

func (router *Router) addRoute(method string, path string, handler HandlerFunc) {
	var route = Route{Method: method, Path: path, Handler: handler}
	router.routes = append(router.routes, route)
}

func (app *App) Get(path string, handler HandlerFunc) {
	app.Router.addRoute("GET", path, handler)
}

func (app *App) Post(path string, handler HandlerFunc) {
	app.Router.addRoute("Post", path, handler)
}

func (app *App) Put(path string, handler HandlerFunc) {
	app.Router.addRoute("PUT", path, handler)
}

func (app *App) Patch(path string, handler HandlerFunc) {
	app.Router.addRoute("PATCH", path, handler)
}

func (app *App) Delete(path string, handler HandlerFunc) {
	app.Router.addRoute("DELETE", path, handler)
}

func (router *Router) match(path string, method string) HandlerFunc {
	for _, v := range router.routes {
		if path == v.Path {
			if method == v.Method {
				return v.Handler
			}
		}
	}

	return nil
}
