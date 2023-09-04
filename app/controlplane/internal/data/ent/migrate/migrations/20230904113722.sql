-- Create "cas_mappings" table
CREATE TABLE "cas_mappings" ("id" uuid NOT NULL, "digest" character varying NOT NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "cas_mapping_cas_backend" uuid NOT NULL, "cas_mapping_workflow_run" uuid NULL, "cas_mapping_organization" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "cas_mappings_cas_backends_cas_backend" FOREIGN KEY ("cas_mapping_cas_backend") REFERENCES "cas_backends" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "cas_mappings_organizations_organization" FOREIGN KEY ("cas_mapping_organization") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "cas_mappings_workflow_runs_workflow_run" FOREIGN KEY ("cas_mapping_workflow_run") REFERENCES "workflow_runs" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "casmapping_digest" to table: "cas_mappings"
CREATE INDEX "casmapping_digest" ON "cas_mappings" ("digest");
