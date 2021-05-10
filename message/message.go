/*
@Time : 2021/5/6 下午11:58
@Author : RenJun
@File : message
@Description :
@CopyRight:
*/
package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const CustomizeNotificationUrl = "https://api.msg.launch.im/message"

type MsgType int

const (
	primary MsgType = iota
	success
	info
	warning
	fail
)

type Message struct {
	Title     string    `json:"title"`
	MsgType   int       `json:"msg_type"`
	Content   string    `json:"content"`
	Group     string    `json:"group,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type SendMessage struct {
	PushID    string `json:"push_id"`
	Nonce     string `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	Message   `json:"message"`
}

func (m MsgType) typeConversion() int {
	switch m {
	case primary:
		return 0
	case success:
		return 1
	case info:
		return 2
	case warning:
		return 3
	case fail:
		return 4
	default:
		return 0
	}
}

func NewMessage(title, content string, msgType MsgType, group ...string) (*Message, error) {
	msg := Message{}
	msg.Title = title
	msg.Content = content
	msg.MsgType = msgType.typeConversion()

	if len(group) > 0 {
		msg.Group = group[0]
	}

	if err := msg.check(); err != nil {
		return nil, err
	}

	return &msg, nil
}

func (m *Message) Send(pushID, pushSecret string) error {
	nonceStr := GenerateNonceStr()
	timestamp := time.Now().Unix()
	msgJson, err := json.Marshal(&m)
	if err != nil {
		return err
	}

	p := make(ParamSign)
	p["push_id"] = pushID
	p["nonce"] = nonceStr
	p["timestamp"] = strconv.FormatInt(timestamp, 10)
	p["message"] = string(msgJson)
	signStr, err := p.sign(pushSecret)
	if err != nil {
		return err
	}

	sm := &SendMessage{
		PushID:    pushID,
		Nonce:     nonceStr,
		Timestamp: timestamp,
		Sign:      signStr,
		Message:   *m,
	}

	return SubmitMessageRequest(sm)
}

func SubmitMessageRequest(m *SendMessage) error {
	b, err := json.Marshal(&m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", CustomizeNotificationUrl, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return errors.New("server response error")
	}
	result, err := ParseResponse(resp)
	if err != nil {
		return err
	}
	if result.Code >= 400 {
		return errors.New(result.Error)
	}
	return nil
}

type HttpResponse struct {
	Code    int    `json:"code"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func ParseResponse(resp *http.Response) (*HttpResponse, error) {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	hr := &HttpResponse{}
	if err = json.Unmarshal(respBody, &hr); err != nil {
		return nil, err
	}

	return hr, nil
}

func (m *Message) check() error {
	if len(m.Title) > 100 || len(m.Title) < 1 {
		return errors.New("title char count is between 1 and 100")
	}

	if len(m.Content) > 4000 || len(m.Content) < 1 {
		return errors.New("content char count is between 1 and 4000")
	}

	if len(m.Group) > 20 {
		return errors.New("group char count not over 20")
	}

	return nil
}

func GenerateNonceStr() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytesStr := []byte(str)
	var result []byte
	randNonceStr := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		result = append(result, bytesStr[randNonceStr.Intn(len(bytesStr))])
	}
	return string(result)
}
