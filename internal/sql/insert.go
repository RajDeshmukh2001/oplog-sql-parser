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

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
