package ui

import (
	"basketsimulation/pkg/service"
	"net/http"
)

const (
	GET = "GET"
)

type ResultHandler struct {
	resultService service.ResultService
}

func NewResultHandler(rs service.ResultService) *ResultHandler {
	return &ResultHandler{resultService: rs}
}

func (h ResultHandler) HandleResult(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GET:
		switch r.URL.Path {
		case "/result":
			h.Result(w)
		case "/dashboard":
			h.Dashboard(w, r)
		default:
			WriteError(w, nil, http.StatusBadRequest)
		}

	default:
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}
}

func (h ResultHandler) Result (w http.ResponseWriter)  {
	r := h.resultService.Result()
	write(w, r)
}

func (h ResultHandler) Dashboard (w http.ResponseWriter, r *http.Request) {
	template := template()
	_ = template.Execute(w, nil)
}