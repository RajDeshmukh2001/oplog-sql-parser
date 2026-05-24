package parser

import (
	"encoding/json"
	"fmt"

	"github.com/RajDeshmukh2001/oplog-sql-parser/internal/model"
)

func ParseOplog(data []byte) (*model.Oplog, error) {
	var oplog model.Oplog

	err := json.Unmarshal(data, &oplog)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse oplog JSON: %w", err)
	}

	return &oplog, nil
}
