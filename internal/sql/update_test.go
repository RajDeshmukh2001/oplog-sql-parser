package sql

import (
	"testing"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func validUpdateOplog() *model.Oplog {
	return &model.Oplog{
		Op: "u",
		Ns: "test.student",
		O: map[string]interface{}{
			"$v": 2,
			"diff": map[string]interface{}{
				"u": map[string]interface{}{
					"name":         "Virat Kohli",
					"is_graduated": false,
				},
			},
		},
		O2: map[string]interface{}{
			"_id": "635b79e231d82a8ab1de863b",
		},
	}
}

func TestGenerateUpdateSQL(t *testing.T) {
	tests := []struct {
		name     string
		oplog    *model.Oplog
		expected string
	}{
		{
			name:     "generate update SQL for update fields",
			oplog:    validUpdateOplog(),
			expected: "UPDATE test.student SET is_graduated = false, name = 'Virat Kohli' WHERE _id = '635b79e231d82a8ab1de863b';",
		},
		{
			name: "generate update SQL for deleted fields",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()

				oplog.O["diff"] = map[string]interface{}{
					"d": map[string]interface{}{
						"roll_no": false,
					},
				}

				return oplog
			}(),
			expected: "UPDATE test.student SET roll_no = NULL WHERE _id = '635b79e231d82a8ab1de863b';",
		},
		{
			name: "generate update SQL for mixed update operations",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()

				oplog.O["diff"] = map[string]interface{}{
					"u": map[string]interface{}{
						"name": "Virat Kohli",
					},
					"d": map[string]interface{}{
						"roll_no": false,
					},
				}

				return oplog
			}(),
			expected: "UPDATE test.student SET name = 'Virat Kohli', roll_no = NULL WHERE _id = '635b79e231d82a8ab1de863b';",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			query, err := GenerateUpdateSQL(tc.oplog)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if query != tc.expected {
				t.Errorf("got %q, want %q", query, tc.expected)
			}
		})
	}
}
