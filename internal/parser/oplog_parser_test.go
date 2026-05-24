package parser

import "testing"

func TestParseOplog(t *testing.T) {
	jsonData := []byte(`{
		"op":"i",
		"ns":"test.student",
		"o":{
			"_id":"d7cb2b45-c2b1-43ea-a049-619bbbf27037",
			"name":"Raj Deshmukh",
			"roll_no": 24,
			"date_of_birth": "2001-06-30"
		}
	}`)

	oplog, err := ParseOplog(jsonData)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if oplog.Op != "i" {
		t.Errorf("expected op to be i, got %s", oplog.Op)
	}

	if oplog.Ns != "test.student" {
		t.Errorf("expected namespace test.student, got %s", oplog.Ns)
	}

	if oplog.O["name"] != "Raj Deshmukh" {
		t.Errorf("expected name to be Raj Deshmukh")
	}

	if oplog.O["date_of_birth"] != "2001-06-30" {
		t.Errorf("expected correct date_of_birth")
	}
}
