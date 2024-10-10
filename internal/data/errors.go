package data

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("no matching record found")

func processGormError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return ErrNoRecord
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrNoRecord
	default:
		return err
	}
}
