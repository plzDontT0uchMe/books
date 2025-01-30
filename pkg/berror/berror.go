package berror

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

// Система обработки ошибок

const (
	TooManyRequestsType  = "too_many_requests"
	AlreadyExistsType    = "already_exists"
	InternalType         = "internal"
	InvalidArgumentType  = "invalid_argument"
	NotFoundType         = "not_found"
	MethodNotAllowedType = "method_not_allowed"
	TimeoutType          = "timeout"
	UnauthorizedType     = "unauthorized"
	UnknownType          = "unknown"
	ValidationType       = "validation"
	ForbiddenType        = "forbidden"
)

type Base struct {
	Type    string            `json:"type"`
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

// getErr - функция для получения ошибки и кода ошибки.
func getErr(err error) (error, int) {
	if errors.Is(err, context.DeadlineExceeded) {
		return Timeout(), http.StatusRequestTimeout
	}

	switch err.(type) {
	case Base:
	default:
		slog.Error(err.Error())
		return Internal(), http.StatusInternalServerError
	}

	e := err.(Base)
	switch e.Type {
	case TooManyRequestsType:
		return err, http.StatusTooManyRequests
	case AlreadyExistsType:
		return err, http.StatusConflict
	case InvalidArgumentType:
		return err, http.StatusBadRequest
	case MethodNotAllowedType:
		return err, http.StatusMethodNotAllowed
	case NotFoundType:
		return err, http.StatusNotFound
	case TimeoutType:
		return err, http.StatusRequestTimeout
	case UnauthorizedType:
		return err, http.StatusUnauthorized
	case ValidationType:
		return err, http.StatusBadRequest
	case ForbiddenType:
		return err, http.StatusForbidden
	case UnknownType, InternalType:
		return Internal(), http.StatusInternalServerError
	default:
		return Internal(), http.StatusInternalServerError
	}
}

// HTTPError заносит в ответ код, соответствующий ошибке.
func HTTPError(w http.ResponseWriter, e error) {
	err, code := getErr(e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = fmt.Fprintln(w, err.Error())
	if err != nil {
		slog.Error(err.Error())
	}
}

func TooManyRequests() Base {
	return Base{
		Type:    TooManyRequestsType,
		Code:    TooManyRequestsType,
		Message: TooManyRequestsType,
	}
}

// AlreadyExists возвращает ошибку с типом already_exists (код 409).
func AlreadyExists() Base {
	return Base{
		Type:    AlreadyExistsType,
		Code:    AlreadyExistsType,
		Message: "already exists",
	}
}

// Internal возвращает ошибку с типом internal (код 500).
func Internal() Base {
	return Base{
		Type:    InternalType,
		Code:    InternalType,
		Message: "internal error",
	}
}

// InvalidArgument возвращает ошибку с типом invalid_argument (код 400).
func InvalidArgument() Base {
	return Base{
		Type:    InvalidArgumentType,
		Code:    InvalidArgumentType,
		Message: "invalid argument",
	}
}

// NotFound возвращает ошибку с типом not_found(код 404).
func NotFound() Base {
	return Base{
		Type:    NotFoundType,
		Code:    NotFoundType,
		Message: "not found",
	}
}

// MethodNotAllowed возвращает ошибку с типом method_not_allowed (код 405).
func MethodNotAllowed() Base {
	return Base{
		Type:    MethodNotAllowedType,
		Code:    MethodNotAllowedType,
		Message: "method not allowed",
	}
}

// Timeout возвращает ошибку с типом timeout (код 408).
func Timeout() Base {
	return Base{
		Type:    TimeoutType,
		Code:    TimeoutType,
		Message: TimeoutType,
	}
}

// Unauthorized возвращает ошибку с типом unauthorized (код 401).
func Unauthorized() Base {
	return Base{
		Type:    UnauthorizedType,
		Code:    UnauthorizedType,
		Message: UnauthorizedType,
	}
}

// Unknown возвращает ошибку с типом unknown (код 500).
func Unknown() Base {
	return Base{
		Type:    UnknownType,
		Code:    UnknownType,
		Message: UnknownType,
		Details: nil,
	}
}

// Validation возвращает ошибку с типом validation (код 400).
func Validation() Base {
	return Base{
		Type:    ValidationType,
		Code:    ValidationType,
		Message: "validation error",
	}
}

func Forbidden() Base {
	return Base{
		Type:    ForbiddenType,
		Code:    ForbiddenType,
		Message: ForbiddenType,
	}
}

// Error возвращает строку, содержащую JSON-представление ошибки.
func (e Base) Error() string {
	bytes, errMarshal := json.Marshal(e)
	if errMarshal != nil {
		return errMarshal.Error()
	}
	return string(bytes)
}

// Obj добавляет к коду и сообщению ошибки источник(объект) её возникновения.
func (e Base) Obj(object string) Base {
	e.Code += "_" + strings.ReplaceAll(object, " ", "_")
	e.Message += ": " + object
	return e
}

// Descr добавляет подробности к сообщнию об ошибке.
func (e Base) Descr(descr string) Base {
	e.Message += " " + descr
	return e
}

// DescrIsEmpty добавляет к сообщнию об ошибке 'is empty'.
func (e Base) DescrIsEmpty() Base {
	e.Message += " is empty"
	return e
}

// Msg устанавливает сообщение ошибки.
func (e Base) Msg(msg string) Base {
	e.Message += msg
	return e
}

// WithDetails добавляет детали к ошибке.
func (e Base) WithDetails(details map[string]string) Base {
	e.Details = details
	return e
}

// Level возвращает уровень логирования ошибки.
func (e Base) Level() slog.Level {
	switch e.Type {
	case AlreadyExistsType, InvalidArgumentType, MethodNotAllowedType, NotFoundType, UnauthorizedType, ValidationType, ForbiddenType, TooManyRequestsType:
		return slog.LevelError
	case TimeoutType, UnknownType, InternalType:
		return slog.LevelError + 4
	default:
		return slog.LevelError + 4
	}
}
