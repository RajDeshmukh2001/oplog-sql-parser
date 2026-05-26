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
