package validator

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

func TestValidateDeleteOplog(t *testing.T) {
	tests := []struct {
		name    string
		oplog   *model.Oplog
		wantErr bool
	}{
		{
			name:    "valid delete oplog",
			oplog:   validDeleteOplog(),
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
				oplog := validDeleteOplog()

				oplog.Op = "i"

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing namespace",
			oplog: func() *model.Oplog {
				oplog := validDeleteOplog()

				oplog.Ns = ""

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing delete operation data",
			oplog: func() *model.Oplog {
				oplog := validDeleteOplog()

				oplog.O = nil

				return oplog
			}(),
			wantErr: true,
		},
		{
			name: "missing _id in delete criteria",
			oplog: func() *model.Oplog {
				oplog := validDeleteOplog()

				oplog.O = map[string]interface{}{}

				return oplog
			}(),
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateDeleteOplog(tc.oplog)

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
