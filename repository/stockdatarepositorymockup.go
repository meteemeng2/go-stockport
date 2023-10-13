package repository

type stockDataRepositoryMockup struct {
	stockdatalist []StockData
}

func NewStockDataRepositoryMockup() *stockDataRepositoryMockup {
	s := stockDataRepositoryMockup{}
	return &s
}

func (r *stockDataRepositoryMockup) GetStockDatas() []StockData {
	copied_stockdata := make([]StockData, len(r.stockdatalist))
	copy(copied_stockdata, r.stockdatalist)
	return copied_stockdata
}

func (r *stockDataRepositoryMockup) PostStockData(s StockData) {
	r.stockdatalist = append(r.stockdatalist, s)
}
