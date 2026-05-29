package sql

import (
	"fmt"
	"strings"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateCreateTableSQL(oplog *model.Oplog) (string, error) {
	columnNames := sortedKeys(oplog.O)
	columns := make([]string, 0, len(columnNames))

	for _, column := range columnNames {
		columnType := getColumnType(oplog.O[column])

		if column == "_id" {
			columns = append(columns, fmt.Sprintf("%s %s PRIMARY KEY", column, columnType))
		} else {
			columns = append(columns, fmt.Sprintf("%s %s", column, columnType))
		}
	}

	query := fmt.Sprintf(
		"CREATE TABLE %s (%s);",
		oplog.Ns,
		strings.Join(columns, ", "),
	)

	return query, nil
}
