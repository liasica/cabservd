package logger

import (
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    log "github.com/sirupsen/logrus"
    "path"
    "time"
)

type Config struct {
    Color  bool
    Level  string
    Age    int
    Json   bool
    Caller bool
}

func LoadWithConfig(cfg Config) {
    rotateOptions := []rotatelogs.Option{
        rotatelogs.WithRotationTime(time.Hour * 24),
    }
    rotateOptions = append(rotateOptions, rotatelogs.WithMaxAge(time.Duration(cfg.Age)*time.Hour))
    w, err := rotatelogs.New(path.Join("runtime/logs", "%Y-%m-%d.log"), rotateOptions...)
    if err != nil {
        log.Errorf("rotatelogs initial failed: %v", err)
        panic(err)
    }

    consoleFormatter := LogFormat{EnableColor: cfg.Color, Console: true, Caller: cfg.Caller}
    fileFormatter := LogFormat{EnableColor: false, SaveJson: cfg.Json, Caller: cfg.Caller}
    log.AddHook(NewLocalHook(w, consoleFormatter, fileFormatter, GetLogLevel(cfg.Level)...))
    log.SetReportCaller(true)
}
