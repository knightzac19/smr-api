// Code generated by ent, DO NOT EDIT.

package versiontarget

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the versiontarget type in the database.
	Label = "version_target"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersionID holds the string denoting the version_id field in the database.
	FieldVersionID = "version_id"
	// FieldTargetName holds the string denoting the target_name field in the database.
	FieldTargetName = "target_name"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// EdgeSmlVersion holds the string denoting the sml_version edge name in mutations.
	EdgeSmlVersion = "sml_version"
	// Table holds the table name of the versiontarget in the database.
	Table = "version_targets"
	// SmlVersionTable is the table that holds the sml_version relation/edge.
	SmlVersionTable = "version_targets"
	// SmlVersionInverseTable is the table name for the Version entity.
	// It exists in this package in order to avoid circular dependency with the "version" package.
	SmlVersionInverseTable = "versions"
	// SmlVersionColumn is the table column denoting the sml_version relation/edge.
	SmlVersionColumn = "version_id"
)

// Columns holds all SQL columns for versiontarget fields.
var Columns = []string{
	FieldID,
	FieldVersionID,
	FieldTargetName,
	FieldKey,
	FieldHash,
	FieldSize,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the VersionTarget queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVersionID orders the results by the version_id field.
func ByVersionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersionID, opts...).ToFunc()
}

// ByTargetName orders the results by the target_name field.
func ByTargetName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetName, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// BySmlVersionField orders the results by sml_version field.
func BySmlVersionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSmlVersionStep(), sql.OrderByField(field, opts...))
	}
}
func newSmlVersionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SmlVersionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SmlVersionTable, SmlVersionColumn),
	)
}