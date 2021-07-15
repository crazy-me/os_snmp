package initialize

import (
	"github.com/crazy-me/os_snmp/utils"
	"github.com/crazy-me/os_snmp/utils/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func LogInit() {
	// 是否有日志目录
	isDir, _ := utils.DirExists(global.APP.Zap.LogPath)
	if !isDir {
		_ = os.Mkdir(global.APP.Zap.LogPath, os.ModePerm)
	}

	write := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, write, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	global.LOGGER = logger.Sugar()

}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   CustomLoggerName(),
		MaxSize:    global.APP.Zap.LogMaxSize,
		MaxBackups: global.APP.Zap.LogMaxBackups,
		MaxAge:     global.APP.Zap.LogMaxAge,
		Compress:   global.APP.Zap.LogCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// CustomTimeEncoder 定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.APP.Zap.LogPrefix + "2006-01-02/15:04:05.000"))
}

// CustomLoggerName 定义日志文件名称
func CustomLoggerName() string {
	return global.APP.Zap.LogPath + string(os.PathSeparator) + time.Now().Format("2006-01-02") + ".log"
}
