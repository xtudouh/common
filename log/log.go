package log

/*
*   Example:
*    logger := log.NewLogger()
*    logger.Error("hello %v", "小黄")
 */

import (
    "github.com/siddontang/go-log/log"
    "os"
    "xtudouh/common/conf"
)

const (
    console = "console"
    file = "file"
)

var (
    logger   *log.Logger
    levelMap = map[string]int{
        "Trace": log.LevelTrace,
        "Debug": log.LevelDebug,
        "Info":  log.LevelInfo,
        "Warn":  log.LevelWarn,
        "Error": log.LevelError,
        "Fatal": log.LevelFatal,
    }
    mode, level, logFile string
    maxSize, backupCount int
)

func NewLogger() *log.Logger {
    if logger == nil {
        var (
            h log.Handler
            err error
        )
        if mode == file {
            h, err = log.NewRotatingFileHandler(logFile, maxSize, backupCount)
        } else {
            h, err = log.NewStreamHandler(os.Stdout)
        }
        if err != nil {
            panic(err)
        }
        logger = log.NewDefault(h)
        logger.SetLevel(levelMap[level])
    }
    return logger
}

func Init() {
    mode = conf.String("log", "MODE", "console")
    level = conf.String("log", "LEVEL", "Debug")
    logFile = conf.String("log", "FILE_NAME")
    maxSize = conf.Int("log", "MAX_SIZE", 1000000)
    backupCount = conf.Int("log", "BACKUP_COUNT", 100)
}
