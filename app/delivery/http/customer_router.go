package http

import (
	"github.com/aulyarahman/twitcat-service/app"
	"github.com/aulyarahman/twitcat-service/app/delivery/controller"
	"github.com/labstack/echo"
	"net/http"
)

func RouterCustomer(e *echo.Echo, cusUseCase app.CustomerUseCaseI) {
	customerCtrl := controller.NewCustomerController(e, cusUseCase)

	r := e.Group("/api/v1/customer")
	r.GET("/", customerCtrl.Get)
	r.GET("/{customerId}", func(c echo.Context) error {
		id := c.Param("customerId")
		return c.String(http.StatusOK, id)
	})
	r.POST("/", customerCtrl.Insert)
	r.PUT("/", customerCtrl.Update)
	r.DELETE("/", customerCtrl.Delete)
}
