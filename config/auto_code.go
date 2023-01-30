package config

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          string `mapstructure:"forum-server" json:"forum-server" yaml:"forum-server"`
	SApi            string `mapstructure:"forum-server-api" json:"forum-server-api" yaml:"forum-server-api"`
	SPlug           string `mapstructure:"forum-server-plug" json:"forum-server-plug" yaml:"forum-server-plug"`
	SInitialize     string `mapstructure:"forum-server-initialize" json:"forum-server-initialize" yaml:"forum-server-initialize"`
	SModel          string `mapstructure:"forum-server-model" json:"forum-server-model" yaml:"forum-server-model"`
	SRequest        string `mapstructure:"forum-server-request" json:"forum-server-request"  yaml:"forum-server-request"`
	SRouter         string `mapstructure:"forum-server-router" json:"forum-server-router" yaml:"forum-server-router"`
	SService        string `mapstructure:"forum-server-service" json:"forum-server-service" yaml:"forum-server-service"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	WApi            string `mapstructure:"web-api" json:"web-api" yaml:"web-api"`
	WForm           string `mapstructure:"web-form" json:"web-form" yaml:"web-form"`
	WTable          string `mapstructure:"web-table" json:"web-table" yaml:"web-table"`
}
