// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package category

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql/internal/todogotype/ent/schema/bigintgql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldTypes holds the string denoting the types field in the database.
	FieldTypes = "types"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldStrings holds the string denoting the strings field in the database.
	FieldStrings = "strings"
	// EdgeTodos holds the string denoting the todos edge name in mutations.
	EdgeTodos = "todos"
	// EdgeSubCategories holds the string denoting the sub_categories edge name in mutations.
	EdgeSubCategories = "sub_categories"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// TodosTable is the table that holds the todos relation/edge.
	TodosTable = "todos"
	// TodosInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodosInverseTable = "todos"
	// TodosColumn is the table column denoting the todos relation/edge.
	TodosColumn = "category_id"
	// SubCategoriesTable is the table that holds the sub_categories relation/edge. The primary key declared below.
	SubCategoriesTable = "category_sub_categories"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldStatus,
	FieldConfig,
	FieldTypes,
	FieldDuration,
	FieldCount,
}

var (
	// SubCategoriesPrimaryKey and SubCategoriesColumn2 are the table columns denoting the
	// primary key for the sub_categories relation (M2M).
	SubCategoriesPrimaryKey = []string{"category_id", "sub_category_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for _, f := range [...]string{FieldStrings} {
		if column == f {
			return true
		}
	}
	return false
}

var (
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() bigintgql.BigInt
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusEnabled  Status = "ENABLED"
	StatusDisabled Status = "DISABLED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusEnabled, StatusDisabled:
		return nil
	default:
		return fmt.Errorf("category: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Category queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByConfig orders the results by the config field.
func ByConfig(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfig, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByCount orders the results by the count field.
func ByCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCount, opts...).ToFunc()
}

// ByTodosCount orders the results by todos count.
func ByTodosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTodosStep(), opts...)
	}
}

// ByTodos orders the results by todos terms.
func ByTodos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubCategoriesCount orders the results by sub_categories count.
func BySubCategoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubCategoriesStep(), opts...)
	}
}

// BySubCategories orders the results by sub_categories terms.
func BySubCategories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTodosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TodosTable, TodosColumn),
	)
}
func newSubCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SubCategoriesTable, SubCategoriesPrimaryKey...),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
