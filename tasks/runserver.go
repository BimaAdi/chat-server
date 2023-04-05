package tasks

import (
	"github.com/BimaAdi/chat-server/docs"
	"github.com/BimaAdi/chat-server/models"
	"github.com/BimaAdi/chat-server/routes"
	"github.com/BimaAdi/chat-server/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer(envPath string) {
	// Initialize environtment variable
	settings.InitiateSettings(envPath)

	// Initiate Database connection
	models.Initiate()

	// Cors Middleware
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOrigins:           []string{},
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowHeaders:           []string{"Origin", "Content-Type", "authorization", "accept"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Content-Length"},
		MaxAge:                 0,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	// Initiate static and template
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	// Initialize gin route
	routes := routes.GetRoutes(router)

	// setup swagger
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = settings.SERVER_HOST + ":" + settings.SERVER_PORT
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// run gin server
	routes.Run(settings.SERVER_HOST + ":" + settings.SERVER_PORT)
}
