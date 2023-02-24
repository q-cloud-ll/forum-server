package forum

import (
	"fmt"
	"forum/dao/redis"
	"forum/global"
	"forum/model/common/response"
	"forum/utils/wx"
	"forum/utils/xerr"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"go.uber.org/zap"
)

type WeChatApi struct{}

// FrmGetQRCodeTicket 生成带参数二维码
func (qra *WeChatApi) FrmGetQRCodeTicket(c *gin.Context) {
	var accessToken string
	token, err := redis.FrmGetAccessToken()
	if err != nil {
		access, err := wx.GetAccess()
		if err != nil {
			response.FailWithMessage(xerr.CodeWxTicketFail, c)
			return
		}
		redis.FrmSetAccessToken(access)
		accessToken = access
	} else {
		accessToken = token
	}

	wxTicket, err := wechatService.FrmGenerateQRCode(accessToken)
	if err != nil {
		response.FailWithMessage(xerr.CodeWxTicketFail, c)
		return
	}
	response.OkWithDetailed(wxTicket, "获取二维码成功", c)
}

// FrmWxGetTicket 微信获取消息回复
func (qra *WeChatApi) FrmWxGetTicket(c *gin.Context) {
	server := global.GVA_WX.GetServer(c.Request, c.Writer)
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		//TODO 对接收到的消息以及处理

		//text := message.NewText("扫码成功！")
		if msg.Event == "SCAN" {
			//text = message.NewText("扫码成功！")
			_, err := redis.FrmSetEvent(fmt.Sprintf("%v", msg.EventKey), fmt.Sprintf("%v", msg.FromUserName))
			if err != nil {
				global.GVA_LOG.Error("Wx scan message failed", zap.Error(err))
				return &message.Reply{MsgType: message.MsgTypeText}
			}
		}
		return &message.Reply{MsgType: message.MsgTypeText}
		//return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		global.GVA_LOG.Error("Wx message server failed", zap.Error(err))
		return
	}
	_ = server.Send()
}

func (qra *WeChatApi) FrmSignInfo(c *gin.Context) {
	url := c.Query("url")
	js := global.GVA_WX.GetJs()
	config, err := js.GetConfig(url)
	if err != nil {
		global.GVA_LOG.Error("get wx sign failed", zap.Error(err))
		response.FailWithMessage(xerr.CodeWxGzhSignFail, c)
		return
	}
	response.OkWithData(config, c)
}
