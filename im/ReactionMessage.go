package im

import (
	"context"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"mbot/infra"
)

// Aciton_like 点赞所有消息
func Aciton_like(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	// todo 需要开通获取群组消息之后才可以用
	fmt.Printf("[ OnP2MessageReceiveV1 access ], data: %s\n", larkcore.Prettify(event.Event))
	fmt.Printf("[ OnP2MessageReceiveV1 access ], event MessageID: %s\n", larkcore.Prettify(event.Event.Message.MessageId))
	messageId := *event.Event.Message.MessageId
	ReactionCallBack(messageId)
	return nil
}

func ReactionCallBack(messageId string) {
	req := larkim.NewCreateMessageReactionReqBuilder().
		MessageId(messageId).
		Body(larkim.NewCreateMessageReactionReqBodyBuilder().
			ReactionType(larkim.NewEmojiBuilder().
				EmojiType(`THUMBSUP`).
				Build()).
			Build()).
		Build()

	// 发起请求
	resp, err := infra.ImClient.Im.V1.MessageReaction.Create(context.Background(), req)

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
