package validator

import (
	"errors"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func ValidateDeleteOplog(oplog *model.Oplog) error {
	if oplog == nil {
		return errors.New("oplog cannot be nil")
	}

	if oplog.Op != "d" {
		return errors.New("invalid operation type")
	}

	if oplog.Ns == "" {
		return errors.New("namespace is required")
	}

	if oplog.O == nil {
		return errors.New("delete criteria is required")
	}

	if _, ok := oplog.O["_id"]; !ok {
		return errors.New("_id is required in delete criteria")
	}

	return nil
}
