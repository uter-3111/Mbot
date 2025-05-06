package im

import (
	"context"
	"encoding/json"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"log"
	"mbot/command"
	"mbot/infra"
)

func ReplyByCommand(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	fmt.Printf("[ OnP2MessageReceiveV1 access ], data: %s\n", larkcore.Prettify(event.Event))
	fmt.Printf("[ OnP2MessageReceiveV1 access ], event MessageID: %s\n", larkcore.Prettify(event.Event.Message.MessageId))
	messageId := *event.Event.Message.MessageId
	content, err := GetTextContent(*event.Event.Message.Content)
	if err != nil {
		log.Printf("GetTextContent err: %v", err)
		return err
	}
	body, msgtype := command.Registry.ParseAndExecute(content)

	ReplyMessageCallBack(messageId, body, msgtype)
	return nil
}

func ReplyMessageCallBack(messageId string, body string, msgtype string) {
	req := larkim.NewReplyMessageReqBuilder().
		MessageId(messageId).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			Content(body).
			MsgType(msgtype).
			Build()).
		Build()

	// 发起请求
	resp, err := infra.ImClient.Im.V1.Message.Reply(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Printf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}

type Message struct {
	Text string `json:"text"`
}

func GetTextContent(content string) (message string, err error) {
	var msg Message
	err = json.Unmarshal([]byte(content), &msg)
	if err != nil {
		return "", err
	}
	return msg.Text, nil
}
