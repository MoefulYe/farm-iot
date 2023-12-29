// Code generated by ent, DO NOT EDIT.

package device

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the device type in the database.
	Label = "device"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBornAt holds the string denoting the born_at field in the database.
	FieldBornAt = "born_at"
	// FieldHashedPasswd holds the string denoting the hashed_passwd field in the database.
	FieldHashedPasswd = "hashed_passwd"
	// FieldDeadAt holds the string denoting the dead_at field in the database.
	FieldDeadAt = "dead_at"
	// FieldReason holds the string denoting the reason field in the database.
	FieldReason = "reason"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the device in the database.
	Table = "devices"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "devices"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "device_children"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "devices"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "device_children"
)

// Columns holds all SQL columns for device fields.
var Columns = []string{
	FieldID,
	FieldBornAt,
	FieldHashedPasswd,
	FieldDeadAt,
	FieldReason,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "devices"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"device_children",
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

// OrderOption defines the ordering options for the Device queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBornAt orders the results by the born_at field.
func ByBornAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBornAt, opts...).ToFunc()
}

// ByHashedPasswd orders the results by the hashed_passwd field.
func ByHashedPasswd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashedPasswd, opts...).ToFunc()
}

// ByDeadAt orders the results by the dead_at field.
func ByDeadAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeadAt, opts...).ToFunc()
}

// ByReason orders the results by the reason field.
func ByReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReason, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}