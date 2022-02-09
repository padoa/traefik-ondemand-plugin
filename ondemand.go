package traefik_ondemand_plugin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/acouvreur/traefik-ondemand-plugin/pkg/service"
	"github.com/acouvreur/traefik-ondemand-plugin/pkg/strategy"
)

// Config the plugin configuration
type Config struct {
	Name        string   `yaml:"name"`
	Names       []string `yaml:"names"`
	ServiceUrl  string   `yaml:"serviceurl"`
	Timeout     string   `yaml:"timeout"`
	ErrorPage   string   `yaml:"errorpage"`
	LoadingPage string   `yaml:"loadingpage"`
	WaitUi      bool     `yaml:"waitui"`
	BlockDelay  string   `yaml:"blockdelay"`
}

// CreateConfig creates a config with its default values
func CreateConfig() *Config {
	return &Config{
		Timeout:     "1m",
		WaitUi:      true,
		BlockDelay:  "1m",
		ErrorPage:   "",
		LoadingPage: "",
	}
}

// Ondemand holds the request for the on demand service
type Ondemand struct {
	strategy strategy.Strategy
}

func buildServiceInfo(url string, name string, timeout time.Duration) (service.Service, error) {
	request := fmt.Sprintf("%s?name=%s&timeout=%s", url, name, timeout.String())
	return service.Service{
		Name:    name,
		Started: false,
		Request: request,
	}, nil
}

// New function creates the configuration
func New(ctx context.Context, next http.Handler, config *Config, middlewareName string) (http.Handler, error) {
	if len(config.ServiceUrl) == 0 {
		return nil, fmt.Errorf("serviceurl cannot be null")
	}

	if len(config.Name) != 0 && len(config.Names) != 0 {
		return nil, fmt.Errorf("both name and names cannot be used simultaneously")
	}
	var serviceNames []string

	if len(config.Name) != 0 {
		serviceNames = append(serviceNames, config.Name)
	} else if len(config.Names) != 0 {
		serviceNames = config.Names
	} else {
		return nil, fmt.Errorf("both name and names cannot be null")
	}

	timeout, err := time.ParseDuration(config.Timeout)

	if err != nil {
		return nil, err
	}
	var services []service.Service

	for _, serviceName := range serviceNames {
		service, err := buildServiceInfo(config.ServiceUrl, serviceName, timeout)
		if err != nil {
			return nil, fmt.Errorf("error while building request for %s", serviceName)
		}
		services = append(services, service)
	}
	strategy, err := config.getServeStrategy(services, middlewareName, next, timeout)

	if err != nil {
		return nil, err
	}

	return &Ondemand{
		strategy: strategy,
	}, nil
}

func (config *Config) getServeStrategy(services []service.Service, name string, next http.Handler, timeout time.Duration) (strategy.Strategy, error) {
	if config.WaitUi {
		return &strategy.DynamicStrategy{
			Services:    services,
			Name:        name,
			Next:        next,
			Timeout:     timeout,
			ErrorPage:   config.ErrorPage,
			LoadingPage: config.LoadingPage,
		}, nil
	} else {

		blockDelay, err := time.ParseDuration(config.BlockDelay)

		if err != nil {
			return nil, err
		}

		return &strategy.BlockingStrategy{
			Services:           services,
			Name:               name,
			Next:               next,
			Timeout:            timeout,
			BlockDelay:         blockDelay,
			BlockCheckInterval: 1 * time.Second,
		}, nil
	}
}

// ServeHTTP retrieve the service status
func (e *Ondemand) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	e.strategy.ServeHTTP(rw, req)
}
