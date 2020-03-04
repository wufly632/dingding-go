package dingding

type Ding struct {
	MsgType string // text
	Token   string
	Secret  string
}

func sign() string {
	return ""
}

func getUri() {
	signStr := sign()
}

// 发送消息
func (ding *Ding) Send(content string) error {

}
