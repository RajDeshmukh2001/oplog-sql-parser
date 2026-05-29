package sql

import (
	"fmt"
	"strings"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func GenerateCreateSchemaSQL(oplog *model.Oplog) (string, error) {
	schemaName := strings.Split(oplog.Ns, ".")[0]

	query := fmt.Sprintf(
		"CREATE SCHEMA %s;",
		schemaName,
	)

	return query, nil
}
