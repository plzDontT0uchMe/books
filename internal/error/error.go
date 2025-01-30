package error

import "backend/go/books/pkg/berror"

func Equal(err error, target berror.Base) bool {
	if e, ok := err.(berror.Base); ok {
		return e.Type == target.Type && e.Code == target.Code && e.Message == target.Message
	}
	return false
}
