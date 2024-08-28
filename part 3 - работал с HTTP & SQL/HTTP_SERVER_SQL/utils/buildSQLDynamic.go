package utils

import (
	"fmt"
	"strings"
)

// BuildSQLQuery creates a dynamic SQL query for any operation with the provided table name, data map, and condition.
func buildSQLDynamic(queryType, tableName string, data map[string]interface{}, condition string, conditionArgs ...interface{}) (string, []interface{}, error) {
	if len(data) == 0 && queryType != "DELETE" && queryType != "SELECT" {
		return "", nil, fmt.Errorf("no data provided for %s operation", queryType)
	}

	var queryParts []string
	var args []interface{}
	var query string

	switch strings.ToUpper(queryType) {
	case "INSERT":
		var columns []string
		var placeholders []string
		for key, value := range data {
			columns = append(columns, key)
			placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)+1))
			args = append(args, value)
		}
		query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	case "UPDATE":
		argID := 1
		for key, value := range data {
			queryParts = append(queryParts, fmt.Sprintf("%s = $%d", key, argID))
			args = append(args, value)
			argID++
		}
		query = fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(queryParts, ", "), condition)

	case "DELETE":
		query = fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, condition)

	case "SELECT":
		var columns []string
		for key := range data {
			columns = append(columns, key)
		}
		if len(columns) == 0 {
			columns = append(columns, "*")
		}
		query = fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(columns, ", "), tableName, condition)

	default:
		return "", nil, fmt.Errorf("unsupported query type: %s", queryType)
	}

	args = append(args, conditionArgs...)

	return query, args, nil
}
