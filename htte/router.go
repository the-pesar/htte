package htte

type Route struct {
	Method  string
	Path    string
	Handler func(Request) string
}

func (app *App) Get(path string, handler func(Request) string) {
	var route = Route{Method: "GET", Path: path, Handler: handler}
	app.routes = append(app.routes, route)
}

func (app *App) match(path string, method string) (Route, bool) {
	for _, v := range app.routes {
		if path == v.Path {
			if method == v.Method {
				return v, true
			}
		}
	}

	return Route{}, false
}
