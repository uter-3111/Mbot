package model

// Command 接口定义命令执行方法
type Command interface {
	GetName() string
	Execute(args []string) (string, string)
}

// BaseCommand 基础结构体（可选）
type BaseCommand struct {
	Name        string
	Description string
}

type Stock struct {
	Name   string `json:"mc"`
	Code   string `json:"dm"`
	Bourse string `json:"jys"` // 交易所
}
type StockInfo struct {
	LastPrice float32 `json:"lastPrice"` // 最新价格
	Open      float32 `json:"open"`      // 开盘价
	High      float32 `json:"high"`      // 最高价
	Low       float32 `json:"low"`       // 最低价
	LastClose float32 `json:"lastClose"` // 昨日收盘价
	Amount    int     `json:"amount"`    // 成交总额
	Volume    int     `json:"volume"`    // 成交量
	Pvolume   int     `json:"pvolume"`   // 原始成交总量
	T         string  `json:"t"`         // 更新时间
}

type Stockfluctuation struct {
	Name      string
	Code      string
	TodayNew  string // 今日涨幅
	StockInfo *StockInfo
}
