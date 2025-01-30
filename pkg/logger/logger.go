package logger

import (
	"log/slog"
	"os"
	"slices"

	c "backend/go/books/internal/constants"
)

// Логгер приложения и его настройки

// hideValueInSecretGroup скрывает значение атрибута в группе "secret", заменяя значение на "...", для того,
// чтобы не печатать в логи конфиденциальные данные.
func hideValueInSecretGroup(groups []string, attribute slog.Attr) slog.Attr {
	if slices.Contains(groups, "secret") {
		return slog.Attr{
			Key:   attribute.Key,
			Value: slog.StringValue("..."),
		}
	}
	return attribute
}

// MustLoad загружает логгер в зависимости от окружения, что позволяет запускать приложение с разным уровнем логирования
// Также можно добавить isDebug, чтобы включать дополнительные логи в зависимости от значения, не уверен, что это хорошая идея.
func MustLoad(env string) error {
	logOpts := slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: hideValueInSecretGroup,
	}

	switch env {
	case c.EnvDevelopment:
		logOpts.Level = slog.LevelDebug
	case c.EnvStage:
		logOpts.Level = slog.LevelInfo
	case c.EnvProduction:
		logOpts.Level = slog.LevelError
	default:
		logOpts.Level = slog.LevelError
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &logOpts)))

	return nil
}
