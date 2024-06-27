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
