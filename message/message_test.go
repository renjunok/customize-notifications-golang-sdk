/*
@Time : 2021/5/7 上午2:53
@Author : RenJun
@File : message_test
@Description :
@CopyRight:
*/
package message

import (
	"net/http"
	"reflect"
	"testing"
)

const TestStringFormat = "\x1b[1;40;31m [fail] %s error = %v, wantErr %v \x1b[0m\n"

func TestGenerateNonceStr(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "GenerateNonceStr()", want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateNonceStr(); len(got) != tt.want {
				t.Errorf(TestStringFormat, tt.name, len(got), tt.want)
			}
		})
	}
}

func GenerateTestStr(count int) string {
	testStr := ""
	for i := 1; i <= count; i++ {
		testStr += "a"
	}
	return testStr
}

func TestMessage_Check(t *testing.T) {
	type fields struct {
		Title   string
		MsgType int
		Content string
		Group   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "title len zero", fields: fields{
			Title:   "",
			MsgType: 0,
			Content: "test",
			Group:   "",
		}, wantErr: true},

		{name: "title len > 100", fields: fields{
			Title:   GenerateTestStr(101),
			MsgType: 0,
			Content: "test",
			Group:   "",
		}, wantErr: true},

		{name: "title len meet the requirements", fields: fields{
			Title:   GenerateTestStr(50),
			MsgType: 0,
			Content: "test",
			Group:   "",
		}, wantErr: false},

		{name: "content len zero", fields: fields{
			Title:   "test",
			MsgType: 0,
			Content: "",
			Group:   "",
		}, wantErr: true},

		{name: "content len > 4000", fields: fields{
			Title:   "test",
			MsgType: 0,
			Content: GenerateTestStr(4001),
			Group:   "",
		}, wantErr: true},

		{name: "content len meet the requirements", fields: fields{
			Title:   "test",
			MsgType: 0,
			Content: GenerateTestStr(2000),
			Group:   "",
		}, wantErr: false},

		{name: "group len = 0", fields: fields{
			Title:   "test",
			MsgType: 0,
			Content: "test",
			Group:   "",
		}, wantErr: false},
		{name: "group len > 20 ", fields: fields{
			Title:   "test",
			MsgType: 0,
			Content: "test",
			Group:   GenerateTestStr(21),
		}, wantErr: true},
		{name: "group len = 10", fields: fields{
			Title:   "test",
			MsgType: 0,
			Content: "test",
			Group:   GenerateTestStr(10),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Title:   tt.fields.Title,
				MsgType: tt.fields.MsgType,
				Content: tt.fields.Content,
				Group:   tt.fields.Group,
			}
			if err := m.check(); (err != nil) != tt.wantErr {
				t.Errorf(TestStringFormat, tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestMessage_Send(t *testing.T) {
	type m struct {
		Title   string
		MsgType int
		Content string
		Group   string
	}
	type args struct {
		pushID     string
		pushSecret string
	}
	tests := []struct {
		name    string
		fields  m
		args    args
		wantErr bool
	}{
		// {name: "test json", fields: m{
		// 	Title:   "title",
		// 	MsgType: 0,
		// 	Content: "content",
		// 	Group:   "",
		// }, args: args{
		// 	pushID:     "abcdef",
		// 	pushSecret: "abcdefg",
		// }, wantErr: false},
		//
		// {name: "test json", fields: m{
		// 	Title:   "title",
		// 	MsgType: 0,
		// 	Group:   "",
		// }, args: args{
		// 	pushID:     "abcdef",
		// 	pushSecret: "abcdefg",
		// }, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Title:   tt.fields.Title,
				MsgType: tt.fields.MsgType,
				Content: tt.fields.Content,
				Group:   tt.fields.Group,
			}
			if err := m.Send(tt.args.pushID, tt.args.pushSecret); (err != nil) != tt.wantErr {
				t.Errorf(TestStringFormat, tt.name, err, tt.wantErr)
				// t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgType_typeConversion(t *testing.T) {
	tests := []struct {
		name string
		m    MsgType
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.typeConversion(); got != tt.want {
				t.Errorf("typeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMessage(t *testing.T) {
	type args struct {
		title   string
		content string
		msgType MsgType
		group   []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMessage(tt.args.title, tt.args.content, tt.args.msgType, tt.args.group...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseResponse(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *HttpResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseResponse(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmitMessageRequest(t *testing.T) {
	type args struct {
		m *SendMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SubmitMessageRequest(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("SubmitMessageRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
