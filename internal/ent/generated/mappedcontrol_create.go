// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/control"
	"github.com/theopenlane/core/internal/ent/generated/mappedcontrol"
	"github.com/theopenlane/core/internal/ent/generated/subcontrol"
)

// MappedControlCreate is the builder for creating a MappedControl entity.
type MappedControlCreate struct {
	config
	mutation *MappedControlMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mcc *MappedControlCreate) SetCreatedAt(t time.Time) *MappedControlCreate {
	mcc.mutation.SetCreatedAt(t)
	return mcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableCreatedAt(t *time.Time) *MappedControlCreate {
	if t != nil {
		mcc.SetCreatedAt(*t)
	}
	return mcc
}

// SetUpdatedAt sets the "updated_at" field.
func (mcc *MappedControlCreate) SetUpdatedAt(t time.Time) *MappedControlCreate {
	mcc.mutation.SetUpdatedAt(t)
	return mcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableUpdatedAt(t *time.Time) *MappedControlCreate {
	if t != nil {
		mcc.SetUpdatedAt(*t)
	}
	return mcc
}

// SetCreatedBy sets the "created_by" field.
func (mcc *MappedControlCreate) SetCreatedBy(s string) *MappedControlCreate {
	mcc.mutation.SetCreatedBy(s)
	return mcc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableCreatedBy(s *string) *MappedControlCreate {
	if s != nil {
		mcc.SetCreatedBy(*s)
	}
	return mcc
}

// SetUpdatedBy sets the "updated_by" field.
func (mcc *MappedControlCreate) SetUpdatedBy(s string) *MappedControlCreate {
	mcc.mutation.SetUpdatedBy(s)
	return mcc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableUpdatedBy(s *string) *MappedControlCreate {
	if s != nil {
		mcc.SetUpdatedBy(*s)
	}
	return mcc
}

// SetDeletedAt sets the "deleted_at" field.
func (mcc *MappedControlCreate) SetDeletedAt(t time.Time) *MappedControlCreate {
	mcc.mutation.SetDeletedAt(t)
	return mcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableDeletedAt(t *time.Time) *MappedControlCreate {
	if t != nil {
		mcc.SetDeletedAt(*t)
	}
	return mcc
}

// SetDeletedBy sets the "deleted_by" field.
func (mcc *MappedControlCreate) SetDeletedBy(s string) *MappedControlCreate {
	mcc.mutation.SetDeletedBy(s)
	return mcc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableDeletedBy(s *string) *MappedControlCreate {
	if s != nil {
		mcc.SetDeletedBy(*s)
	}
	return mcc
}

// SetTags sets the "tags" field.
func (mcc *MappedControlCreate) SetTags(s []string) *MappedControlCreate {
	mcc.mutation.SetTags(s)
	return mcc
}

// SetMappingType sets the "mapping_type" field.
func (mcc *MappedControlCreate) SetMappingType(s string) *MappedControlCreate {
	mcc.mutation.SetMappingType(s)
	return mcc
}

// SetNillableMappingType sets the "mapping_type" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableMappingType(s *string) *MappedControlCreate {
	if s != nil {
		mcc.SetMappingType(*s)
	}
	return mcc
}

// SetRelation sets the "relation" field.
func (mcc *MappedControlCreate) SetRelation(s string) *MappedControlCreate {
	mcc.mutation.SetRelation(s)
	return mcc
}

// SetNillableRelation sets the "relation" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableRelation(s *string) *MappedControlCreate {
	if s != nil {
		mcc.SetRelation(*s)
	}
	return mcc
}

// SetID sets the "id" field.
func (mcc *MappedControlCreate) SetID(s string) *MappedControlCreate {
	mcc.mutation.SetID(s)
	return mcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mcc *MappedControlCreate) SetNillableID(s *string) *MappedControlCreate {
	if s != nil {
		mcc.SetID(*s)
	}
	return mcc
}

// AddControlIDs adds the "controls" edge to the Control entity by IDs.
func (mcc *MappedControlCreate) AddControlIDs(ids ...string) *MappedControlCreate {
	mcc.mutation.AddControlIDs(ids...)
	return mcc
}

// AddControls adds the "controls" edges to the Control entity.
func (mcc *MappedControlCreate) AddControls(c ...*Control) *MappedControlCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mcc.AddControlIDs(ids...)
}

// AddSubcontrolIDs adds the "subcontrols" edge to the Subcontrol entity by IDs.
func (mcc *MappedControlCreate) AddSubcontrolIDs(ids ...string) *MappedControlCreate {
	mcc.mutation.AddSubcontrolIDs(ids...)
	return mcc
}

// AddSubcontrols adds the "subcontrols" edges to the Subcontrol entity.
func (mcc *MappedControlCreate) AddSubcontrols(s ...*Subcontrol) *MappedControlCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mcc.AddSubcontrolIDs(ids...)
}

// Mutation returns the MappedControlMutation object of the builder.
func (mcc *MappedControlCreate) Mutation() *MappedControlMutation {
	return mcc.mutation
}

// Save creates the MappedControl in the database.
func (mcc *MappedControlCreate) Save(ctx context.Context) (*MappedControl, error) {
	if err := mcc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mcc.sqlSave, mcc.mutation, mcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mcc *MappedControlCreate) SaveX(ctx context.Context) *MappedControl {
	v, err := mcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcc *MappedControlCreate) Exec(ctx context.Context) error {
	_, err := mcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcc *MappedControlCreate) ExecX(ctx context.Context) {
	if err := mcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcc *MappedControlCreate) defaults() error {
	if _, ok := mcc.mutation.CreatedAt(); !ok {
		if mappedcontrol.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized mappedcontrol.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := mappedcontrol.DefaultCreatedAt()
		mcc.mutation.SetCreatedAt(v)
	}
	if _, ok := mcc.mutation.UpdatedAt(); !ok {
		if mappedcontrol.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized mappedcontrol.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := mappedcontrol.DefaultUpdatedAt()
		mcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mcc.mutation.Tags(); !ok {
		v := mappedcontrol.DefaultTags
		mcc.mutation.SetTags(v)
	}
	if _, ok := mcc.mutation.ID(); !ok {
		if mappedcontrol.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized mappedcontrol.DefaultID (forgotten import generated/runtime?)")
		}
		v := mappedcontrol.DefaultID()
		mcc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mcc *MappedControlCreate) check() error {
	return nil
}

func (mcc *MappedControlCreate) sqlSave(ctx context.Context) (*MappedControl, error) {
	if err := mcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected MappedControl.ID type: %T", _spec.ID.Value)
		}
	}
	mcc.mutation.id = &_node.ID
	mcc.mutation.done = true
	return _node, nil
}

func (mcc *MappedControlCreate) createSpec() (*MappedControl, *sqlgraph.CreateSpec) {
	var (
		_node = &MappedControl{config: mcc.config}
		_spec = sqlgraph.NewCreateSpec(mappedcontrol.Table, sqlgraph.NewFieldSpec(mappedcontrol.FieldID, field.TypeString))
	)
	_spec.Schema = mcc.schemaConfig.MappedControl
	if id, ok := mcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mcc.mutation.CreatedAt(); ok {
		_spec.SetField(mappedcontrol.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mcc.mutation.UpdatedAt(); ok {
		_spec.SetField(mappedcontrol.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mcc.mutation.CreatedBy(); ok {
		_spec.SetField(mappedcontrol.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := mcc.mutation.UpdatedBy(); ok {
		_spec.SetField(mappedcontrol.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := mcc.mutation.DeletedAt(); ok {
		_spec.SetField(mappedcontrol.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := mcc.mutation.DeletedBy(); ok {
		_spec.SetField(mappedcontrol.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := mcc.mutation.Tags(); ok {
		_spec.SetField(mappedcontrol.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := mcc.mutation.MappingType(); ok {
		_spec.SetField(mappedcontrol.FieldMappingType, field.TypeString, value)
		_node.MappingType = value
	}
	if value, ok := mcc.mutation.Relation(); ok {
		_spec.SetField(mappedcontrol.FieldRelation, field.TypeString, value)
		_node.Relation = value
	}
	if nodes := mcc.mutation.ControlsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mappedcontrol.ControlsTable,
			Columns: mappedcontrol.ControlsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(control.FieldID, field.TypeString),
			},
		}
		edge.Schema = mcc.schemaConfig.MappedControlControls
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mcc.mutation.SubcontrolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mappedcontrol.SubcontrolsTable,
			Columns: mappedcontrol.SubcontrolsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subcontrol.FieldID, field.TypeString),
			},
		}
		edge.Schema = mcc.schemaConfig.MappedControlSubcontrols
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MappedControlCreateBulk is the builder for creating many MappedControl entities in bulk.
type MappedControlCreateBulk struct {
	config
	err      error
	builders []*MappedControlCreate
}

// Save creates the MappedControl entities in the database.
func (mccb *MappedControlCreateBulk) Save(ctx context.Context) ([]*MappedControl, error) {
	if mccb.err != nil {
		return nil, mccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mccb.builders))
	nodes := make([]*MappedControl, len(mccb.builders))
	mutators := make([]Mutator, len(mccb.builders))
	for i := range mccb.builders {
		func(i int, root context.Context) {
			builder := mccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MappedControlMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mccb *MappedControlCreateBulk) SaveX(ctx context.Context) []*MappedControl {
	v, err := mccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mccb *MappedControlCreateBulk) Exec(ctx context.Context) error {
	_, err := mccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccb *MappedControlCreateBulk) ExecX(ctx context.Context) {
	if err := mccb.Exec(ctx); err != nil {
		panic(err)
	}
}
