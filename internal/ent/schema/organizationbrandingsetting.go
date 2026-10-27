package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/gertd/go-pluralize"
	"github.com/theopenlane/entx"
	"github.com/theopenlane/entx/accessmap"

	"github.com/theopenlane/core/common/models"

	"github.com/theopenlane/core/v2/internal/ent/privacy/policy"
	"github.com/theopenlane/core/v2/internal/ent/validator"
)

type OrganizationBrandingSetting struct {
	SchemaFuncs

	ent.Schema
}

const SchemaOrganizationBrandingSetting = "organization_branding_setting"

func (OrganizationBrandingSetting) Name() string {
	return SchemaOrganizationBrandingSetting
}

func (OrganizationBrandingSetting) GetType() any {
	return OrganizationBrandingSetting.Type
}

func (OrganizationBrandingSetting) PluralName() string {
	return pluralize.NewClient().Plural(SchemaOrganizationBrandingSetting)
}

func (OrganizationBrandingSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("remove_assessment_branding").
			Comment("whether to hide the Openlane branding on the Organization's assessments").
			Default(false).
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			),
		field.String("primary_color").
			Comment("the primary brand color of the Organization").
			Validate(validator.HexColorValidator).
			Optional(),
		field.String("secondary_color").
			Comment("the secondary brand color of the Organization").
			Validate(validator.HexColorValidator).
			Optional(),
		field.String("background_color").
			Comment("the background color of the Organization").
			Validate(validator.HexColorValidator).
			Optional(),
		field.String("text_color").
			Comment("the text color of the Organization").
			Validate(validator.HexColorValidator).
			Optional(),
		field.String("font_family").
			Comment("the font family of the Organization").
			Optional(),
		field.String("logo_remote_url").
			Comment("the logo URL of the Organization").
			MaxLen(urlMaxLen).
			Validate(validator.ValidateURL()).
			Optional().
			Nillable(),
		field.String("logo_local_file_id").
			Comment("the local logo file id of the Organization, takes precedence over the logo remote URL").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			).
			Nillable(),
	}
}

func (b OrganizationBrandingSetting) Mixin() []ent.Mixin {
	return mixinConfig{
		excludeTags: true,
		additionalMixins: []ent.Mixin{
			newOrgOwnedMixin(b),
		},
	}.getMixins(b)
}

func (b OrganizationBrandingSetting) Edges() []ent.Edge {
	return []ent.Edge{
		uniqueEdgeTo(&edgeDefinition{
			fromSchema: b,
			name:       "logo_file",
			t:          File.Type,
			field:      "logo_local_file_id",
			annotations: []schema.Annotation{
				accessmap.EdgeViewCheck(File{}.Name()),
			},
		}),
	}
}

func (OrganizationBrandingSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(ownerFieldName).
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

func (OrganizationBrandingSetting) Policy() ent.Policy {
	return policy.NewPolicy(
		policy.WithQueryRules(
			policy.CheckOrgReadAccess(),
		),
		policy.WithMutationRules(
			policy.CheckOrgWriteAccess(),
		),
	)
}

func (OrganizationBrandingSetting) Modules() []models.OrgModule {
	return []models.OrgModule{
		models.CatalogBaseModule,
	}
}

func (OrganizationBrandingSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.FileCategory(SchemaOrganizationBrandingSetting),
		entx.FGACrudParent(Organization{}.Name()),
	}
}
