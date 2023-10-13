package repository

import (
	"gorm.io/gorm"
)

type stockDataRepository struct {
	db *gorm.DB
}

func NewStockDataRepository(db *gorm.DB) *stockDataRepository {
	s := stockDataRepository{
		db: db,
	}
	return &s
}

func (r *stockDataRepository) GetStockDatas() []StockData {
	var stockdatas = new([]StockData)
	result := r.db.Find(&stockdatas)
	if result.Error != nil {
		panic(result.Error)
	}
	return *stockdatas
}

func (r *stockDataRepository) PostStockData(s StockData) {
	result := r.db.Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
}
