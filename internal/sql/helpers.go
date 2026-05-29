package sql

import "sort"

func sortedKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))

	for key := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func getColumnType(value interface{}) string {
	switch value.(type) {
	case string:
		return "VARCHAR(255)"
	case float64:
		return "FLOAT"
	case bool:
		return "BOOLEAN"
	default:
		return "VARCHAR(255)"
	}
}
