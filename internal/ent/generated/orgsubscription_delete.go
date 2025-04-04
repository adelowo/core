// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
	"github.com/theopenlane/core/internal/ent/generated/orgsubscription"
)

// OrgSubscriptionDelete is the builder for deleting a OrgSubscription entity.
type OrgSubscriptionDelete struct {
	config
	hooks    []Hook
	mutation *OrgSubscriptionMutation
}

// Where appends a list predicates to the OrgSubscriptionDelete builder.
func (osd *OrgSubscriptionDelete) Where(ps ...predicate.OrgSubscription) *OrgSubscriptionDelete {
	osd.mutation.Where(ps...)
	return osd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (osd *OrgSubscriptionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, osd.sqlExec, osd.mutation, osd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (osd *OrgSubscriptionDelete) ExecX(ctx context.Context) int {
	n, err := osd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (osd *OrgSubscriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgsubscription.Table, sqlgraph.NewFieldSpec(orgsubscription.FieldID, field.TypeString))
	_spec.Node.Schema = osd.schemaConfig.OrgSubscription
	ctx = internal.NewSchemaConfigContext(ctx, osd.schemaConfig)
	if ps := osd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, osd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	osd.mutation.done = true
	return affected, err
}

// OrgSubscriptionDeleteOne is the builder for deleting a single OrgSubscription entity.
type OrgSubscriptionDeleteOne struct {
	osd *OrgSubscriptionDelete
}

// Where appends a list predicates to the OrgSubscriptionDelete builder.
func (osdo *OrgSubscriptionDeleteOne) Where(ps ...predicate.OrgSubscription) *OrgSubscriptionDeleteOne {
	osdo.osd.mutation.Where(ps...)
	return osdo
}

// Exec executes the deletion query.
func (osdo *OrgSubscriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := osdo.osd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgsubscription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (osdo *OrgSubscriptionDeleteOne) ExecX(ctx context.Context) {
	if err := osdo.Exec(ctx); err != nil {
		panic(err)
	}
}
