package validator

import (
	"errors"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func ValidateUpdateOplog(oplog *model.Oplog) error {
	if oplog == nil {
		return errors.New("oplog cannot be nil")
	}

	if oplog.Op != "u" {
		return errors.New("invalid operation type")
	}

	if oplog.Ns == "" {
		return errors.New("namespace is required")
	}

	if oplog.O == nil {
		return errors.New("update operation data is required")
	}

	diff, ok := oplog.O["diff"].(map[string]interface{})
	if !ok {
		return errors.New("diff data is required")
	}

	_, hasUpdates := diff["u"]
	_, hasDeletes := diff["d"]

	if !hasUpdates && !hasDeletes {
		return errors.New("no update operations found")
	}

	if oplog.O2 == nil {
		return errors.New("update criteria is required")
	}

	if _, ok := oplog.O2["_id"]; !ok {
		return errors.New("_id is required in update criteria")
	}

	return nil
}
