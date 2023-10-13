package service

type IService interface {
	GetStockDatas() []StockDataDto
	PostStockData(s StockDataDto)
}

type ServiceInterfaceTest interface {
	GetHello() string
	UpdateHello(text string)
	GetAllUser() []UserDto
	AddUser(u *UserDto)
}

type StockDataDto struct {
	StockSymbol string  `json:"stock_symbol"`
	StockName   string  `json:"stock_name"`
	Description string  `json:"description"`
	Industry    string  `json:"industry"`
	MarketCap   float32 `json:"market_cap"`
}
