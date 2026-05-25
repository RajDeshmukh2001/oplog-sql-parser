package sql

import (
	"fmt"
	"sort"
	"strings"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateInsertSQL(oplog *model.Oplog) (string, error) {
	var columns []string
	var values []string

	for column := range oplog.O {
		columns = append(columns, column)
	}
	sort.Strings(columns)

	for _, column := range columns {
		value := oplog.O[column]

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
