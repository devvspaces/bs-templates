package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/go-todo-app/internal/config"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/tracer"
	"github.com/mecitsemerci/go-todo-app/internal/wired"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// AppServer application runtime
type AppServer struct {
	Close func() error
	Start func() error
}

func registerHandlers(apiRouteGroup *gin.RouterGroup) error {

	groupV1 := apiRouteGroup.Group("/v1")
	{
		todoHandler, err := wired.InitializeTodoHandler()
		if err != nil {
			return err
		}
		todoHandler.Register(groupV1)
	}

	healthHandler := wired.InitializeHealthHandler()
	healthHandler.Register(apiRouteGroup)

	return nil
}

// NewAppServer returns AppServer
func NewAppServer() (*AppServer, error) {
	router := gin.Default()

	//Application configuration
	config.Load()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		config.AppConfig.FrontendURL,
		config.AppConfig.FrontendURL + "/",
	}
	corsConfig.AllowCredentials = true

	// Middleware
	router.Use(
		cors.New(corsConfig),
	)

	apiRouteGroup := router.Group("/api")

	if err := registerHandlers(apiRouteGroup); err != nil {
		return nil, err
	}

	// Opentracing configuration
	traceCloser := tracer.Init()

	// Swagger Configuration
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))

	return &AppServer{
		Close: func() error {
			return traceCloser.Close()
		},
		Start: func() error {
			return router.Run()
		},
	}, nil
}
