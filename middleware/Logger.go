package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"snail/conf"
	"snail/util"
	"syscall"
	"time"
)

var LogClient = logrus.New()

func Logger() gin.HandlerFunc {

	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err!= nil{
		fmt.Println("err", err)
	}

	LogClient.Out = src
	LogClient.SetLevel(logrus.DebugLevel)

	// 错误重定向到日志文件
	err = syscall.Dup2(int(src.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		LogClient.Fatalf("Failed to redirect stderr to file: %v", err)
	}

	//记录函数名 todo 待完善
	LogClient.SetReportCaller(true)

	basePath := conf.ServerConf["logDir"]
	// 检查目录，不存在则创建
	if ok, _ := util.PathExists(basePath); !ok{
		_ = os.Mkdir(basePath, os.ModePerm)
	}
	apiLogPath := basePath + "/snail.log"
	logWriter, err := rotatelogs.New(
		apiLogPath+"%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	LogClient.AddHook(lfHook)

	return func (c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		LogClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}
