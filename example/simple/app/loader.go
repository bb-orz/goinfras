package main

import (
	"errors"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
)

func viperLoader() *viper.Viper {
	var err error
	var runtimeViper *viper.Viper
	var fileName, configFilePath, configFileName, configFileExt string

	runtimeViper = viper.New()

	// 1. 从环境变量导入配置项
	err = loadConfigFromEnv(runtimeViper)
	if err != nil {
		panic("Viper Loading ENV Error:" + err.Error())
	}

	// 2. 从配置文件导入配置项
	if flagConfigFile != "" {
		configFilePath, fileName = filepath.Split(flagConfigFile)
		configFileExt = path.Ext(fileName)
		configFileName = fileName[0 : len(fileName)-len(configFileExt)]
		configFileExt = configFileExt[1:]
		err = loadConfigFromFile(runtimeViper, configFilePath, configFileName, configFileExt)
		if err != nil {
			panic("Viper Loading Config File Error:" + err.Error())
		}
	}

	// 3. 从远程配置系统导入配置项
	if flagRemoteProvider != "" || flagRemoteEndpoint != "" || flagRemotePath != "" {
		err = loadConfigFromRemote(runtimeViper, flagRemoteProvider, flagRemoteEndpoint, flagRemotePath)
		if err != nil {
			panic("Viper Loading Remote Config Error:" + err.Error())
		}
	}

	return runtimeViper
}

// Viper读取环境变量
func loadConfigFromEnv(runtimeViper *viper.Viper) error {
	// 若有需要读取环境变量
	// runtimeViper.SetEnvPrefix("env_prefix") // 添加需加载系统环境变量的前缀
	// runtimeViper.AllowEmptyEnv(true) // 是否允许环境变量为空值,默认为false
	// err := runtimeViper.BindEnv("your env var") // 绑定特定环境变量值到viper
	// runtimeViper.AutomaticEnv() // 自动载入所有环境变量，如设置SetEnvPrefix，则加载有特定前缀的所有环境变量

	return nil
}

// 读取配置渠道：本地文件读取（json/yaml/ini/...）或远程配置数据（etcd/consul/...）
// Viper 读取本地配置文件
func loadConfigFromFile(runtimeViper *viper.Viper, cfgPath, cfgName, cfgType string) error {
	var err error
	runtimeViper.AddConfigPath(cfgPath) // 设置配置文件读取路径，默认windows环境下为%GOPATH，linux环境下为$GOPATH
	runtimeViper.SetConfigName(cfgName) // 设置读取的配置文件名
	runtimeViper.SetConfigType(cfgType) // 设置配置文件类型

	if err = runtimeViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；可以读取当前目录,如果需要可以忽略

		} else {
			return err
		}
	}

	return nil
}

// Viper 读取远程配置系统
func loadConfigFromRemote(runtimeViper *viper.Viper, provider, endpoint, path string) error {
	var err error

	switch provider {
	case "etcd":
		err = runtimeViper.AddRemoteProvider(provider, endpoint, path)
		if err != nil {
			return err
		}

		// 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
		runtimeViper.SetConfigType("json")
		err = viper.ReadRemoteConfig()

		if err != nil {
			return err
		}

	case "consul":
		err = runtimeViper.AddRemoteProvider(provider, endpoint, path)
		if err != nil {
			return err
		}
		// 需要显示设置成json
		runtimeViper.SetConfigType("json")
		err = viper.ReadRemoteConfig()
		if err != nil {
			return err
		}
	default:
		return errors.New("Only Support etcd/consul K/V System. ")
	}

	return nil
}
