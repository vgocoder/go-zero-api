package config

import (
	"embed"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/zeromicro/go-zero/rest"
)

//go:embed *.toml
var configFS embed.FS

type Config struct {
	rest.RestConf
	// 数据库配置
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
	} `json:"database" yaml:"database"`

	// Redis配置
	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int
	} `json:"redis" yaml:"redis"`

	// 日志配置
	CustomLog struct {
		Mode  string
		Level string
		Path  string
	} `json:"customLog" yaml:"customLog"`
}

// LoadConfig 从指定路径加载配置文件
func LoadConfig(configFile string) (*Config, error) {
	var c Config

	// 如果配置文件路径为空，使用默认的开发环境配置
	if configFile == "" {
		content, err := configFS.ReadFile("dev.toml")
		if err != nil {
			return nil, fmt.Errorf("failed to read default config: %v", err)
		}
		if _, err := toml.Decode(string(content), &c); err != nil {
			return nil, fmt.Errorf("failed to decode default config: %v", err)
		}
		return &c, nil
	}

	// 从指定路径加载配置文件
	content, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}
	if _, err := toml.Decode(string(content), &c); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &c, nil
}
