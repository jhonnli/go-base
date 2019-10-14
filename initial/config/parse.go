package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/jhonnli/golibs"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	_key = "?*-NMn5hJMXoTkm7=dFUYvUJu35UUN_&"
	_iv  = "$8^82_4nc=r045FN"
)

// InitConfig 读取并解析配置到Config对象
// 没有检查配置是否完整
func InitConfig() {
	DEPLOY_ENV := os.Getenv("DEPLOY_ENV")
	switch DEPLOY_ENV {
	case "dev":
		viper.SetConfigFile("config_dev.yaml")
	case "test":
		viper.SetConfigFile("config_test.yaml")
	case "prod":
		viper.SetConfigFile("config_prod.yaml")
	default:
		log.Fatalln("没有找到环境变量DEPLOY_ENV")
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 启动时解析出错，程序退出
	conf, err := parseConfig()
	if err != nil {
		os.Exit(1)
	}
	Config = *conf

	// 当配置文件被修改，自动重新读取
	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		if event.Op == fsnotify.Write {
			// 解析出错，拒绝错误的配置
			conf, err := parseConfig()
			if err != nil {
				return
			}
			Config = *conf
			log.Println("config reload.")
		}
	})
}

// parseConfig 解析并解密配置
func parseConfig() (*config, error) {
	var conf config
	//  解析
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	decode := func(str *string) {
		if err != nil {
			return
		}
		if *str == "" {
			return
		}

		var tmpBytes []byte
		tmpBytes, err = golibs.AesDecrypt(golibs.HexStringToBytes(*str), []byte(_key), []byte(_iv))
		if err != nil {
			return
		}
		*str = string(tmpBytes)
	}

	// 解密
	decode(&conf.DB.Pwd)
	decode(&conf.App.Secret)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
