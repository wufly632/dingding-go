package main

import (
	"github.com/CatchZeng/dingtalk/client"
	"github.com/CatchZeng/dingtalk/message"
)

func main() {
	dingTalk := client.DingTalk{
		AccessToken: "5600b43c34249a88599f344d56faadff3e010f495614ee1851b0b89723c88f13",
		Secret:      "SEC76c895c621f6bff08ce278164f2e23c994f029b695dca3008744e2f477430a49",
	}

	msg := message.NewTextMessage().SetContent("测试文本&at 某个人").SetAt([]string{"177010xxx60"}, false)
	dingTalk.Send(msg)
}
