package service

import "go-echo-stockport/repository"

type stockDataService struct {
	stockrepo repository.IStockDataRepository
}

func NewStockDataService(r repository.IStockDataRepository) *stockDataService {
	s := stockDataService{
		stockrepo: r,
	}
	return &s
}

func (service *stockDataService) GetStockDatas() []StockDataDto {
	repo_stocks := service.stockrepo.GetStockDatas()
	ret_stockdtolist := make([]StockDataDto, len(repo_stocks))
	for i, v := range repo_stocks {
		stockdto := StockDataDto{
			StockSymbol: v.StockSymbol,
			StockName:   v.StockName,
			Description: v.Description,
			Industry:    v.Industry,
			MarketCap:   v.MarketCap,
		}
		ret_stockdtolist[i] = stockdto
	}
	return ret_stockdtolist
}

func (service *stockDataService) PostStockData(s StockDataDto) {
	repo_stockdata := repository.StockData{
		StockSymbol: s.StockSymbol,
		StockName:   s.StockName,
		Description: s.Description,
		Industry:    s.Industry,
		MarketCap:   s.MarketCap,
	}
	service.stockrepo.PostStockData(repo_stockdata)
}
