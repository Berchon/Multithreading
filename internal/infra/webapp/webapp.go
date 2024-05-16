package webapp

import (
	"fmt"
	"net/http"

	"github.com/Berchon/Multithreading/configs"
	"github.com/Berchon/Multithreading/internal/infra/webapp/route"
)

type WebApp interface {
	Start()
}

type webApp struct{}

func New() WebApp {
	return &webApp{}
}

func (webApp *webApp) Start() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	router := route.ConfigureApplicationRoutes(configs.WebServerPort)

	port := fmt.Sprintf(":%s", configs.WebServerPort)
	http.ListenAndServe(port, router)
}
