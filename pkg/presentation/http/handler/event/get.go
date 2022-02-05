package event

import (
	"encoding/json"
	"github.com/3110Y/reminder/pkg/application/service"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type Get struct {
	Event *service.EventService
}

func (g Get) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	eventIDString := chi.URLParam(request, "id")
	eventID, err := strconv.ParseUint(eventIDString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	event, err := g.Event.Get(uint(eventID))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	jsonEvent, err := json.Marshal(event)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	_, err = response.Write(jsonEvent)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
