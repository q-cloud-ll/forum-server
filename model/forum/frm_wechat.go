package forum

import "encoding/json"

type FrmQRCodeInfo struct {
	CodeUrl string `json:"codeUrl"`
}

type AuthInfo struct {
	Token  string `json:"token"`
	Status int    `json:"status"`
}

func (u *AuthInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

type FrmWXAppConfig struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Token     string `json:"token"`
}

type FrmWxTokenMessages struct {
	Ticket  string `json:"ticket"`
	Url     string `json:"url"`
	StrData string `json:"str_data"`
	ErrMsg  string `json:"err_msg"`
	ErrCode int    `json:"err_code"`
}