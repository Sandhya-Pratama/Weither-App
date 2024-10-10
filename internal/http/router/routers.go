package router

import "github.com/labstack/echo/v4"

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func PublicRoutes() []*Route  {	
	return []*Route{}
}

func PrivateRoutes() []*Route  {	
	return []*Route{}
}