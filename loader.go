package goinfras

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

// 环境变量配置值参数
type EnvConfigArgs struct {
	Prefix        string            // 环境变量前缀
	AllowEmptyEnv bool              // 是否允许读取空值的环境变量
	AutomaticEnv  bool              // 是否载入所有环境变量，如设置Prefix，则只筛选有前缀的载入
	Keys          []string          // 绑定特定环境变量
	Replacer      *strings.Replacer // 键名字符替换器，用于排除一些前缀符合但不需要的环境变量
}

// 文件配置值参数
type FileConfigArgs struct {
	Path string // 配置文件路径
	Name string // 配置文件名
	Type string // 配置文件类型
}

// 远程配置值参数
type RemoteConfigArgs struct {
	Provider string // 远程配置服务提供者，only etcd、consult
	Endpoint string // 配置服务节点
	Path     string // 键路径
	Type     string // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
}

func ViperLoader(envCfgArgs *EnvConfigArgs, fileCfgARgs *FileConfigArgs, remoteCfgArgs *RemoteConfigArgs) (*viper.Viper, error) {
	var err error
	var viperCfg *viper.Viper

	viperCfg = viper.New()

	// 1. 从环境变量导入配置项
	if envCfgArgs != nil {
		if err = loadConfigFromEnv(viperCfg, envCfgArgs); err != nil {
			return nil, err
		}
	}

	// 2. 从配置文件导入配置项
	if fileCfgARgs != nil {
		if err = loadConfigFromFile(viperCfg, fileCfgARgs); err != nil {
			return nil, err
		}
	}

	// 3. 从远程配置系统导入配置项
	if remoteCfgArgs != nil {
		if err = loadConfigFromRemote(viperCfg, remoteCfgArgs); err != nil {
			return nil, err
		}
	}

	return viperCfg, nil
}

// Viper读取环境变量
func loadConfigFromEnv(viperCfg *viper.Viper, cfg *EnvConfigArgs) error {
	viperCfg.AllowEmptyEnv(cfg.AllowEmptyEnv) // 是否允许环境变量为空值,默认为false
	if cfg.Replacer != nil {
		viperCfg.SetEnvKeyReplacer(cfg.Replacer) // 替换一些不需要的变量
	}
	if cfg.Prefix != "" {
		viperCfg.SetEnvPrefix(cfg.Prefix) // 添加需加载系统环境变量的前缀字符串
	}
	if len(cfg.Keys) > 0 {
		err := viperCfg.BindEnv(cfg.Keys...) // 绑定特定环境变量值到viper
		if err != nil {
			return err
		}
	}
	if cfg.AutomaticEnv {
		viperCfg.AutomaticEnv() // 自动载入所有环境变量，如设置SetEnvPrefix，则加载有特定前缀的所有环境变量
	}
	return nil
}

// 读取配置渠道：本地文件读取（json/yaml/ini/...）或远程配置数据（etcd/consul/...）
// Viper 读取本地配置文件
func loadConfigFromFile(viperCfg *viper.Viper, cfg *FileConfigArgs) error {
	var err error
	viperCfg.AddConfigPath(cfg.Path) // 设置配置文件读取路径，默认windows环境下为%GOPATH，linux环境下为$GOPATH
	viperCfg.SetConfigName(cfg.Name) // 设置读取的配置文件名
	viperCfg.SetConfigType(cfg.Type) // 设置配置文件类型
	if err = viperCfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；可以读取当前目录,如果需要可以忽略

		} else {
			return err
		}
	}

	return nil
}

// Viper 读取远程配置系统
func loadConfigFromRemote(viperCfg *viper.Viper, cfg *RemoteConfigArgs) error {
	var err error
	switch cfg.Provider {
	case "etcd", "consul":
		if err = viperCfg.AddRemoteProvider(cfg.Provider, cfg.Endpoint, cfg.Path); err != nil {
			return err
		}

		// 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
		viperCfg.SetConfigType(cfg.Type)

		if err = viper.ReadRemoteConfig(); err != nil {
			return err
		}
	default:
		return errors.New("Only Support etcd/consul K/V System. ")
	}

	return nil
}
