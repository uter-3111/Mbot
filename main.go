package main

import (
	"context"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
	"mbot/im"
	"mbot/infra"
	"os"
	"strings"
)

func main() {
	infra.InitBot()
	// 注册事件回调，OnP2MessageReceiveV1 为接收消息 v2.0；OnCustomizedEvent 内的 message 为接收消息 v1.0。
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2MessageReceiveV1(im.ReplyByCommand).
		OnCustomizedEvent("", func(ctx context.Context, event *larkevent.EventReq) error {
			fmt.Printf("[ OnCustomizedEvent access ], type: message, data: %s\n", string(event.Body))
			return nil
		})
	// 创建Client
	if strings.Contains(os.Getenv("HOSTNAME"), "mbot") {
		infra.Loglevel = larkcore.LogLevelError
	} else {
		infra.Loglevel = larkcore.LogLevelDebug
	}
	fmt.Printf("LogLeve:%v \n", infra.Loglevel)
	cli := larkws.NewClient(infra.BotAppid, infra.BotSecret,
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(infra.Loglevel),
	)

	// 启动客户端
	fmt.Println("McDonald's start success")
	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
