package httpserver

import "net/http"
type routes struct {
	path string
	method string
	handler http.HandlerFunc
}
type routeList struct {
	routeList []routes
}
func createRouteList()routeList  {
	r:=[]routes{}
	return routeList{routeList: r}
}

func (r *routeList)addRoutes(route routes) {
	r.routeList=append(r.routeList, route)
}

