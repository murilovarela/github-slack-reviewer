-- create "organizations" table
CREATE TABLE "public"."organizations" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "email" text NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_organizations_deleted_at" to table: "organizations"
CREATE INDEX "idx_organizations_deleted_at" ON "public"."organizations" ("deleted_at");
-- create index "uni_organizations_email" to table: "organizations"
CREATE UNIQUE INDEX "uni_organizations_email" ON "public"."organizations" ("email");
-- create "configs" table
CREATE TABLE "public"."configs" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "organization_id" bigint NULL,
  "time_to_summary" bigint NULL,
  "time_to_reminder" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_organizations_config" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_configs_deleted_at" to table: "configs"
CREATE INDEX "idx_configs_deleted_at" ON "public"."configs" ("deleted_at");
-- create "pullrequests" table
CREATE TABLE "public"."pullrequests" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "organization_id" bigint NULL,
  "github_ref" text NULL,
  "latest_message_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_organizations_pullrequests" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_pullrequests_deleted_at" to table: "pullrequests"
CREATE INDEX "idx_pullrequests_deleted_at" ON "public"."pullrequests" ("deleted_at");
-- create index "uni_pullrequests_github_ref" to table: "pullrequests"
CREATE UNIQUE INDEX "uni_pullrequests_github_ref" ON "public"."pullrequests" ("github_ref");
-- create "messages" table
CREATE TABLE "public"."messages" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "organization_id" bigint NULL,
  "slack_ref" text NULL,
  "content" text NULL,
  "message_type" text NULL,
  "pullrequest_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_organizations_messages" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_pullrequests_messages" FOREIGN KEY ("pullrequest_id") REFERENCES "public"."pullrequests" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_messages_deleted_at" to table: "messages"
CREATE INDEX "idx_messages_deleted_at" ON "public"."messages" ("deleted_at");
-- create index "uni_messages_slack_ref" to table: "messages"
CREATE UNIQUE INDEX "uni_messages_slack_ref" ON "public"."messages" ("slack_ref");
-- create "reviewers" table
CREATE TABLE "public"."reviewers" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "organization_id" bigint NULL,
  "github_id" text NULL,
  "slack_id" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_organizations_reviewers" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_reviewers_deleted_at" to table: "reviewers"
CREATE INDEX "idx_reviewers_deleted_at" ON "public"."reviewers" ("deleted_at");
-- create "pullrequest_reviewers" table
CREATE TABLE "public"."pullrequest_reviewers" (
  "pullrequest_id" bigint NOT NULL,
  "reviewer_id" bigint NOT NULL,
  PRIMARY KEY ("pullrequest_id", "reviewer_id"),
  CONSTRAINT "fk_pullrequest_reviewers_pullrequest" FOREIGN KEY ("pullrequest_id") REFERENCES "public"."pullrequests" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_pullrequest_reviewers_reviewer" FOREIGN KEY ("reviewer_id") REFERENCES "public"."reviewers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
