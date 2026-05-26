package sql

import (
	"fmt"
	"strings"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateUpdateSQL(oplog *model.Oplog) (string, error) {
	diff := oplog.O["diff"].(map[string]interface{})
	updates := diff["u"].(map[string]interface{})
	id := oplog.O2["_id"]

	var assignments []string
	columns := sortedKeys(updates)

	for _, column := range columns {
		value := updates[column]

		assignments = append(
			assignments,
			fmt.Sprintf("%s = %s", column, formatValue(value)),
		)
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE _id = %s;",
		oplog.Ns,
		strings.Join(assignments, ", "),
		formatValue(id),
	)

	return query, nil
}
