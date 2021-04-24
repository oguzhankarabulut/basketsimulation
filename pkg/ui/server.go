package ui

import (
	"basketsimulation/pkg/infrastructure/mongo"
	"basketsimulation/pkg/service"
	"net/http"
)


type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Run(mr mongo.MatchRepository , mpr mongo.MatchPlayerRepository) error {
	rs := service.NewResultService(mr, mpr)
	rh := NewResultHandler(rs)

	http.HandleFunc("/result", rh.HandleResult)
	http.HandleFunc("/dashboard", rh.HandleResult)
	return http.ListenAndServe("localhost:8000", nil)
}
