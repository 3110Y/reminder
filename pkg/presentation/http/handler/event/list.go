package event

import (
	"encoding/json"
	"github.com/3110Y/reminder/pkg/application/service"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type List struct {
	Event *service.EventService
}

func (l List) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	eventIDString := chi.URLParam(request, "id")
	eventID, err := strconv.ParseUint(eventIDString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	LimitString := chi.URLParam(request, "limit")
	limit, err := strconv.Atoi(LimitString)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	eventList, err := l.Event.List(uint(eventID), limit)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	jsonEvent, err := json.Marshal(eventList)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	_, err = response.Write(jsonEvent)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
