package event

import (
	"github.com/3110Y/reminder/pkg/application/service"
	"net/http"
)

type Get struct {
	Event *service.EventService
}

func (g Get) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
