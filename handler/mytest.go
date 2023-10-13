package handler

import (
	"go-echo-stockport/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlerTest struct {
	service service.ServiceInterfaceTest
}

func NewHandlerTest(s service.ServiceInterfaceTest) *handlerTest {
	hdlt := handlerTest{
		service: s,
	}
	return &hdlt
}

func (eh *handlerTest) HandleGetHello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, eh.service.GetHello())
	}
}

func (eh *handlerTest) HandleUpdateHello() echo.HandlerFunc {
	return func(c echo.Context) error {
		txt := c.Param("txt")
		eh.service.UpdateHello(txt)
		return c.NoContent(http.StatusOK)
	}
}

func (eh *handlerTest) HandleGetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, eh.service.GetAllUser())
	}
}

func (eh *handlerTest) HandleAddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(service.UserDto)
		if err := c.Bind(u); err != nil {
			return err
		}
		eh.service.AddUser(u)
		return c.NoContent(http.StatusOK)
	}
}
