package test

import (
	"crypto/tls"
	"errors"
	"fmt"
	"gin/structs/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Error(c *gin.Context) {
	// var a int

	// a = 1
	// b, _ := strconv.Atoi(c.PostForm("b"))

	// d := a / b

	// fmt.Println(d)

	c.Error(errors.New("测试全局错误捕获"))

}

func Panic(c *gin.Context) {
	a := 1
	b, _ := strconv.Atoi(c.PostForm("b"))

	d := a / b

	fmt.Println(d)

	res := response.GetResponse()
	c.JSON(http.StatusOK, res)
}

// 发送邮件
func SendEmail(c *gin.Context) {
	m := gomail.NewMessage()
	m.SetHeader("From", "lufeijun_1234@126.com") //发送者
	m.SetHeader("To", "jipeng@zhufaner.com")     //接受者可以有多个
	body := "这是测试，这是演习"                          //发送内容
	m.SetHeader("Subject", "演习演习演习")             // 邮件标题
	m.SetBody("text/html", body)                 // 发送邮件内容
	d := gomail.NewDialer("smtp.126.com", 465, "lufeijun_1234@126.com", "JP418955279jp")
	//这里第一个参数为服务器地址，第二个为端口号，第三个为发送者邮箱号
	//第四个如果是qq邮箱为授权玛而其他邮箱是密码
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("***%s\n", err.Error())
	}

	res := response.GetResponse()
	c.JSON(http.StatusOK, res)
}

func Log(c *gin.Context) {

	// log1()

	log2()

	// log3()

	res := response.GetResponse()
	c.JSON(http.StatusOK, res)
	return
}

func log1() {
	logger, _ := zap.NewProduction()

	logger.Error(
		"错误信息",
		zap.String("name", "张三"),
	)

	logger.Info(
		"错误信息",
		zap.String("name", "张三"),
	)
}

func log2() {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	config := zap.Config{
		Level:            atom,                                                           // 日志级别
		Development:      true,                                                           // 开发模式，堆栈跟踪
		Encoding:         "json",                                                         // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                                  // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"},            // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"./logs/" + time.Now().Format("2006-01-02") + ".log"}, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}
	// 构建日志
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}
	logger.Info("log 初始化成功")
	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	logger.Error(
		"报错了",
		zap.String("name", "1232"),
	)

}

/**
* 分割函数好像有问题，
* 我用 ab -n 1000 -c 10 http://127.0.0.1:8008/api/test/log ，测试，超过 1 M 时，rename 函数调用失败，直接日志覆盖了
 */
func log3() {

	hook := lumberjack.Logger{
		Filename:   "./logs/" + time.Now().Format("2006-01-02") + ".log", // 日志文件路径
		MaxSize:    1,                                                    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,                                                    // 日志文件最多保存多少个备份
		MaxAge:     7,                                                    // 文件最多保存多少天
		Compress:   true,                                                 // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger := zap.New(core, caller, development, filed)
	logger.Info("log 初始化成功")
	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

}
