package validator

import (
	"errors"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func ValidateInsertOplog(oplog *model.Oplog) error {
	if oplog == nil {
		return errors.New("oplog cannot be nil")
	}

	if oplog.Op != "i" {
		return errors.New("invalid operation type")
	}

	if oplog.Ns == "" {
		return errors.New("namespace is required")
	}

	if len(oplog.O) == 0 {
		return errors.New("document data is required")
	}

	return nil
}
