package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/internal/ent/generated/internalpolicy"
	"github.com/theopenlane/core/internal/graphapi/model"
	"github.com/theopenlane/gqlgen-plugins/graphutils"
	"github.com/theopenlane/utils/rout"
)

// CreateInternalPolicy is the resolver for the createInternalPolicy field.
func (r *mutationResolver) CreateInternalPolicy(ctx context.Context, input generated.CreateInternalPolicyInput) (*model.InternalPolicyCreatePayload, error) {
	// grab preloads to set max result limits
	graphutils.GetPreloads(ctx, r.maxResultLimit)

	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, input.OwnerID); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	res, err := withTransactionalMutation(ctx).InternalPolicy.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "internalpolicy"})
	}

	return &model.InternalPolicyCreatePayload{
		InternalPolicy: res,
	}, nil
}

// CreateBulkInternalPolicy is the resolver for the createBulkInternalPolicy field.
func (r *mutationResolver) CreateBulkInternalPolicy(ctx context.Context, input []*generated.CreateInternalPolicyInput) (*model.InternalPolicyBulkCreatePayload, error) {
	if len(input) == 0 {
		return nil, rout.NewMissingRequiredFieldError("input")
	}

	// set the organization in the auth context if its not done for us
	// this will choose the first input OwnerID when using a personal access token
	if err := setOrganizationInAuthContextBulkRequest(ctx, input); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	// grab preloads to set max result limits
	graphutils.GetPreloads(ctx, r.maxResultLimit)

	return r.bulkCreateInternalPolicy(ctx, input)
}

// CreateBulkCSVInternalPolicy is the resolver for the createBulkCSVInternalPolicy field.
func (r *mutationResolver) CreateBulkCSVInternalPolicy(ctx context.Context, input graphql.Upload) (*model.InternalPolicyBulkCreatePayload, error) {
	// grab preloads to set max result limits
	graphutils.GetPreloads(ctx, r.maxResultLimit)

	data, err := unmarshalBulkData[generated.CreateInternalPolicyInput](input)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal bulk data")

		return nil, err
	}

	if len(data) == 0 {
		return nil, rout.NewMissingRequiredFieldError("input")
	}

	// set the organization in the auth context if its not done for us
	// this will choose the first input OwnerID when using a personal access token
	if err := setOrganizationInAuthContextBulkRequest(ctx, data); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	return r.bulkCreateInternalPolicy(ctx, data)
}

// UpdateInternalPolicy is the resolver for the updateInternalPolicy field.
func (r *mutationResolver) UpdateInternalPolicy(ctx context.Context, id string, input generated.UpdateInternalPolicyInput) (*model.InternalPolicyUpdatePayload, error) {
	// grab preloads to set max result limits
	graphutils.GetPreloads(ctx, r.maxResultLimit)

	res, err := withTransactionalMutation(ctx).InternalPolicy.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "internalpolicy"})
	}

	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, &res.OwnerID); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.ErrPermissionDenied
	}

	// setup update request
	req := res.Update().SetInput(input).AppendTags(input.AppendTags)

	res, err = req.Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "internalpolicy"})
	}

	return &model.InternalPolicyUpdatePayload{
		InternalPolicy: res,
	}, nil
}

// DeleteInternalPolicy is the resolver for the deleteInternalPolicy field.
func (r *mutationResolver) DeleteInternalPolicy(ctx context.Context, id string) (*model.InternalPolicyDeletePayload, error) {
	if err := withTransactionalMutation(ctx).InternalPolicy.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, parseRequestError(err, action{action: ActionDelete, object: "internalpolicy"})
	}

	if err := generated.InternalPolicyEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &model.InternalPolicyDeletePayload{
		DeletedID: id,
	}, nil
}

// InternalPolicy is the resolver for the internalPolicy field.
func (r *queryResolver) InternalPolicy(ctx context.Context, id string) (*generated.InternalPolicy, error) {
	// get preloads to set max result limits
	graphutils.GetPreloads(ctx, r.maxResultLimit)

	query, err := withTransactionalMutation(ctx).InternalPolicy.Query().Where(internalpolicy.ID(id)).CollectFields(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "internalpolicy"})
	}

	res, err := query.Only(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "internalpolicy"})
	}

	return res, nil
}
