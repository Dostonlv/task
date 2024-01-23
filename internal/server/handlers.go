package server

import (
	"strings"

	blogsHttp "github.com/Dostonlv/task/blogs/delivery/http"
	"github.com/Dostonlv/task/docs"
	"github.com/Dostonlv/task/internal/blogs/repository"
	"github.com/Dostonlv/task/internal/blogs/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	MAX_HEADER_SIZE = 1 << 20 // 1 MB
	STACK_SIZE      = 1 << 10 // 1 KB
	BODY_LIMIT      = "2M"
	GZIP_LEVEL      = 5
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	cRepo := repository.NewBlogsRepository(s.db)

	blogUC := usecase.NewBlogsUseCase(s.cfg, cRepo, s.logger)

	// Init handlers
	blogHandlers := blogsHttp.NewBlogsHandlers(s.cfg, blogUC, s.logger)

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Blog and News API."
	docs.SwaggerInfo.Description = "Blog and News API Server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         STACK_SIZE,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: GZIP_LEVEL,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit(BODY_LIMIT))

	v1 := s.echo.Group("/v1")
	blogsHttp.MapHandlers(v1, blogHandlers)

	return nil
}
