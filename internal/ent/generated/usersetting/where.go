// Code generated by ent, DO NOT EDIT.

package usersetting

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/theopenlane/core/internal/ent/generated/predicate"
	"github.com/theopenlane/core/pkg/enums"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldDeletedBy, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUserID, v))
}

// Locked applies equality check predicate on the "locked" field. It's identical to LockedEQ.
func Locked(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldLocked, v))
}

// SilencedAt applies equality check predicate on the "silenced_at" field. It's identical to SilencedAtEQ.
func SilencedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldSilencedAt, v))
}

// SuspendedAt applies equality check predicate on the "suspended_at" field. It's identical to SuspendedAtEQ.
func SuspendedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldSuspendedAt, v))
}

// EmailConfirmed applies equality check predicate on the "email_confirmed" field. It's identical to EmailConfirmedEQ.
func EmailConfirmed(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldEmailConfirmed, v))
}

// IsWebauthnAllowed applies equality check predicate on the "is_webauthn_allowed" field. It's identical to IsWebauthnAllowedEQ.
func IsWebauthnAllowed(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldIsWebauthnAllowed, v))
}

// IsTfaEnabled applies equality check predicate on the "is_tfa_enabled" field. It's identical to IsTfaEnabledEQ.
func IsTfaEnabled(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldIsTfaEnabled, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldPhoneNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldUpdatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldDeletedBy, v))
}

// TagsIsNil applies the IsNil predicate on the "tags" field.
func TagsIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldTags))
}

// TagsNotNil applies the NotNil predicate on the "tags" field.
func TagsNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldTags))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldUserID))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldUserID, v))
}

// LockedEQ applies the EQ predicate on the "locked" field.
func LockedEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldLocked, v))
}

// LockedNEQ applies the NEQ predicate on the "locked" field.
func LockedNEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldLocked, v))
}

// SilencedAtEQ applies the EQ predicate on the "silenced_at" field.
func SilencedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldSilencedAt, v))
}

// SilencedAtNEQ applies the NEQ predicate on the "silenced_at" field.
func SilencedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldSilencedAt, v))
}

// SilencedAtIn applies the In predicate on the "silenced_at" field.
func SilencedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldSilencedAt, vs...))
}

// SilencedAtNotIn applies the NotIn predicate on the "silenced_at" field.
func SilencedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldSilencedAt, vs...))
}

// SilencedAtGT applies the GT predicate on the "silenced_at" field.
func SilencedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldSilencedAt, v))
}

// SilencedAtGTE applies the GTE predicate on the "silenced_at" field.
func SilencedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldSilencedAt, v))
}

// SilencedAtLT applies the LT predicate on the "silenced_at" field.
func SilencedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldSilencedAt, v))
}

// SilencedAtLTE applies the LTE predicate on the "silenced_at" field.
func SilencedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldSilencedAt, v))
}

// SilencedAtIsNil applies the IsNil predicate on the "silenced_at" field.
func SilencedAtIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldSilencedAt))
}

// SilencedAtNotNil applies the NotNil predicate on the "silenced_at" field.
func SilencedAtNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldSilencedAt))
}

// SuspendedAtEQ applies the EQ predicate on the "suspended_at" field.
func SuspendedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldSuspendedAt, v))
}

// SuspendedAtNEQ applies the NEQ predicate on the "suspended_at" field.
func SuspendedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldSuspendedAt, v))
}

// SuspendedAtIn applies the In predicate on the "suspended_at" field.
func SuspendedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldSuspendedAt, vs...))
}

// SuspendedAtNotIn applies the NotIn predicate on the "suspended_at" field.
func SuspendedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldSuspendedAt, vs...))
}

// SuspendedAtGT applies the GT predicate on the "suspended_at" field.
func SuspendedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldSuspendedAt, v))
}

// SuspendedAtGTE applies the GTE predicate on the "suspended_at" field.
func SuspendedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldSuspendedAt, v))
}

// SuspendedAtLT applies the LT predicate on the "suspended_at" field.
func SuspendedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldSuspendedAt, v))
}

