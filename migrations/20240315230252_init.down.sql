-- reverse: create "pullrequest_reviewers" table
DROP TABLE "public"."pullrequest_reviewers";
-- reverse: create index "idx_reviewers_deleted_at" to table: "reviewers"
DROP INDEX "public"."idx_reviewers_deleted_at";
-- reverse: create "reviewers" table
DROP TABLE "public"."reviewers";
-- reverse: create index "uni_messages_slack_ref" to table: "messages"
DROP INDEX "public"."uni_messages_slack_ref";
-- reverse: create index "idx_messages_deleted_at" to table: "messages"
DROP INDEX "public"."idx_messages_deleted_at";
-- reverse: create "messages" table
DROP TABLE "public"."messages";
-- reverse: create index "uni_pullrequests_github_ref" to table: "pullrequests"
DROP INDEX "public"."uni_pullrequests_github_ref";
-- reverse: create index "idx_pullrequests_deleted_at" to table: "pullrequests"
DROP INDEX "public"."idx_pullrequests_deleted_at";
-- reverse: create "pullrequests" table
DROP TABLE "public"."pullrequests";
-- reverse: create index "idx_configs_deleted_at" to table: "configs"
DROP INDEX "public"."idx_configs_deleted_at";
-- reverse: create "configs" table
DROP TABLE "public"."configs";
-- reverse: create index "uni_organizations_email" to table: "organizations"
DROP INDEX "public"."uni_organizations_email";
-- reverse: create index "idx_organizations_deleted_at" to table: "organizations"
DROP INDEX "public"."idx_organizations_deleted_at";
-- reverse: create "organizations" table
DROP TABLE "public"."organizations";
