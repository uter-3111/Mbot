package im

import lark "github.com/larksuite/oapi-sdk-go/v3"

var Client *lark.Client

func init() {
	Client = lark.NewClient("cli_a75d57eaf3b0901c", "Ppvsb2hV7hfCHbveppi4XdSK7cbQM7Jv")
}
