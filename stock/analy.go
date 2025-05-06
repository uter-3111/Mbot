package stock

import (
	"fmt"
	"mbot/model"
)

func AnalyRate(stock *model.Stockfluctuation) {
	stock.TodayNew = fmt.Sprintf("%.2f%%", ((stock.StockInfo.LastPrice-stock.StockInfo.LastClose)/stock.StockInfo.LastClose)*100)
}
