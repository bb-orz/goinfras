package config

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/constant"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Parse() *AppConfig {
	fmt.Println("Config Init...")
	// 获取配置文件目录
	confPath := getConfigPath()

	// 读取yaml配置文件并解析
	var (
		err           error
		baseConfFile  []byte
		logConfFile   []byte
		mysqlConfFile []byte
		redisConfFile []byte
		mongoConfFile []byte
		mqConfFile    []byte
		corsConfFile  []byte
		ossConfFile   []byte
		etcdConfFile  []byte

		baseConf  *Base
		logConf   *Log
		mysqlConf *Mysql
		redisConf *Redis
		mongoConf *Mongodb
		mqConf    *Mq
		corsConf  *Cors
		ossConf   *Oss
		etcdConf  *Etcd
	)

	baseConfFile, err = ioutil.ReadFile(confPath + "/base.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(baseConfFile, &baseConf)
	infras.FailHandler(err)

	logConfFile, err = ioutil.ReadFile(confPath + "/log.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(logConfFile, &logConf)
	infras.FailHandler(err)

	mysqlConfFile, err = ioutil.ReadFile(confPath + "/mysql.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(mysqlConfFile, &mysqlConf)
	infras.FailHandler(err)

	redisConfFile, err = ioutil.ReadFile(confPath + "/redis.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(redisConfFile, &redisConf)
	infras.FailHandler(err)

	mongoConfFile, err = ioutil.ReadFile(confPath + "/mongodb.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(mongoConfFile, &mongoConf)
	infras.FailHandler(err)

	mqConfFile, err = ioutil.ReadFile(confPath + "/mq.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(mqConfFile, &mqConf)
	infras.FailHandler(err)

	corsConfFile, err = ioutil.ReadFile(confPath + "/cors.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(corsConfFile, &corsConf)
	infras.FailHandler(err)

	ossConfFile, err = ioutil.ReadFile(confPath + "/oss.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(ossConfFile, &ossConf)
	infras.FailHandler(err)

	etcdConfFile, err = ioutil.ReadFile(confPath + "/etcd.yaml")
	infras.FailHandler(err)
	err = yaml.Unmarshal(etcdConfFile, &etcdConf)
	infras.FailHandler(err)

	return &AppConfig{
		BaseConf:  baseConf,
		LogConf:   logConf,
		MysqlConf: mysqlConf,
		RedisConf: redisConf,
		MongoConf: mongoConf,
		MqConf:    mqConf,
		CorsConf:  corsConf,
		OssConf:   ossConf,
		EtcdConf:  etcdConf,
	}

}

// 启动时根据系统环境变量获取配置文件目录参数,如没有则在当前目录的config文件夹查找
func getConfigPath() string {
	// 先从启动命令行参数获取配置根目录信息
	var configRootPath string
	if len(os.Args) > 1 {
		configRootPath = os.Args[1]
	}
	if configRootPath == "" {
		configRootPath = "./build/config"
	}

	// 再从环境变量获取具体环境的信息 (dev|product|testing)
	var currentEnv string
	currentEnv = os.Getenv(constant.OsEnvVarName)
	if currentEnv == "" {
		currentEnv = constant.DefaultEnv
	}
	confPath := configRootPath + "/" + currentEnv + "/"

	dirs, err := ioutil.ReadDir(confPath)
	if os.IsNotExist(err) {
		infras.ErrorHandler(err)
	}
	if len(dirs) == 0 {
		panic("Not Found Config Files...")
	}

	return confPath
}
