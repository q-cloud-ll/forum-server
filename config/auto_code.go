package config

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          string `mapstructure:"forum" json:"forum" yaml:"forum"`
	SApi            string `mapstructure:"forum-api" json:"forum-api" yaml:"forum-api"`
	SPlug           string `mapstructure:"forum-plug" json:"forum-plug" yaml:"forum-plug"`
	SInitialize     string `mapstructure:"forum-initialize" json:"forum-initialize" yaml:"forum-initialize"`
	SModel          string `mapstructure:"forum-model" json:"forum-model" yaml:"forum-model"`
	SRequest        string `mapstructure:"forum-request" json:"forum-request"  yaml:"forum-request"`
	SRouter         string `mapstructure:"forum-router" json:"forum-router" yaml:"forum-router"`
	SService        string `mapstructure:"forum-service" json:"forum-service" yaml:"forum-service"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	WApi            string `mapstructure:"web-api" json:"web-api" yaml:"web-api"`
	WForm           string `mapstructure:"web-form" json:"web-form" yaml:"web-form"`
	WTable          string `mapstructure:"web-table" json:"web-table" yaml:"web-table"`
}
