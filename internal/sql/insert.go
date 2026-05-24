package sql

import (
	"fmt"
	"strings"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateInsertSQL(oplog *model.Oplog) (string, error) {
	var columns []string
	var values []string

	for column, value := range oplog.O {
		columns = append(columns, column)

		switch v := value.(type) {
		case string:
			values = append(values, fmt.Sprintf("'%s'", v))
		default:
			values = append(values, fmt.Sprintf("%v", v))
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		oplog.Ns,
		strings.Join(columns, ", "),
		strings.Join(values, ", "),
	)

	return query, nil
}
