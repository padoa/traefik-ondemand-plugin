package strategy

import (
	"log"
	"net/http"
	"time"

	"github.com/acouvreur/traefik-ondemand-plugin/pkg/pages"
	"github.com/acouvreur/traefik-ondemand-plugin/pkg/service"
)

type DynamicStrategy struct {
	Services    []service.Service
	Name        string
	Next        http.Handler
	Timeout     time.Duration
	LoadingPage string
	ErrorPage   string
}

// ServeHTTP retrieve the service status
func (e *DynamicStrategy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var allServicesReady bool
	for _, service := range e.Services {
		allServicesReady = true
		log.Printf("Sending request: %s", service.Request)
		status, err := getServiceStatus(service.Request)
		log.Printf("Status: %s", status)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(pages.GetErrorPage(e.ErrorPage, e.Name, err.Error())))
		}

		if status == "started" {
			service.Started = true
		} else if status == "starting" {
			service.Started = false
			allServicesReady = false
		} else {
			// Error
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(pages.GetErrorPage(e.ErrorPage, e.Name, status)))
		}
	}
	if allServicesReady {
		// All services are ready, forward request
		e.Next.ServeHTTP(rw, req)
	} else {
		// Services still starting, notify client
		rw.WriteHeader(http.StatusAccepted)
		rw.Write([]byte(pages.GetLoadingPage(e.LoadingPage, e.Name, e.Services, e.Timeout)))
	}
}
