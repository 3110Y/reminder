package event

import (
	"github.com/3110Y/reminder/pkg/application/service"
	"net/http"
)

type List struct {
	Event *service.EventService
}

func (l List) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
