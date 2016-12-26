package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

func hello(ctx *context.Context) {
	//配置微信参数
	config := &wechat.Config{
		AppID:          "wx60bf1662294a3444",
		AppSecret:      "1f3f1f506e63c4bcd0c2f2d88281cf33",
		Token:          "huxuesongtoken",
		EncodingAESKey: "EAOxf3lgmlfLzFsyZm346qXA9joUuG9Nmll7HUHNUep",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息

		text := message.NewText("我已收到您的消息会尽快处理")

		switch msg.MsgType {
		//文本消息
		case message.MsgTypeText:
		//do something
			fmt.Print("文本消息")
			//return &message.Reply{message.NewArticle("那个嫁给了穷人的女同学", "爆漫画", "http://zxpic.gtimg.com/infonew/0/wechat_pics_-10190065.static/640", "http://v.juhe.cn/weixin/redirect?wid=wechat_20161226038185")}
		//图片消息
		case message.MsgTypeImage:
		//do something
			fmt.Print("图片消息")
		//语音消息
		case message.MsgTypeVoice:
		//do something
			fmt.Print("语音消息")
		//视频消息
		case message.MsgTypeVideo:
		//do something
			fmt.Print("视频消息")
		//小视频消息
		case message.MsgTypeShortVideo:
		//do something
			fmt.Print("小视频消息")
		//地理位置消息
		case message.MsgTypeLocation:
		//do something
			fmt.Print("地理位置消息")
		//链接消息
		case message.MsgTypeLink:
		//do something
			fmt.Print("连接消息")
		//事件推送消息
		case message.MsgTypeEvent:
			fmt.Print("事件推送消息")
		}
		return &message.Reply{message.MsgTypeLink ,message.NewArticle("那个嫁给了穷人的女同学", "爆漫画", "http://zxpic.gtimg.com/infonew/0/wechat_pics_-10190065.static/640", "http://v.juhe.cn/weixin/redirect?wid=wechat_20161226038185")}
		//return &message.Reply{message.MsgTypeText, text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	beego.Any("/wx_connect", hello)
	beego.Run(":80")
}
