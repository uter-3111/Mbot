package infra

import (
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"os"
	"strings"
)

var BotAppid string
var BotSecret string
var ImClient *lark.Client

func InitBot() {
	if strings.Contains(os.Getenv("HOSTNAME"), "mbot") {
		BotAppid = ProdBotAppid
		BotSecret = ProdBotSecret
	} else {
		//BotAppid = TestBotAppid
		//BotSecret = TestBotSecret
		BotAppid = ProdBotAppid
		BotSecret = ProdBotSecret
	}
	fmt.Printf("BotAppid: %v\n", BotAppid)
	ImClient = lark.NewClient(BotAppid, BotSecret)
}
