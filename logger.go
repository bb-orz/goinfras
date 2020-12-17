package goinfras

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 记录位置常量
const (
	ApplicationPosition = "Application"
	StarterPosition     = "Starter"
)

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

	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	magenta = "\033[97;45m"

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
	LogLevel  string    // 记录日志级别
	TimeStamp time.Time // 记录时间戳
	Message   string    // 记录信息
}

// 日志输出位置标示
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

// 每种错误级别输出不同的颜色
func (p *LogFormatterParams) LogLevelColor() string {
	switch p.LogLevel {
	case DebugLevel:
		return green
	case InfoLevel:
		return white
	case WarningLevel:
		return yellow
	case ErrorLevel:
		return red
	case FailLevel:
		return magenta
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
	var positionColor, logLevelColor, resetColor string

	positionColor = param.LogPositionColor()
	logLevelColor = param.LogLevelColor()
	resetColor = param.ResetColor()

	return fmt.Sprintf("[%s %s %s] | %v | %s [%s] >>>>>>  %s %s",
		positionColor, param.Position, resetColor,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		logLevelColor, param.LogLevel, param.Message, resetColor,
	)
}

// 启动日志文件输出格式
var fileLogFormatter = func(param LogFormatterParams) string {
	return fmt.Sprintf("[%s ] | %v | [%s] >>>>>>  %s",
		param.Position,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.LogLevel, param.Message,
	)
}

// 启动器日志记录器
type StarterLogger struct {
	Outputs []*StarterLoggerOutput
}

func (l *StarterLogger) ApplicationDebug(msg string) {
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

func (l *StarterLogger) ApplicationInfo(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  InfoLevel,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) ApplicationWarning(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  WarningLevel,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) ApplicationError(err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  ErrorLevel,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}
func (l *StarterLogger) ApplicationFail(err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  ApplicationPosition,
			LogLevel:  FailLevel,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}

func (l *StarterLogger) StarterDebug(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			LogLevel:  DebugLevel,
			TimeStamp: time.Now(),
			Message:   msg,
			// 是否增加caller
		}))
	}
}

func (l *StarterLogger) StarterInfo(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			LogLevel:  InfoLevel,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) StarterWarning(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			LogLevel:  WarningLevel,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) StarterError(err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			LogLevel:  ErrorLevel,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}
func (l *StarterLogger) StarterFail(err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			Position:  StarterPosition,
			LogLevel:  FailLevel,
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

func NewCommandStarterLogger(starterLoggerFile string) (*CommonStarterLogger, error) {
	var err error
	logger := new(CommonStarterLogger)
	logger.Outputs = make([]*StarterLoggerOutput, 0)

	// 标准输出
	outputStd := new(StarterLoggerOutput)
	outputStd.Formatter = defaultLogFormatter
	outputStd.Writers = os.Stdout
	logger.Outputs = append(logger.Outputs, outputStd)

	outputFile := new(StarterLoggerOutput)
	outputFile.Formatter = fileLogFormatter
	// TODO 创建输出文件
	outputFile.Writers, err = os.Create(starterLoggerFile)
	if err != nil {
		return nil, err
	}
	logger.Outputs = append(logger.Outputs, outputFile)
	return logger, nil
}
