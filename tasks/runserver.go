package tasks

import (
	"log"

	"github.com/BimaAdi/chat-server/docs"
	"github.com/BimaAdi/chat-server/models"
	"github.com/BimaAdi/chat-server/routes"
	"github.com/BimaAdi/chat-server/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
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

	// Initiate socket io
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	// Initialize gin route
	routes := routes.GetRoutes(router)

	// setup swagger
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = settings.SERVER_HOST + ":" + settings.SERVER_PORT
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// run gin server
	routes.Run(settings.SERVER_HOST + ":" + settings.SERVER_PORT)
}
