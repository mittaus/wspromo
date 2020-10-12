package main

import (
	"log"

	application "dev.azure.com/spsa/wspromo/application"
	config "dev.azure.com/spsa/wspromo/config"
	server_routes "dev.azure.com/spsa/wspromo/infraestructure/gin.server"
	v1 "dev.azure.com/spsa/wspromo/infraestructure/gin.server/v1"
	repository "dev.azure.com/spsa/wspromo/infraestructure/repository"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
)

func main() {
	var dbs []*gorm.DB
	//inicializar base de datos
	configDb := &config.GORMConfig{}
	err := configor.Load(configDb, "config.yml")
	if err != nil {
		panic("No se encontró archivo de configuración")
	}
	dbs, _ = repository.NewDb(configDb)
	repository := repository.NewDatacenterRepository(dbs)
	app := application.NewDatacenter(repository)
	datacenterHandler := v1.NewDatacenterHandler(*app)
	handler := server_routes.NewRouterHandler(*datacenterHandler)

	// serverPort, _ := strconv.Atoi(os.Getenv("server.port"))
	ginServer := config.NewServer(
		8080,
		config.DebugMode,
	)

	handler.SetRoutes(ginServer.Router)

	ginServer.Start()

	log.Println("server has started")
}
