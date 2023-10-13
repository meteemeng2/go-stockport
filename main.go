package main

import (
	"go-echo-stockport/handler"
	"go-echo-stockport/initializers"
	"go-echo-stockport/repository"
	"go-echo-stockport/service"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	e := echo.New()

	// var_repository := repository.NewRepositoryTest("test from repo")
	// var_service := service.NewServiceTest(var_repository)
	// var_handler := handler.NewHandlerTest(var_service)
	// echosetting_test(e, var_handler)

	var_repository := repository.NewStockDataRepository(initializers.DB)
	var_service := service.NewStockDataService(var_repository)
	var_handler := handler.NewStockDataHandler(var_service)
	echosetting(e, var_handler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func echosetting(e *echo.Echo, h handler.IStockDatahandler) {
	e.GET("/StockData", h.GetStockDatas())
	e.POST("/StockData", h.PostStockData())
}

func echosetting_test(e *echo.Echo, h handler.HandlerInterfaceTest) {
	e.GET("/gethello", h.HandleGetHello())
	e.GET("/updatehello/:txt", h.HandleUpdateHello())
	e.GET("/getalluser", h.HandleGetAllUser())
	e.POST("/adduser", h.HandleAddUser())
}
