package validator

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
					"name": "Virat Kohli",
				},
			},
		},
		O2: map[string]interface{}{
			"_id": "635b79e231d82a8ab1de863b",
		},
	}
}

func TestValidateUpdateOplog(t *testing.T) {
	tests := []struct {
		name    string
		oplog   *model.Oplog
		wantErr bool
	}{
		{
			name:    "valid update oplog",
			oplog:   validUpdateOplog(),
			wantErr: false,
		},
		{
			name:    "nil oplog",
			oplog:   nil,
			wantErr: true,
		},
		{
			name: "invalid operation type",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()
				oplog.Op = "i"

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing namespace",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()
				oplog.Ns = ""

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing update operation data",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()
				oplog.O = nil

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing diff data",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()

				oplog.O = map[string]interface{}{
					"$v": 2,
				}

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing update operations",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()

				oplog.O["diff"] = map[string]interface{}{}

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing update criteria",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()
				oplog.O2 = nil

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing _id in update criteria",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()

				oplog.O2 = map[string]interface{}{}

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "valid delete operation",
			oplog: func() *model.Oplog {
				oplog := validUpdateOplog()

				oplog.O["diff"] = map[string]interface{}{
					"d": map[string]interface{}{
						"roll_no": false,
					},
				}

				return oplog
			}(),
			wantErr: false,
		},
		{
			name: "valid mixed update and delete operations",
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
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateUpdateOplog(tc.oplog)

			if (err != nil) != tc.wantErr {
				t.Fatalf(
					"expected error=%v, got err=%v",
					tc.wantErr,
					err,
				)
			}
		})
	}
}
