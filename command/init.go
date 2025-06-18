package command

import (
	"crypto/tls"
	"net/http"
)

var Registry *CommandRegistry

func init() {
	Registry = NewCommandRegistry()

	// 注册命令
	Registry.Register(NewWeatherCommand())
	Registry.Register(NewTodoCommand())
	Registry.Register(NewStockCommand())
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证
	}
}
