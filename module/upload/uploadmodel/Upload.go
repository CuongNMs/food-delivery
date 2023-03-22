package uploadmodel

import (
	"errors"
)

type Upload struct {
}

var (
	ErrFileIsNotImage = errors.New("File is not image")
)
