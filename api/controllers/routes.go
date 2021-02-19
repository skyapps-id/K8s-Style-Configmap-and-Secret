package controllers

import "github.com/skyapps-id/K8s-Style-Configmap-and-Secret/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

}
