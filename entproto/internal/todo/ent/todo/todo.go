// Code generated by ent, DO NOT EDIT.

package todo

import (
	"fmt"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTask holds the string denoting the task field in the database.
	FieldTask = "task"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "user_id"
	// Table holds the table name of the todo in the database.
	Table = "todos"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "todos"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "todo_user"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldTask,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "todos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"todo_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusInProgress, StatusDone:
		return nil
	default:
		return fmt.Errorf("todo: invalid enum value for status field: %q", s)
	}
}
