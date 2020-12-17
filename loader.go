package goinfras

import (
	"errors"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"strings"
)

// Viper读取环境变量
/**
 * @Description: Viper读取环境变量
 * @param viperCfg 			Viper实例
 * @param Prefix			指定需读取的环境变量前缀
 * @param Keys				绑定特定环境变量
 * @param AllowEmptyEnv		是否允许读取空值的环境变量
 * @param AutomaticEnv		是否载入所有环境变量，如设置Prefix，则只筛选有前缀的载入
 * @param Replacer			键名字符替换器，用于排除一些前缀符合但不需要的环境变量
 * @return error
 */
func LoadViperConfigFromEnv(viperCfg *viper.Viper, envPrefix string, envKeys []string, allowEmptyEnv, automaticEnv bool, envKeyReplacer *strings.Replacer) error {
	viperCfg.AllowEmptyEnv(allowEmptyEnv) // 是否允许环境变量为空值,默认为false
	if envKeyReplacer != nil {
		viperCfg.SetEnvKeyReplacer(envKeyReplacer) // 替换一些不需要的变量
	}
	if envPrefix != "" {
		viperCfg.SetEnvPrefix(envPrefix) // 添加需加载系统环境变量的前缀字符串
	}
	if len(envKeys) > 0 {
		err := viperCfg.BindEnv(envKeys...) // 绑定特定环境变量值到viper
		if err != nil {
			return err
		}
	}
	if automaticEnv {
		viperCfg.AutomaticEnv() // 自动载入所有环境变量，如设置SetEnvPrefix，则加载有特定前缀的所有环境变量
	}
	return nil
}

/**
 * @Description: Viper 读取配置文件
 * @param viperCfg 	Viper实例
 * @param configFilePath 配置文件路径及文件名
 * @return error
 */
func LoadViperConfigFromFile(viperCfg *viper.Viper, configFilePath string) error {
	var (
		err       error
		cPath     string // 配置文件路径部分
		cFileName string // 配置文件文件全名部分
		cName     string // 配置文件名称部分（不包含扩展名）
		cExt      string // 配置文件扩展名部分
	)
	cPath, cFileName = filepath.Split(configFilePath)    // 分离路径和文件全名
	cExt = path.Ext(cFileName)                           // 获取包含点符的扩展名
	cExt = cExt[1:]                                      // 去掉点符
	cName = cFileName[0 : len(cFileName)-len(cFileName)] // 获取分离拓展名的文件名
	viperCfg.AddConfigPath(cPath)                        // 设置配置文件读取路径，默认windows环境下为%GOPATH，linux环境下为$GOPATH
	viperCfg.SetConfigName(cName)                        // 设置读取的配置文件名
	viperCfg.SetConfigType(cExt)                         // 设置配置文件类型
	if err = viperCfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；可以读取当前目录,如果需要可以忽略

		} else {
			return err
		}
	}

	return nil
}

/**
 * @Description:  Viper 读取远程配置系统
 * @param viperCfg  Viper实例
 * @param remoteProvider  远程配置提供者etcd/consult
 * @param remoteEndpoint  远程配置主机节点
 * @param remotePath	  远程配置键节点
 * @param remoteType	  因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
 * @return error
 */
func LoadViperConfigFromRemote(viperCfg *viper.Viper, remoteProvider, remoteEndpoint, remotePath, remoteType string) error {
	var err error
	switch remoteProvider {
	case "etcd", "consul":
		if err = viperCfg.AddRemoteProvider(remoteProvider, remoteEndpoint, remotePath); err != nil {
			return err
		}
		viperCfg.SetConfigType(remoteType)
		if err = viper.ReadRemoteConfig(); err != nil {
			return err
		}
	default:
		return errors.New("Only Support etcd/consul K/V System. ")
	}

	return nil
}
