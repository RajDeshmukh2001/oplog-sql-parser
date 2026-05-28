package sql

import "fmt"

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
