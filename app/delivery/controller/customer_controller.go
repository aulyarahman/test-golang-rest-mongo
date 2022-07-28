package controller

import (
	"context"
	"encoding/json"
	"github.com/aulyarahman/twitcat-service/app"
	"github.com/aulyarahman/twitcat-service/app/model"
	"github.com/labstack/echo"
)

type CustomerController struct {
	e       *echo.Echo
	usecase app.CustomerUseCaseI
}

func NewCustomerController(e *echo.Echo, usecase app.CustomerUseCaseI) *CustomerController {
	return &CustomerController{
		e:       e,
		usecase: usecase,
	}
}

func (c *CustomerController) Get(ec echo.Context) error {
	data, err := c.usecase.GetData(context.Background())
	if err != nil {
		return err
	}
	return ec.JSON(200, data)
}

func (c *CustomerController) GetById(ec echo.Context) error {
	var id model.CustomerById
	data, err := c.usecase.GetDataById(context.Background(), id)
	if err != nil {
		return err
	}
	return ec.JSON(200, data)
}

func (c *CustomerController) Insert(ec echo.Context) error {
	var req model.Customer
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	data, err := c.usecase.InsertData(context.Background(), req)
	if err != nil {
		return err
	}

	return ec.JSON(200, data)
}

func (c *CustomerController) Update(ec echo.Context) error {
	var req model.Customer
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	data, err := c.usecase.InsertData(context.Background(), req)
	if err != nil {
		return err
	}

	return ec.JSON(200, data)
}

func (c *CustomerController) Delete(ec echo.Context) error {
	var req model.Customer
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}

	data, err := c.usecase.DeleteData(context.Background(), req)
	if err != nil {
		return err
	}
	return ec.JSON(200, data)
}
