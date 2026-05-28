package sql

import (
	"fmt"
	"strings"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateInsertSQL(oplog *model.Oplog) (string, error) {
	var values []string
	columns := sortedKeys(oplog.O)

	for _, column := range columns {
		value := oplog.O[column]
		values = append(values, formatValue(value))
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		oplog.Ns,
		strings.Join(columns, ", "),
		strings.Join(values, ", "),
	)

	return query, nil
}
