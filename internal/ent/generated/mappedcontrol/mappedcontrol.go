// Code generated by ent, DO NOT EDIT.

package mappedcontrol

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the mappedcontrol type in the database.
	Label = "mapped_control"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldMappingType holds the string denoting the mapping_type field in the database.
	FieldMappingType = "mapping_type"
	// FieldRelation holds the string denoting the relation field in the database.
	FieldRelation = "relation"
	// EdgeControls holds the string denoting the controls edge name in mutations.
	EdgeControls = "controls"
	// EdgeSubcontrols holds the string denoting the subcontrols edge name in mutations.
	EdgeSubcontrols = "subcontrols"
	// Table holds the table name of the mappedcontrol in the database.
	Table = "mapped_controls"
	// ControlsTable is the table that holds the controls relation/edge. The primary key declared below.
	ControlsTable = "mapped_control_controls"
	// ControlsInverseTable is the table name for the Control entity.
	// It exists in this package in order to avoid circular dependency with the "control" package.
	ControlsInverseTable = "controls"
	// SubcontrolsTable is the table that holds the subcontrols relation/edge. The primary key declared below.
	SubcontrolsTable = "mapped_control_subcontrols"
	// SubcontrolsInverseTable is the table name for the Subcontrol entity.
	// It exists in this package in order to avoid circular dependency with the "subcontrol" package.
	SubcontrolsInverseTable = "subcontrols"
)

// Columns holds all SQL columns for mappedcontrol fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldTags,
	FieldMappingType,
	FieldRelation,
}

var (
	// ControlsPrimaryKey and ControlsColumn2 are the table columns denoting the
	// primary key for the controls relation (M2M).
	ControlsPrimaryKey = []string{"mapped_control_id", "control_id"}
	// SubcontrolsPrimaryKey and SubcontrolsColumn2 are the table columns denoting the
	// primary key for the subcontrols relation (M2M).
	SubcontrolsPrimaryKey = []string{"mapped_control_id", "subcontrol_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/theopenlane/core/internal/ent/generated/runtime"
var (
	Hooks        [3]ent.Hook
	Interceptors [1]ent.Interceptor
	Policy       ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the MappedControl queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByMappingType orders the results by the mapping_type field.
func ByMappingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMappingType, opts...).ToFunc()
}

// ByRelation orders the results by the relation field.
func ByRelation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelation, opts...).ToFunc()
}

// ByControlsCount orders the results by controls count.
func ByControlsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newControlsStep(), opts...)
	}
}

// ByControls orders the results by controls terms.
func ByControls(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newControlsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubcontrolsCount orders the results by subcontrols count.
func BySubcontrolsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubcontrolsStep(), opts...)
	}
}

// BySubcontrols orders the results by subcontrols terms.
func BySubcontrols(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubcontrolsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newControlsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ControlsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ControlsTable, ControlsPrimaryKey...),
	)
}
func newSubcontrolsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubcontrolsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SubcontrolsTable, SubcontrolsPrimaryKey...),
	)
}
