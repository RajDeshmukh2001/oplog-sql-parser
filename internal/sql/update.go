package sql

import (
	"fmt"
	"strings"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateUpdateSQL(oplog *model.Oplog) (string, error) {
	diff := oplog.O["diff"].(map[string]interface{})
	id := oplog.O2["_id"]

	var assignments []string

	if updateFields, ok := diff["u"].(map[string]interface{}); ok {
		columns := sortedKeys(updateFields)

		for _, column := range columns {
			value := updateFields[column]

			assignments = append(
				assignments,
				fmt.Sprintf("%s = %s", column, formatValue(value)),
			)
		}
	}

	if deleteFields, ok := diff["d"].(map[string]interface{}); ok {
		columns := sortedKeys(deleteFields)

		for _, column := range columns {
			assignments = append(
				assignments,
				fmt.Sprintf("%s = NULL", column),
			)
		}
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE _id = %s;",
		oplog.Ns,
		strings.Join(assignments, ", "),
		formatValue(id),
	)

	return query, nil
}
