package controllers

import "github.com/Ilhamsat/go/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/api/test", middlewares.SetMiddlewareJSON(s.Tsv)).Methods("GET")

}
