package http

import (
	"net/http"

	"github.com/jadahbakar/backend-golang/internal/master/entities"
)

type handler struct {
	masterService entities.MasterService
}

type RedirectHandler interface {
	Get(http.ResponseWriter)
	Post(http.ResponseWriter)
}

func NewHandler(masterService entities.MasterService) RedirectHandler {
	return &handler{masterService: masterService}
}

func (h *handler) Get(http.ResponseWriter) {

}
func (h *handler) Post(http.ResponseWriter) {

}
