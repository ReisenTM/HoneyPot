package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"image_server/internal/global"
	"os"
	"path"
	"sync"
	"time"
)

type MyLog struct {
}

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		var appName = entry.Data["appName"]
		if appName == nil {
			appName = global.Config.Logger.AppName
		}
		fmt.Fprintf(b, "%s=>[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", appName, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}

// MyHook 是自定义 logrus Hook，用于按级别将日志写入不同的文件。
// 实现了 logrus.Hook 接口：
//
//	type Hook interface {
//	    Levels() []logrus.Level       // 返回本 Hook 关心的日志级别
//	    Fire(*logrus.Entry) error     // 针对每条日志执行的处理逻辑
//	}
type MyHook struct {
	file     *os.File
	errFile  *os.File
	fileDate string
	logPath  string
	mu       sync.Mutex
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {
	hook.mu.Lock()
	defer hook.mu.Unlock()

	timer := entry.Time.Format("2006-01-02")
	line, err := entry.String()
	if err != nil {
		return fmt.Errorf("failed to format log entry: %v", err)
	}

	if hook.fileDate != timer {
		if err := hook.rotateFiles(timer); err != nil {
			return err
		}
	}

	if _, err := hook.file.Write([]byte(line)); err != nil {
		return fmt.Errorf("failed to write to log file: %v", err)
	}

	if entry.Level <= logrus.ErrorLevel {
		if _, err := hook.errFile.Write([]byte(line)); err != nil {
			return fmt.Errorf("failed to write to error log file: %v", err)
		}
	}

	return nil
}

// rotateFiles 日志轮换
func (hook *MyHook) rotateFiles(timer string) error {
	if hook.file != nil {
		if err := hook.file.Close(); err != nil {
			return fmt.Errorf("failed to close log file: %v", err)
		}
	}
	if hook.errFile != nil {
		if err := hook.errFile.Close(); err != nil {
			return fmt.Errorf("failed to close error log file: %v", err)
		}
	}

	dirName := fmt.Sprintf("%s/%s", hook.logPath, timer)
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	infoFilename := fmt.Sprintf("%s/info.log", dirName)
	errFilename := fmt.Sprintf("%s/err.log", dirName)

	var err error
	hook.file, err = os.OpenFile(infoFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	hook.errFile, err = os.OpenFile(errFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open error log file: %v", err)
	}

	hook.fileDate = timer
	return nil
}

// Levels 哪些级别的日志能走 Fire 方法
func (hook *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func GetLogger() *logrus.Entry {
	logger := logrus.New()
	l := global.Config.Logger
	level, err := logrus.ParseLevel(l.Level)
	if err != nil {
		logrus.Warnf("日志级别配置错误 自动修改为 info")
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	logger.AddHook(&MyHook{logPath: "logs"})
	if l.Format == "json" {
		//因为elk，默认json输出
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.DateTime,
		})
	} else {
		//如果没设为json模式就是text输出
		logger.SetFormatter(&MyLog{})
	}
	//显示调用日志的代码位置
	logger.SetReportCaller(true)
	return logger.WithField("appName", l.AppName)
}

func SetLogDefault() {
	l := global.Config.Logger
	logrus.SetFormatter(&MyLog{})
	logrus.SetReportCaller(true)
	logrus.WithField("appName", l.AppName)
}
