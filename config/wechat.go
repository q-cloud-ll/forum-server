package config

// WeChat 微信公众号配置
type WeChat struct {
	AppID     string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
	Token     string `mapstructure:"token" json:"token" yaml:"token"`
}
