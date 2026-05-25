package sql

import (
	"testing"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func TestGenerateInsertSQL(t *testing.T) {
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

	resultQuery, err := GenerateInsertSQL(oplog)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedQuery := "INSERT INTO test.student (_id, date_of_birth, is_graduated, name, roll_no) VALUES ('d7cb2b45-c2b1-43ea-a049-619bbbf27037', '2001-06-30', false, 'Raj Deshmukh', 24);"

	if resultQuery != expectedQuery {
		t.Errorf("got %q, want %q", resultQuery, expectedQuery)
	}
}

func TestFormatValue(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "string value",
			input:    "Raj",
			expected: "'Raj'",
		},
		{
			name:     "integer value",
			input:    51,
			expected: "51",
		},
		{
			name:     "boolean true",
			input:    true,
			expected: "true",
		},
		{
			name:     "boolean false",
			input:    false,
			expected: "false",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := formatValue(tc.input)

			if got != tc.expected {
				t.Errorf("got %q, want %q", got, tc.expected)
			}
		})
	}
}
