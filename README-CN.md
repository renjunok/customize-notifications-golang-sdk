# Customize Notifications SDK for go

[![Go](https://github.com/renjunok/customize-notifications-golang-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/renjunok/customize-notifications-golang-sdk/actions/workflows/ci.yml)

## [README in English](https://github.com/renjunok/customize-notifications-golang-sdk/blob/main/README.md)

## 关于
> - 此Go SDK基于定制通知App应用API构建
> - 使用此SDK，用户可以在任意时间发送事件消息给自己

## 版本
> - v1.0.0

## 运行环境
> - Go 1.5及以上

## 安装方法
### GitHub安装
> - 执行命令`go get github.com/renjunok/customize-notifications-golang-sdk`获取远程代码包。
> - 在您的代码中使用`import "github.com/renjunok/customize-notifications-golang-sdk/message"`引入 Go SDK的包。
>
## 使用方法
> - iOS用户在AppStore商城安装App应用
> - 打开应用后点击左上角菜单获取配置信息, ID 和 Secret
> - 在你需要自定义通知信息中执行下列代码

	m, err := message.NewMessage("test title", "test content", 0, "developer group")
	if err != nil {
		// handler err
		return
	}
	err = m.Send("your id", "your secret")
	if err != nil {
		// handler err
		return
	}

## 注意事项
> - title, content, msgType 字段必填, group选填
> - title字符数最大值100, content最大值4000, group最大值20
> - msgType类型分别为 0 primary, 1 success, 2 info, 3 warning, 4 fail
> - 接口调用速率1分钟内最多三次，超出调用的请求不处理
