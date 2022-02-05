package event

import (
	"encoding/json"
	"github.com/3110Y/reminder/pkg/application/service"
	"net/http"
)

type Add struct {
	Event *service.EventService
}

func (a *Add) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	event := service.Event{}
	err := json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}
	_, err = a.Event.Add(event)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
