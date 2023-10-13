package handler

import (
	"go-echo-stockport/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type stockDataHandler struct {
	service service.IService
}

func NewStockDataHandler(s service.IService) *stockDataHandler {
	sdh := stockDataHandler{
		service: s,
	}
	return &sdh
}

func (sdh *stockDataHandler) GetStockDatas() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, sdh.service.GetStockDatas())
	}
}

func (sdh *stockDataHandler) PostStockData() echo.HandlerFunc {
	return func(c echo.Context) error {
		sdto := new(service.StockDataDto)
		if err := c.Bind(sdto); err != nil {
			return err
		}
		sdh.service.PostStockData(*sdto)
		return c.NoContent(http.StatusOK)
	}
}
