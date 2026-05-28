package sql

import (
	"testing"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func validDeleteOplog() *model.Oplog {
	return &model.Oplog{
		Op: "d",
		Ns: "test.student",
		O: map[string]interface{}{
			"_id": "635b79e231d82a8ab1de863b",
		},
	}
}

func TestGenerateDeleteSQL(t *testing.T) {
	tests := []struct {
		name     string
		oplog    *model.Oplog
		expected string
	}{
		{
			name:     "generate delete SQL",
			oplog:    validDeleteOplog(),
			expected: "DELETE FROM test.student WHERE _id = '635b79e231d82a8ab1de863b';",
		},
		{
			name: "generate delete SQL with numeric id",
			oplog: func() *model.Oplog {
				oplog := validDeleteOplog()

				oplog.O["_id"] = 101

				return oplog
			}(),
			expected: "DELETE FROM test.student WHERE _id = 101;",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			query, err := GenerateDeleteSQL(tc.oplog)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if query != tc.expected {
				t.Errorf("got %q, want %q", query, tc.expected)
			}
		})
	}
}
