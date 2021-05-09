/*
@Time : 2021/5/7 上午4:43
@Author : RenJun
@File : example
@Description :
@CopyRight:
*/
package main

import (
	"github.com/renjunok/customize-notifications-golang-sdk/message"
)

func main() {
	m, err := message.NewMessage("test title", "test content", 0, "developer group")
	if err != nil {
		// handler err
		return
	}
	err = m.Send("abcdef", "abcdefg")
	if err != nil {
		// handler err
		return
	}
}
