package event

import (
	"github.com/3110Y/reminder/pkg/application/service"
	"net/http"
)

type Delete struct {
	Event *service.EventService
}

func (d Delete) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
