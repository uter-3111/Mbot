package ZTAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mbot/infra"
	"mbot/model"
	"net/http"
	"strconv"
	"strings"
)

var StockList []model.Stock
var StockMapNameToCode = map[string]string{}
var StockMapCodeToName = map[string]string{}

// stock info
func Quert_stock_info(name string, code string) (stock *model.Stockfluctuation, err error) {
	var info *model.StockInfo
	stock = new(model.Stockfluctuation)
	stock.Name = name
	stock.Code = code
	url := fmt.Sprintf("https://api.zhituapi.com/hs/real/time/%v?token=%v", stock.Code, infra.ZTAPIToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	stock.StockInfo = info
	return
}

// 查询所有stock
func Query_all_stock() error {
	url := fmt.Sprintf("https://api.zhituapi.com/hs/list/all?token=%v", infra.ZTAPIToken)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &StockList)
	if err != nil {
		return err
	}
	ConvertToKV(StockList)
	return nil
}

// 转换为 map[stock代码]stock名称
func ConvertToKV(stocks []model.Stock) {
	for _, s := range stocks {
		StockMapNameToCode[s.Name] = s.Code
		if len(strings.Split(s.Code, ".")) > 1 {
			StockMapCodeToName[strings.Split(s.Code, ".")[0]] = s.Name // 以代码为键，名称为值
		} else {
			StockMapCodeToName[s.Code] = s.Name // 以代码为键，名称为值
		}
	}
}

func IsCode(stock string) bool {
	_, err := strconv.Atoi(stock)
	return err == nil
}
