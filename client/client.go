package client

import (
	"bytes"
	"dingding/message"
	"dingding/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DingTalk struct {
	AccessToken string
	Secret      string
}

type Response struct {
	ErrMsg  string `json:"err_msg"`
	ErrCode int64  `json:"err_code"`
}

const httpTimoutSecond = time.Duration(30) * time.Second

// 发送消息
func (dt *DingTalk) Send(message message.Message) (Response, error) {
	res := Response{}
	reqByte, err := message.ToByte()
	if err != nil {
		return res, err
	}

	// 获取推送的url
	pushUrl, err := security.GetDingTalkURL(dt.AccessToken, dt.Secret)
	if err != nil {
		return res, err
	}
	req, err := http.NewRequest("POST", pushUrl, bytes.NewReader(reqByte))
	if err != nil {
		return res, err
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	client.Timeout = httpTimoutSecond

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
	}

	if res.ErrCode != 0 {
		return res, fmt.Errorf("send message to dingtalk error = %s", res.ErrMsg)
	}

	return res, nil
}
