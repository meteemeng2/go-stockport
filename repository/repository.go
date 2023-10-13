package repository

type IStockDataRepository interface {
	GetStockDatas() []StockData
	PostStockData(s StockData)
}

type RepositoryInterfaceTest interface {
	GetHello() string
	UpdateHello(text string)
	GetAllUser() []User
	AddUser(u User)
}

type StockData struct {
	StockSymbol string  `json:"stock_symbol" gorm:"column:stock_symbol"`
	StockName   string  `json:"stock_name" gorm:"column:stock_name"`
	Description string  `json:"description" gorm:"column:description"`
	Industry    string  `json:"industry" gorm:"column:industry"`
	MarketCap   float32 `json:"market_cap" gorm:"column:market_cap"`
}