// SuspendedAtLTE applies the LTE predicate on the "suspended_at" field.
func SuspendedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldSuspendedAt, v))
}

// SuspendedAtIsNil applies the IsNil predicate on the "suspended_at" field.
func SuspendedAtIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldSuspendedAt))
}

// SuspendedAtNotNil applies the NotNil predicate on the "suspended_at" field.
func SuspendedAtNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldSuspendedAt))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.UserStatus) predicate.UserSetting {
	vc := v
	return predicate.UserSetting(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.UserStatus) predicate.UserSetting {
	vc := v
	return predicate.UserSetting(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.UserStatus) predicate.UserSetting {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSetting(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.UserStatus) predicate.UserSetting {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSetting(sql.FieldNotIn(FieldStatus, v...))
}

// EmailConfirmedEQ applies the EQ predicate on the "email_confirmed" field.
func EmailConfirmedEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldEmailConfirmed, v))
}

// EmailConfirmedNEQ applies the NEQ predicate on the "email_confirmed" field.
func EmailConfirmedNEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldEmailConfirmed, v))
}

// IsWebauthnAllowedEQ applies the EQ predicate on the "is_webauthn_allowed" field.
func IsWebauthnAllowedEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldIsWebauthnAllowed, v))
}

// IsWebauthnAllowedNEQ applies the NEQ predicate on the "is_webauthn_allowed" field.
func IsWebauthnAllowedNEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldIsWebauthnAllowed, v))
}

// IsWebauthnAllowedIsNil applies the IsNil predicate on the "is_webauthn_allowed" field.
func IsWebauthnAllowedIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldIsWebauthnAllowed))
}

// IsWebauthnAllowedNotNil applies the NotNil predicate on the "is_webauthn_allowed" field.
func IsWebauthnAllowedNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldIsWebauthnAllowed))
}

// IsTfaEnabledEQ applies the EQ predicate on the "is_tfa_enabled" field.
func IsTfaEnabledEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldIsTfaEnabled, v))
}

// IsTfaEnabledNEQ applies the NEQ predicate on the "is_tfa_enabled" field.
func IsTfaEnabledNEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldIsTfaEnabled, v))
}

// IsTfaEnabledIsNil applies the IsNil predicate on the "is_tfa_enabled" field.
func IsTfaEnabledIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldIsTfaEnabled))
}

// IsTfaEnabledNotNil applies the NotNil predicate on the "is_tfa_enabled" field.
func IsTfaEnabledNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldIsTfaEnabled))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberIsNil applies the IsNil predicate on the "phone_number" field.
func PhoneNumberIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldPhoneNumber))
}

// PhoneNumberNotNil applies the NotNil predicate on the "phone_number" field.
func PhoneNumberNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldPhoneNumber))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserSetting {
	return predicate.UserSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserSetting
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserSetting {
	return predicate.UserSetting(func(s *sql.Selector) {
		step := newUserStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserSetting
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDefaultOrg applies the HasEdge predicate on the "default_org" edge.
func HasDefaultOrg() predicate.UserSetting {
	return predicate.UserSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DefaultOrgTable, DefaultOrgColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.UserSetting
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDefaultOrgWith applies the HasEdge predicate on the "default_org" edge with a given conditions (other predicates).
func HasDefaultOrgWith(preds ...predicate.Organization) predicate.UserSetting {
	return predicate.UserSetting(func(s *sql.Selector) {
		step := newDefaultOrgStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.UserSetting
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.UserSetting {
	return predicate.UserSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, FilesTable, FilesPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.File
		step.Edge.Schema = schemaConfig.UserSettingFiles
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.UserSetting {
	return predicate.UserSetting(func(s *sql.Selector) {
		step := newFilesStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.File
		step.Edge.Schema = schemaConfig.UserSettingFiles
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserSetting) predicate.UserSetting {
	return predicate.UserSetting(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserSetting) predicate.UserSetting {
	return predicate.UserSetting(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserSetting) predicate.UserSetting {
	return predicate.UserSetting(sql.NotPredicates(p))
}
