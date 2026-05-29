package sql

import (
	"testing"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func TestGenerateCreateTableSQL(t *testing.T) {
	oplog := &model.Oplog{
		Ns: "test.student",
		O: map[string]interface{}{
			"_id":           "635b79e231d82a8ab1de863b",
			"name":          "Selena Miller",
			"roll_no":       51.0,
			"is_graduated":  false,
			"date_of_birth": "2000-01-30",
		},
	}

	query, err := GenerateCreateTableSQL(oplog)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "CREATE TABLE test.student (_id VARCHAR(255) PRIMARY KEY, date_of_birth VARCHAR(255), is_graduated BOOLEAN, name VARCHAR(255), roll_no FLOAT);"

	if query != expected {
		t.Errorf("got %q, want %q", query, expected)
	}
}
