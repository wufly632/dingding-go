package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"
)

const DingTalkOApi = "oapi.dingtalk.com"

var dingTalkURL url.URL = url.URL{
	Scheme: "https",
	Host:   DingTalkOApi,
	Path:   "robot/send",
}

var timestamp = strconv.FormatInt(time.Now().Unix()*1000, 10)

func GetDingTalkURL(accessToken string, secret string) (string, error) {
	dingTu := dingTalkURL
	value := url.Values{}
	value.Set("access_token", accessToken)
	value.Set("timestamp", timestamp)
	sign, err := sign(timestamp, secret)
	if err != nil {
		dingTu.RawQuery = value.Encode()
		return dingTu.String(), err
	}
	value.Set("sign", sign)
	dingTu.RawQuery = value.Encode()
	return dingTu.String(), nil

}

func sign(timestamp string, secret string) (string, error) {
	// 组数据
	strToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	// 加密
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, strToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
