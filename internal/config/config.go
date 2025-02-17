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
	// 添加自定义配置项
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int
	}
	// 日志配置
	Log struct {
		Level      string // 日志级别：debug, info, warn, error
		FilePath   string // 日志文件路径
		MaxSize    int    // 每个日志文件的最大大小（MB）
		MaxBackups int    // 保留的旧日志文件最大数量
		MaxAge     int    // 保留的旧日志文件最大天数
		Compress   bool   // 是否压缩旧日志文件
	}
}

// LoadConfig 根据环境加载对应的配置文件
func LoadConfig() (*Config, error) {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev" // 默认使用开发环境配置
	}

	configFile := fmt.Sprintf("%s.toml", env)
	configData, err := configFS.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var cfg Config
	if err := toml.Unmarshal(configData, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &cfg, nil
}
