package event

import (
	"encoding/json"
	"github.com/3110Y/reminder/pkg/application/service"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type Edit struct {
	Event *service.EventService
}

func (e Edit) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	eventIDString := chi.URLParam(request, "id")
	eventID, err := strconv.ParseUint(eventIDString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	event := service.Event{}
	err = json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}
	_, err = e.Event.Edit(event, uint(eventID))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
