package config

import (
	"fortune/pkg/log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil

}

func (c *Config) initConfig() error {
	if c.Name != "" {
		//如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		//如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}

func (c *Config) initLog() {
	// 从配置文件读取日志级别，默认为 info
	logLevelStr := viper.GetString("log.level")
	if logLevelStr == "" {
		logLevelStr = "info"
	}

	// 解析日志级别
	var logLevel log.Level
	switch logLevelStr {
	case "debug":
		logLevel = log.DEBUG
	case "info":
		logLevel = log.INFO
	case "warn":
		logLevel = log.WARN
	case "error":
		logLevel = log.ERROR
	case "dpanic":
		logLevel = log.DPANIC
	case "panic":
		logLevel = log.PANIC
	case "fatal":
		logLevel = log.FATAL
	default:
		logLevel = log.INFO
	}

	// 使用配置的日志级别初始化日志
	log.Init(log.GetConsoleZapcore(logLevel))
	log.Infof("Log level set to: %s", logLevelStr)
}
