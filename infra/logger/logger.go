package logger

import (
	"io"
	"runtime/debug"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

var log zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel := config.Envs.Logger.LogLevel

		var logWriter io.Writer = zerolog.ConsoleWriter{
			Out:        config.Envs.Logger.Output,
			TimeFormat: time.RFC3339,
			FieldsExclude: []string{
				"user_agent",
				"go_version",
				"git_revision",
			},
		}

		fileLogger := &lumberjack.Logger{
			Filename:   config.Envs.Logger.Filename,
			MaxSize:    config.Envs.Logger.MaxSize,
			MaxBackups: config.Envs.Logger.MaxBackups,
			MaxAge:     config.Envs.Logger.MaxAge,
			Compress:   config.Envs.Logger.Compress,
		}

		output := zerolog.MultiLevelWriter(fileLogger, logWriter)

		var gitRevision string

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
	})

	return log
}
