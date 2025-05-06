package command

import (
	"fmt"
	"mbot/ZTAPI"
	"mbot/model"
	"mbot/stock"
)

type StockCommand struct {
	BaseCommand *model.BaseCommand
}

func (c *StockCommand) GetName() string {
	return c.BaseCommand.Name
}

func NewStockCommand() *StockCommand {
	err := ZTAPI.Query_all_stock()
	if err != nil {
		fmt.Println("stock command err:%s", err)
	}
	return &StockCommand{
		BaseCommand: &model.BaseCommand{
			Name:        "stock",
			Description: "querying stock real time info",
		},
	}
}

func (c *StockCommand) Execute(args []string) (body string, msgtype string) {
	// args:[stock 002201]
	IsStock, name, code := IsStock(args[1])
	if !IsStock {
		return fmt.Sprintf(`{"text":"此代码不存在,请确认！"}`), "text"
	}
	stockInfo, err := ZTAPI.Quert_stock_info(name, code)
	if err != nil {
		return fmt.Sprintf(`{"text":"遇到错误: %s"}`, err), "text"
	}
	stock.AnalyRate(stockInfo)
	color := "red"
	if stockInfo.StockInfo.LastPrice < stockInfo.StockInfo.LastClose {
		color = "green"
	}
	fmt.Println(color)
	body = fmt.Sprintf(`{"schema":"2.0","config":{"update_multi":true,"style":{"text_size":{"normal_v2":{"default":"normal","pc":"normal","mobile":"heading"}}}},"body":{"direction":"vertical","padding":"12px 12px 12px 12px","elements":[{"tag":"column_set","background_style":"grey-100","horizontal_spacing":"8px","horizontal_align":"left","columns":[{"tag":"column","width":"auto","elements":[{"tag":"div","text":{"tag":"plain_text","content":"🚀  %s ","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"}],"vertical_spacing":"8px","horizontal_align":"left","vertical_align":"top"},{"tag":"column","width":"weighted","elements":[{"tag":"div","text":{"tag":"plain_text","content":"%s","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"}],"vertical_spacing":"8px","horizontal_align":"left","vertical_align":"top","weight":1}],"margin":"0px 0px 0px 0px"},{"tag":"column_set","background_style":"grey-100","horizontal_spacing":"8px","horizontal_align":"left","columns":[{"tag":"column","width":"auto","elements":[],"vertical_spacing":"8px","horizontal_align":"left","vertical_align":"top"},{"tag":"column","width":"auto","elements":[{"tag":"div","text":{"tag":"plain_text","content":"当前价格","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"},{"tag":"div","text":{"tag":"plain_text","content":"%0.2f","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"}],"vertical_spacing":"8px","horizontal_align":"left","vertical_align":"top"},{"tag":"column","width":"auto","elements":[{"tag":"div","text":{"tag":"plain_text","content":"涨幅","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"},{"tag":"div","text":{"tag":"plain_text","content":"%s","text_size":"normal_v2","text_align":"left","text_color":"%s"},"margin":"0px 0px 0px 0px"}],"padding":"0px 0px 0px 0px","direction":"vertical","horizontal_spacing":"8px","vertical_spacing":"8px","horizontal_align":"left","vertical_align":"top","margin":"0px 0px 0px 0px"},{"tag":"column","width":"auto","elements":[{"tag":"div","text":{"tag":"plain_text","content":"昨日收盘","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"},{"tag":"div","text":{"tag":"plain_text","content":"%0.2f","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"}],"vertical_spacing":"8px","horizontal_align":"left","vertical_align":"top"},{"tag":"column","width":"weighted","elements":[{"tag":"div","text":{"tag":"plain_text","content":"数据更新时间","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"},{"tag":"div","text":{"tag":"plain_text","content":"%s","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"}],"vertical_spacing":"8px","horizontal_align":"left","vertical_align":"top","weight":1}],"margin":"0px 0px 0px 0px"},{"tag":"div","text":{"tag":"plain_text","content":"%s","text_size":"normal_v2","text_align":"left","text_color":"default"},"margin":"0px 0px 0px 0px"}]}}`, stockInfo.Code, stockInfo.Name, stockInfo.StockInfo.LastPrice, stockInfo.TodayNew, color, stockInfo.StockInfo.LastClose, stockInfo.StockInfo.T, "这个位置留着写名人名言")

	return body, "interactive"
}

func IsStock(stock string) (IsStock bool, name string, code string) {
	if ZTAPI.IsCode(stock) {
		name = ZTAPI.StockMapCodeToName[stock]
		code = stock
	} else {
		name = stock
		code = ZTAPI.StockMapNameToCode[stock]
	}
	if name == "" || code == "" {
		return false, name, code
	}
	return true, name, code

}
