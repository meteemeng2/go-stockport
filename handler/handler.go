package handler

import (
	"github.com/labstack/echo/v4"
)

type IStockDatahandler interface {
	GetStockDatas() echo.HandlerFunc
	PostStockData() echo.HandlerFunc
}

type HandlerInterfaceTest interface {
	HandleGetHello() echo.HandlerFunc
	HandleUpdateHello() echo.HandlerFunc
	HandleGetAllUser() echo.HandlerFunc
	HandleAddUser() echo.HandlerFunc
}
