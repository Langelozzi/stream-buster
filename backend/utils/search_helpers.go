package utils

import (
	"github.com/STREAM-BUSTER/stream-buster/models/search"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func BuildSearchQuery(db *gorm.DB, model interface{}, searchObj search.Search) *gorm.DB {
	modelType := reflect.TypeOf(model)
	dbTableName := ToSnakeCase(modelType.Name()) + "s"
	tablePrimaryKey := dbTableName + ".id"

	query := db.Model(model)

	// Maintain a list of joined tables, so we don't join them again
	joinedTables := make(map[string]bool)

	for index, param := range searchObj.Params {
		var queryStr string
		var values []interface{}

		if strings.Contains(param.Field, ".") && strings.Trim(param.AssociationForeignKey, " ") != "" {
			// If the field contains a dot, we assume it's a join
			// We need to split the field and join the tables on the primary key
			// Example: "user_documents.user_id" will be joined with "user_document" table
			fields := strings.Split(param.Field, ".")
			joiningTable := fields[0]

			// If the table is not already joined, we join it
			if _, ok := joinedTables[joiningTable]; !ok && joiningTable != dbTableName {
				query = query.Joins("JOIN " + joiningTable + " ON " + fields[0] + "." + param.AssociationForeignKey + " = " + tablePrimaryKey)

				// Add the joining table to the list of joined tables
				joinedTables[joiningTable] = true
			}

			param.Field = fields[0] + "." + fields[1]
		}

		if param.Operator == search.IsNull || param.Operator == search.IsNotNull {
			// For IS NULL and IS NOT NULL operators, the Value should be nil
			queryStr = param.Field + " " + string(param.Operator)
		} else if param.Operator == search.Like || param.Operator == search.NotLike {
			// For LIKE and NOT LIKE operators, we need to wrap the value with %value%
			queryStr = param.Field + " " + string(param.Operator) + " ?"
			values = append(values, param.Value.(string))
		} else if param.Operator == search.In || param.Operator == search.NotIn {
			// For IN operator, the Value should be a slice
			queryStr = param.Field + " " + string(param.Operator) + " (?)"
			values = append(values, param.Value)
		} else {
			// For other operators, we assume single value
			queryStr = param.Field + " " + string(param.Operator) + " ?"
			values = append(values, param.Value)
		}

		if searchObj.LogicalOperator == search.And || index == 0 {
			query = query.Where(queryStr, values...)
		} else {
			query = query.Or(queryStr, values...)
		}
	}

	return query
}
