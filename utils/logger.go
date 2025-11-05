package utils

import (
	"io" // 导入 io 包
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {
	logFile, err := os.OpenFile(DMPRuntimeLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// 使用 io.MultiWriter 将文件和标准输出 (os.Stdout) 合并
	multiWriter := io.MultiWriter(logFile, os.Stdout)

	// 创建一个替换时间的函数
	customTimeFormat := "2006-01-02 15:04:05"
	replaceTime := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			t := a.Value.Time()
			a.Value = slog.StringValue(t.Format(customTimeFormat))
		}
		return a
	}

	Logger = slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{ // 将 multiWriter 传入
		AddSource:   true,           // 记录错误位置
		Level:       slog.LevelInfo, // 设置日志级别
		ReplaceAttr: replaceTime,
	}))

	Logger.Info("Logger initialized. Outputting to both file and console.")
}
