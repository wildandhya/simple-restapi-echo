package routes

import (
	"go-echo/controllers"

	"go-echo/middleware"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	auth := e.Group("/auth")

	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	user := e.Group("/user", middleware.IsAuthenticate)

	user.GET("/all", controllers.GetAllUser)
	user.POST("/", controllers.InsertPegawai)
	user.PUT("/update", controllers.UpdateUser)
	user.DELETE("/:id", controllers.DeleteUser)

	return e
}
