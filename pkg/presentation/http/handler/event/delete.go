package event

import (
	"github.com/3110Y/reminder/pkg/application/service"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type Delete struct {
	Event *service.EventService
}

func (d *Delete) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	eventIDString := chi.URLParam(request, "id")
	eventID, err := strconv.ParseUint(eventIDString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	err = d.Event.Delete(uint(eventID))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
