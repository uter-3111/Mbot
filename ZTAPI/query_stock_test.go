package ZTAPI

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetAll(t *testing.T) {
	err := Query_all_stock()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(StockMapNameToCode["九鼎新材"])
	fmt.Println(StockMapCodeToName["002201"])

}

func TestGetInfo(t *testing.T) {
	err := Query_all_stock()
	if err != nil {
		fmt.Println("get stock list err，panic")
	}
	info, err := Quert_stock_info("九鼎新材", "002201")
	if err != nil {
		fmt.Println(err)
	}
	info.TodayNew = fmt.Sprintf("%.2f%%", ((info.StockInfo.LastPrice-info.StockInfo.LastClose)/info.StockInfo.LastClose)*100)

	fmt.Println(info, info.StockInfo)

}

func TestConvertToKV(t *testing.T) {
	a := "001.SH"
	b := strings.Split(a, ".")
	fmt.Println(b)
}
