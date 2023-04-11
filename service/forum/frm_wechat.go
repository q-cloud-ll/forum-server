package forum

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"forum/dao/redis"
	"forum/global"
	"forum/model/forum"
	"forum/utils/tool"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

type WeChatService struct{}

// FrmGenerateQRCode 生成二维码
func (qrs *WeChatService) FrmGenerateQRCode(token string) (mes forum.FrmWxTokenMessages, err error) {
	post := `{
	"expire_seconds":300,
	"action_name": "QR_LIMIT_STR_SCENE",
		"action_info": {
		"scene": {
			"scene_str": "%s"
			}
		}
	}`
	str := tool.Krand(16, tool.KC_RAND_KIND_ALL)
	_, err = redis.FrmGetWxSceneStr(str)
	if err != nil {
		_ = redis.FrmSetWxSceneStr(str)
		getTicketReq := fmt.Sprintf(post, str)
		var ticket forum.FrmWxTokenMessages
		var jsonTicketReq = []byte(getTicketReq)
		buffer := bytes.NewBuffer(jsonTicketReq)
		request, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="+token, buffer)
		if err != nil {
			global.GVA_LOG.Error("http.NewRequest failed,err:", zap.Error(err))
			return ticket, err
		}
		cli := http.Client{}
		resp, err := cli.Do(request.WithContext(context.TODO()))
		if err != nil {
			global.GVA_LOG.Error("cli.Do failed,err:", zap.Error(err))
			return ticket, err
		}
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			global.GVA_LOG.Error("ioutil.ReadAll failed,err:", zap.Error(err))
			return ticket, err
		}
		err = json.Unmarshal(respBytes, &ticket)
		if err != nil {
			global.GVA_LOG.Error("json.Unmarshal failed,err:", zap.Error(err))
			return ticket, err
		}
		ticket.StrData = str
		return ticket, nil
	} else {
		return forum.FrmWxTokenMessages{}, errors.New("请重新请求")
	}
}
