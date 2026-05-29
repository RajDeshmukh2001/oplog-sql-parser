package sql

import (
	"testing"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func TestGenerateCreateSchemaSQL(t *testing.T) {
	oplog := &model.Oplog{
		Op: "i",
		Ns: "test.student",
		O: map[string]interface{}{
			"_id":           "d7cb2b45-c2b1-43ea-a049-619bbbf27037",
			"name":          "Raj Deshmukh",
			"roll_no":       24,
			"is_graduated":  false,
			"date_of_birth": "2001-06-30",
		},
	}

	query, err := GenerateCreateSchemaSQL(oplog)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "CREATE SCHEMA test;"

	if query != expected {
		t.Errorf("got %q, want %q", query, expected)
	}
}
