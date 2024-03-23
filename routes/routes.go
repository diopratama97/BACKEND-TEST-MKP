package route

import (
	"backend-test-mkp/controllers"
	"backend-test-mkp/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Backend is run ...")
	})

	//auth
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	e.GET("/show", controllers.ListShow, middleware.IsAuth)
	e.POST("/show", controllers.CreateShow, middleware.IsAuth)
	e.PUT("/show/:id", controllers.UpdateShow, middleware.IsAuth)
	e.DELETE("/show/:id", controllers.DeleteShow, middleware.IsAuth)
	e.GET("/show/:id", controllers.DetailShow, middleware.IsAuth)

	return e
}
