package event

import (
	"github.com/3110Y/reminder/pkg/application/service"
	"net/http"
)

type Edit struct {
	Event *service.EventService
}

func (e Edit) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
