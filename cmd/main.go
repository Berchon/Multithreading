package main

import (
	_ "github.com/Berchon/Multithreading/docs"
	"github.com/Berchon/Multithreading/internal/infra/webapp"
	_ "github.com/swaggo/http-swagger"
)

// @title Challenge Multithreading
// @version 1.0
// @description App receive PostCode and return address data from API BRASILAPI and VIACEP more faster.
// @termsOfService http://swagger.io/terms/

// @contact.name Meu Nome
// @contact.url http://www.meu-nome.com.br
// @contact.email meu-nome@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	app := webapp.New()
	app.Start()
}
