package sql

import (
	"fmt"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateUpdateSQL(oplog *model.Oplog) (string, error) {
	diff := oplog.O["diff"].(map[string]interface{})
	updates := diff["u"].(map[string]interface{})
	id := oplog.O2["_id"]

	var column string
	var value interface{}

	for key, val := range updates {
		column = key
		value = val
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s = %s WHERE _id = %s;",
		oplog.Ns,
		column,
		formatValue(value),
		formatValue(id),
	)

	return query, nil
}
