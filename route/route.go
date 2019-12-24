package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/letanthang/nc_user/config"
	"github.com/letanthang/nc_user/db"
	"github.com/letanthang/nc_user/handler"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {
	g := e.Group("/api/user/v1/private")
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(config.Config.JWTSecret.JWTKey),
		Claims:     &db.UserClaims{},
	}
	g.Use(middleware.JWTWithConfig(JWTConfig))
	g.PUT("/user", handler.UpdateUser)
}

func Staff(e *echo.Echo) {
	//g := e.Group("/api/user/v1/staff")
	//g.POST("/student", handler.AddStudent)
	//g.PUT("/student", handler.UpdateStudent)
}

func Public(e *echo.Echo) {
	g := e.Group("/api/user/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.POST("/user/register", handler.RegisterUser)

}
