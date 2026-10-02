package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/gertd/go-pluralize"
	"github.com/theopenlane/core/common/models"
	"github.com/theopenlane/entx"
	"github.com/theopenlane/entx/accessmap"
	"github.com/theopenlane/iam/entfga"

	"github.com/theopenlane/core/v2/internal/ent/generated"
	"github.com/theopenlane/core/v2/internal/ent/interceptors"
	"github.com/theopenlane/core/v2/internal/ent/privacy/policy"
	"github.com/theopenlane/core/v2/internal/ent/privacy/rule"
)

// TrustcenterNDARequestSetting holds the schema definition for the TrustcenterNDARequestSetting entity
type TrustcenterNDARequestSetting struct {
	SchemaFuncs

	ent.Schema
}

// SchemaTrustcenterNDARequestSettingis the name of the schema in snake case
const SchemaTrustcenterNDARequestSetting = "trustcenter_nda_request_setting"

// Name is the name of the schema in snake case
func (TrustcenterNDARequestSetting) Name() string {
	return SchemaTrustcenterNDARequestSetting
}

// GetType returns the type of the schema
func (TrustcenterNDARequestSetting) GetType() any {
	return TrustcenterNDARequestSetting.Type
}

// PluralName returns the plural name of the schema
func (TrustcenterNDARequestSetting) PluralName() string {
	return pluralize.NewClient().Plural(SchemaTrustcenterNDARequestSetting)
}

// Fields of the TrustcenterNDARequestSetting
func (TrustcenterNDARequestSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("approver_group_id").
			Comment("group whose members approve trust center NDA requests").
			Optional().
			Nillable(),
		field.Bool("approval_required").
			Comment("whether NDA requests require approval before being processed").
			Default(false).
			Optional(),
		field.Bool("auto_approve").
			Comment("Auto approve NDA requests based on certain rules").
			Default(false).
			Optional(),
		field.Bool("work_email_only").
			Comment("require only work email when auto approving a nda request").
			Default(true).
			Optional(),
		field.Bool("use_domain_blocklist").
			Comment("Enable the use of a domain blocklist").
			Default(false).
			Optional(),
		field.Bool("use_domain_allowlist").
			Comment("Enable the use of a domain allowlist").
			Default(false).
			Optional(),
		field.Bool("approve_from_existing_domain").
			Comment("automatically approve if the request uses the same domain as that of an existing domain").
			Default(true).
			Optional(),
		field.Bool("approve_if_contact_exists").
			Comment("automatically approve if the request email is already a contact object").
			Default(true).
			Optional(),
	}
}

// Mixin of the TrustcenterNDARequestSetting
func (t TrustcenterNDARequestSetting) Mixin() []ent.Mixin {
	return mixinConfig{
		excludeTags: true,
		additionalMixins: []ent.Mixin{
			// newObjectOwnedMixin[generated.TrustCenterSetting](t,
			// 	withParents(TrustCenter{}),
			// ),
			newGroupPermissionsMixin(withSkipViewPermissions()),
		},
	}.getMixins(t)
}

// Edges of the TrustcenterNDARequestSetting
func (t TrustcenterNDARequestSetting) Edges() []ent.Edge {
	return []ent.Edge{
		uniqueEdgeTo(&edgeDefinition{
			fromSchema: t,
			name:       "approver_group",
			t:          Group.Type,
			field:      "nda_approver_group_id",
			annotations: []schema.Annotation{
				accessmap.EdgeViewCheck(Group{}.Name()),
			},
		}),
	}
}

// Indexes of the TrustcenterNDARequestSetting
func (TrustcenterNDARequestSetting) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations of the TrustcenterNDARequestSetting
func (TrustcenterNDARequestSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.FileCategory(SchemaTrustCenterSetting),
		entfga.SettingsChecks("trust_center"),
		entx.FGACrudSkip(entx.SkipDelete | entx.SkipCreate),
	}
}

// Hooks of the TrustcenterNDARequestSetting
func (TrustcenterNDARequestSetting) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Interceptors of the TrustcenterNDARequestSetting
func (TrustcenterNDARequestSetting) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.InterceptorTrustCenterChild(),
	}
}

// Modules this schema has access to
func (TrustcenterNDARequestSetting) Modules() []models.OrgModule {
	return []models.OrgModule{
		models.CatalogTrustCenterModule,
	}
}

// Policy of the TrustcenterNDARequestSetting
func (TrustcenterNDARequestSetting) Policy() ent.Policy {
	return policy.NewPolicy(
		policy.WithOnMutationRules(ent.OpCreate,
			rule.AllowIfTrustCenterEditor(),
			policy.CanCreateObjectsUnderParents([]string{
				TrustCenter{}.Name(),
			}),
			policy.CheckOrgWriteAccess(),
		),
		policy.WithOnMutationRules(ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne,
			rule.AllowIfTrustCenterEditor(),
			// entfga.CheckEditAccess[*generated.TrustCenterSettingMutation](),
		),
	)
}
