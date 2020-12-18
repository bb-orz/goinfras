package goinfras

import (
	"fmt"
	"github.com/spf13/viper"
	"runtime"
	"strconv"
)

const (
	KeyConfig = "_vpcfg"
	KeyLogger = "_logger"
)

// 资源启动器上下文，用来在服务资源初始化、安装、启动和停止的生命周期中变量和对象的传递
type StarterContext map[string]interface{}

// 创建一个默认最少配置启动器上下文
func CreateDefaultStarterContext(vpcfg *viper.Viper, logger IStarterLogger) *StarterContext {
	sctx := &StarterContext{}
	sctx.SetConfigs(vpcfg)
	sctx.SetLogger(logger)
	return sctx
}

func (s StarterContext) Configs() *viper.Viper {
	p := s[KeyConfig]
	if p == nil {
		panic("配置还没有被初始化")
	}
	return p.(*viper.Viper)
}
func (s StarterContext) SetConfigs(vpcfg *viper.Viper) {
	s[KeyConfig] = vpcfg
}

func (s StarterContext) Logger() IStarterLogger {
	p := s[KeyLogger]
	if p == nil {
		panic("日志记录器还没有被初始化")
	}
	return p.(IStarterLogger)
}
func (s StarterContext) SetLogger(logger IStarterLogger) {
	s[KeyLogger] = logger
}

// 有错误则记录启动器警告日志
func (s StarterContext) PassWarning(name, step string, err error) {
	if err != nil {
		var path string
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + " : " + strconv.Itoa(line)
		}
		s.Logger().SWarning(name, step, fmt.Sprintf("Warning: %s >>> [ %s ] \n", err.Error(), path))
	}
}

// err == nil 返回 true，否则记录启动器错误日志并返回false
func (s StarterContext) PassError(name, step string, err error) bool {
	if err == nil {
		return true
	} else {
		var path string
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + " : " + strconv.Itoa(line)
		}
		s.Logger().SError(name, step, fmt.Errorf("ERROR: %s >>> [ %s ] \n", err.Error(), path))
		return false
	}
}

// err == nil,返回true ;err != nil 致命错误处理，直接panic
func (s StarterContext) PassFail(name, step string, err error) bool {
	if err == nil {
		return true
	} else {
		var path string
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + " : " + strconv.Itoa(line)
		}
		s.Logger().SFail(name, step, fmt.Errorf("FAIL %s >>> [ %s ]", err.Error(), path))
		panic("")
	}
}
