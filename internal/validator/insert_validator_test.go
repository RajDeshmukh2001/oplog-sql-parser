package validator

import (
	"testing"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func TestValidateInsertOplog(t *testing.T) {
	tests := []struct {
		name    string
		oplog   *model.Oplog
		wantErr bool
	}{
		{
			name: "valid insert oplog",
			oplog: &model.Oplog{
				Op: "i",
				Ns: "test.student",
				O: map[string]interface{}{
					"_id":           "d7cb2b45-c2b1-43ea-a049-619bbbf27037",
					"name":          "Raj Deshmukh",
					"roll_no":       24,
					"is_graduated":  false,
					"date_of_birth": "2001-06-30",
				},
			},
			wantErr: false,
		},
		{
			name:    "nil oplog",
			oplog:   nil,
			wantErr: true,
		},
		{
			name: "invalid operation type",
			oplog: &model.Oplog{
				Op: "u",
				Ns: "test.student",
				O: map[string]interface{}{
					"_id":           "d7cb2b45-c2b1-43ea-a049-619bbbf27037",
					"name":          "Raj Deshmukh",
					"roll_no":       24,
					"is_graduated":  false,
					"date_of_birth": "2001-06-30",
				},
			},
			wantErr: true,
		},
		{
			name: "missing namespace",
			oplog: &model.Oplog{
				Op: "i",
				Ns: "",
				O: map[string]interface{}{
					"_id":           "d7cb2b45-c2b1-43ea-a049-619bbbf27037",
					"name":          "Raj Deshmukh",
					"roll_no":       24,
					"is_graduated":  false,
					"date_of_birth": "2001-06-30",
				},
			},
			wantErr: true,
		},
		{
			name: "empty document",
			oplog: &model.Oplog{
				Op: "i",
				Ns: "test.student",
				O:  map[string]interface{}{},
			},
			wantErr: true,
		},
		{
			name: "nil document map",
			oplog: &model.Oplog{
				Op: "i",
				Ns: "test.student",
				O:  nil,
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateInsertOplog(tc.oplog)

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
