package infra

import (
	"encoding/json"
	"fmt"
	"mbot/model"
	"testing"
)

func TestInitlist(t *testing.T) {
	var StockList []model.Stock
	err := json.Unmarshal([]byte(Stock_all), &StockList)
	fmt.Println(err)
	fmt.Println("get tock list success for local", StockList[0].Name)
}
