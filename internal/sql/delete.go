package sql

import (
	"fmt"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateDeleteSQL(oplog *model.Oplog) (string, error) {
	id := oplog.O["_id"]

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE _id = %s;",
		oplog.Ns,
		formatValue(id),
	)

	return query, nil
}
