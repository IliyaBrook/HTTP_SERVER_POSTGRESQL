package pkg

import (
	"fmt"
	"strings"
)

// BuildSQLDynamic builds a dynamic SQL query for different types of operations (INSERT, UPDATE, DELETE, SELECT).
//
// Parameters:
//   - queryType: Type of the SQL query ("INSERT", "UPDATE", "DELETE", "SELECT").
//   - tableName: Name of the table to perform the operation on.
//   - data: A map where keys are column names and values are the values to insert or update.
//     For SELECT queries, keys are the column names to select (if the map is empty, it selects all columns).
//   - condition: A string representing the WHERE clause (e.g., "id = $1").
//     **Important**: Ensure there are spaces around operators in the condition string (e.g., use "id = $1" instead of "id=$1").
//   - conditionArgs: Values for placeholders in the WHERE clause.
//
// Returns:
// - A string containing the constructed SQL query.
// - A slice of interface{} containing the arguments to pass to the query.
// - An error if the queryType is not supported or if data is empty for non-DELETE/SELECT queries.
//
// Example Usage:
// 1:
// query, args, err := pkg.BuildSQLDynamic("UPDATE", "users", map[string]interface{}{"name": "John", "age": 30}, "id = $1", 1)
// 2:
// query, args, err := pkg.BuildSQLDynamic("UPDATE", "users", updatedUserData, "id = $1", updateUserId)
//
//	if err != nil {
//		pkg.ResponseErrorText(err, w, "failed to build update query")
//		return
//	}
func BuildSQLDynamic(queryType, tableName string, data map[string]interface{}, condition string, conditionArgs ...interface{}) (string, []interface{}, error) {
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

		condition = reindexConditionParams(condition, argID)
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

func reindexConditionParams(condition string, startIdx int) string {
	var newConditionParts []string
	parts := strings.Split(condition, " ")

	for _, part := range parts {
		if strings.HasPrefix(part, "$") {
			newConditionParts = append(newConditionParts, fmt.Sprintf("$%d", startIdx))
			startIdx++
		} else {
			newConditionParts = append(newConditionParts, part)
		}
	}

	return strings.Join(newConditionParts, " ")
}
