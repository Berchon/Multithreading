package route

import (
	"fmt"

	"github.com/Berchon/Multithreading/internal/business/usecase"
	"github.com/Berchon/Multithreading/internal/infra/service"
	"github.com/Berchon/Multithreading/internal/infra/webapp/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func ConfigureApplicationRoutes(port string) *chi.Mux {
	router := chi.NewRouter()
	registerRoutes(port, router)
	return router
}

func registerRoutes(port string, router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	getAddressByCepService := service.NewAddressService()
	getAddressByCepUsecase := usecase.NewGetAddressByCepUsecase(getAddressByCepService)
	getAddressByCep := handler.NewGetAddressByCepHandler(getAddressByCepUsecase)
	router.Get("/{cep}", getAddressByCep.Handle)

	urlSwagger := fmt.Sprintf("http://localhost:%s/docs/doc.json", port)
	swaggerHandlerFunc := httpSwagger.Handler(httpSwagger.URL(urlSwagger))
	router.Get("/docs/*", swaggerHandlerFunc)
}
