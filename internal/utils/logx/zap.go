package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var sugar *zap.SugaredLogger

func init() {
	NewLogger()
	//log.Println("zap log init success")
}

func NewLogger() {
	core := newCore(zap.DebugLevel)
	caller := zap.AddCaller()
	// 向上跳一层，否则日志中caller显示不正确
	callerSkip := zap.AddCallerSkip(1)
	// 构造日志
	sugar = zap.New(core, caller, callerSkip).Sugar()
	return
}

func newCore(level zapcore.Level) zapcore.Core {

	//日志文件路径配置
	hook := lumberjack.Logger{
		Filename: ".cache/local.log", // 日志文件路径
		MaxSize:  1024,               // 每个日志文件保存的最大尺寸 单位：M
		Compress: true,               // 是否压缩
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                        // json时时间键
		LevelKey:       "level",                       // json时日志等级键
		NameKey:        "name",                        // json时日志记录器名
		CallerKey:      "caller",                      // json时日志文件信息键
		MessageKey:     "message",                     // json时日志消息键
		StacktraceKey:  "stackTrace",                  // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,     // 友好日志换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 友好日志等级名大小写（info INFO）
		EncodeTime:     timeEncoder,                   // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 日志文件信息（包/文件.go:行号）
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台(os.Stdout)和文件(&hook)
		atomicLevel, // 日志级别
	)
}

/**
 * @Description: 格式化时间
 * @param t
 * @param enc
 */
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
