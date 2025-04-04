-- +goose Up
-- modify "organization_setting_history" table
ALTER TABLE "organization_setting_history" DROP COLUMN "billing_address";
ALTER TABLE "organization_setting_history" ADD COLUMN "billing_address" jsonb;
-- modify "organization_settings" table
ALTER TABLE "organization_settings" DROP COLUMN "billing_address";
ALTER TABLE "organization_settings" ADD COLUMN "billing_address" jsonb;

-- +goose Down
-- reverse: modify "organization_settings" table
ALTER TABLE "organization_settings" DROP COLUMN "billing_address";
ALTER TABLE "organization_settings" ADD COLUMN "billing_address" character varying;
-- reverse: modify "organization_setting_history" table
ALTER TABLE "organization_setting_history" DROP COLUMN "billing_address";
ALTER TABLE "organization_setting_history" ADD COLUMN "billing_address" character varying;
