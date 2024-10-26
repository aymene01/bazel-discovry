package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger initializes and returns a configured zerolog Logger.
// It supports logging to console, file, or both, with configurable levels and formats.
func InitLogger(level, format, env, filePath string) (zerolog.Logger, error) {
	// Configure global zerolog time format to Unix time for consistency
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "timestamp"

	var writers []io.Writer

	// Console logging setup (useful for dev environments)
	if env == "development" {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		writers = append(writers, consoleWriter)
	} else if format == "console" {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		writers = append(writers, consoleWriter)
	}

	// File logging setup
	if filePath != "" {
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return zerolog.Logger{}, err
		}
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return zerolog.Logger{}, err
		}
		writers = append(writers, file)
	}

	// JSON output for structured logging in production
	if format == "json" && env == "production" {
		writers = append(writers, os.Stderr)
	}

	// Multi-output setup (console + file or JSON)
	log.Logger = log.Output(zerolog.MultiLevelWriter(writers...)).With().Caller().Timestamp().Logger()

	// Set log level based on input
	switch level {
	case "debug":
		log.Logger = log.Logger.Level(zerolog.DebugLevel)
	case "info":
		log.Logger = log.Logger.Level(zerolog.InfoLevel)
	case "warn":
		log.Logger = log.Logger.Level(zerolog.WarnLevel)
	case "error":
		log.Logger = log.Logger.Level(zerolog.ErrorLevel)
	default:
		log.Logger = log.Logger.Level(zerolog.InfoLevel) // Default to info level
	}

	return log.Logger, nil
}

// Usage example for initializing the logger:
// logger, err := logger.InitLogger("debug", "console", "development", "")
// if err != nil {
//     panic("Failed to initialize logger: " + err.Error())
// }
// logger.Info().Msg("Logger initialized!")
