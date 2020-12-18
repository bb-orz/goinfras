package goinfras

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 记录位置命名常量
const (
	ApplicationPosition = "Application"
	StarterPosition     = "Starter"
)

// 启动步骤常量
const (
	StepInit  = "Init"
	StepSetup = "Setup"
	StepStart = "Start"
	StepCheck = "Check"
	StepStop  = "Stop"
)

// 日志等级命名常量
const (
	DebugLevel   = "Debug"
	InfoLevel    = "Info"
	WarningLevel = "Warning"
	ErrorLevel   = "Error"
	FailLevel    = "Fail"
)

const (
	blue = "\033[97;44m"
	cyan = "\033[97;46m"

	whitef   = "\033[37m"
	yellowf  = "\033[33m"
	bluef    = "\033[34m"
	greenf   = "\033[32m"
	redf     = "\033[31m"
	magentaf = "\033[35m"

	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[1;97;41m"
	magenta = "\033[1;97;45m"

	reset = "\033[0m"
)

// 可定义多个输出
type StarterLoggerOutput struct {
	Formatter LogFormatter // 格式转化器
	Writers   io.Writer    // 输出
}

// 格式转化签名函数
type LogFormatter func(params LogFormatterParams) string

// 格式化输出参数
type LogFormatterParams struct {
	Position  string    // 日志记录位置
	Name      string    // 启动器名称
	Step      string    // 启动器步骤
	LogLevel  string    // 记录日志级别
	TimeStamp time.Time // 记录时间戳
	Message   string    // 记录信息
}

// 日志输出位置颜色标示
func (p *LogFormatterParams) LogPositionColor() string {
	switch p.Position {
	case ApplicationPosition:
		return blue
	case StarterPosition:
		return cyan
	default:
		return cyan
	}
}

// 启动器步骤颜色标示
func (p *LogFormatterParams) LogStepColor() string {
	switch p.Step {
	case StepInit:
		return white
	case StepSetup:
		return yellow
	case StepStart:
		return blue
	case StepCheck:
		return green
	default:
		return cyan
	}
}

// 每种错误级别输出不同的颜色
func (p *LogFormatterParams) LogLevelColor() string {
	switch p.LogLevel {
	case DebugLevel:
		return greenf
	case InfoLevel:
		return whitef
	case WarningLevel:
		return yellowf
	case ErrorLevel:
		return redf
	case FailLevel:
		return magentaf
	default:
		return white
	}
}

// 颜色重置
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// 启动日志默认终端颜色输出格式
var defaultLogFormatter = func(param LogFormatterParams) string {
	var positionColor, stepColor, logLevelColor, resetColor string

	positionColor = param.LogPositionColor()
	stepColor = param.LogStepColor()
	logLevelColor = param.LogLevelColor()
	resetColor = param.ResetColor()

	return fmt.Sprintf("[%s %s %s %s] | %s %s %s | %v | %s [%s] >>>>>> %s %s \n",
		positionColor, param.Name, param.Position, resetColor,
		stepColor, param.Step, resetColor,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		logLevelColor, param.LogLevel, param.Message, resetColor,
	)
}

// 启动日志文件输出格式
var commonWriterFormatter = func(param LogFormatterParams) string {
	return fmt.Sprintf("[%s %s] | %s | %v | [%s] >>>>>>  %s",
		param.Name,
		param.Position,
		param.Step,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.LogLevel, param.Message,
	)
}

type IStarterLogger interface {
	ADebug(msg string)
	AInfo(msg string)
	AWarning(msg string)
	AError(err error)
	AFail(err error)
	SDebug(name, step, msg string)
	SInfo(name, step, msg string)
	SWarning(name, step, msg string)
	SError(name, step string, err error)
	SFail(name, step string, err error)
}

// 启动器日志记录器
type StarterLogger struct {
	Outputs []*StarterLoggerOutput
}

func (l *StarterLogger) ADebug(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  DebugLevel,
			TimeStamp: time.Now(),
			Message:   msg,
			// 是否增加caller
		}))
	}
}

func (l *StarterLogger) AInfo(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  InfoLevel,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) AWarning(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  WarningLevel,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) AError(err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  ErrorLevel,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}
func (l *StarterLogger) AFail(err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  FailLevel,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}

func (l *StarterLogger) SDebug(name, step, msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			Name:      name,
			LogLevel:  DebugLevel,
			Step:      step,
			TimeStamp: time.Now(),
			Message:   msg,
			// 可增加caller
		}))
	}
}

func (l *StarterLogger) SInfo(name, step, msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			Name:      name,
			LogLevel:  InfoLevel,
			Step:      step,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) SWarning(name, step, msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			Name:      name,
			LogLevel:  WarningLevel,
			Step:      step,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) SError(name, step string, err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			Name:      name,
			LogLevel:  ErrorLevel,
			Step:      step,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}
func (l *StarterLogger) SFail(name, step string, err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			Name:      name,
			LogLevel:  FailLevel,
			Step:      step,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}

// 标准颜色输出日志记录器
type CommandLineStarterLogger struct {
	StarterLogger
}

// 针对终端输出的默认启动日志记录器
func NewCommandLineStarterLogger() *CommandLineStarterLogger {
	logger := new(CommandLineStarterLogger)
	output := new(StarterLoggerOutput)
	output.Formatter = defaultLogFormatter
	output.Writers = os.Stdout
	logger.Outputs = make([]*StarterLoggerOutput, 0)
	logger.Outputs = append(logger.Outputs, output)
	return logger
}

// 通用启动日志记录器（输出到终端和启动日志文件）
type CommonStarterLogger struct {
	StarterLogger
}

func NewStarterLoggerWithWriters(writers ...io.Writer) *CommonStarterLogger {
	logger := new(CommonStarterLogger)
	logger.Outputs = make([]*StarterLoggerOutput, 0)

	// 标准输出
	outputStd := new(StarterLoggerOutput)
	outputStd.Formatter = defaultLogFormatter
	outputStd.Writers = os.Stdout
	logger.Outputs = append(logger.Outputs, outputStd)

	for _, w := range writers {
		outputFile := new(StarterLoggerOutput)
		outputFile.Formatter = commonWriterFormatter
		outputFile.Writers = w
		logger.Outputs = append(logger.Outputs, outputFile)
	}

	return logger
}
