package config

type SnowFlake struct {
	StartTime string `mapstructure:"start-time" json:"start-time" yaml:"start-time"`
	MachineID int64  `mapstructure:"machine-id" json:"machine-id" yaml:"machine-id"`
}
