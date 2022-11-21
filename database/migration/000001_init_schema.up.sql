CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_login" timestamptz NOT NULL DEFAULT (now()),
  "profile_blurb" varchar,
  "banned" boolean
);

CREATE TABLE "DomainBans" (
  "user_id" bigint,
  "banned_from" varchar,
  "banned_by" bigint,
  "banned_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Comments" (
  "id" bigserial PRIMARY KEY,
  "pathid" bigint NOT NULL,
  "author" bigint NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "parent" bigint,
  "hidden" boolean DEFAULT false,
  "removed" boolean DEFAULT false
);

CREATE TABLE "VoteRecords" (
  "comment_id" bigint,
  "category" varchar,
  "user_id" bigint,
  "value" int8,
  PRIMARY KEY ("comment_id", "category")
);

CREATE TABLE "Domains" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL
);

CREATE TABLE "Paths" (
  "id" bigserial PRIMARY KEY,
  "domain" varchar,
  "path" varchar
);

CREATE TABLE "DomainModeratorAssignments" (
  "id" bigserial PRIMARY KEY,
  "domain" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "granted_at" timestamptz NOT NULL DEFAULT (now()),
  "granted_by" bigint NOT NULL
);

CREATE TABLE "GlobalModeratorAssignments" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "granted_at" timestamptz NOT NULL DEFAULT (now()),
  "granted_by" bigint NOT NULL
);

CREATE TABLE "AdminAssignments" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint
);

CREATE TABLE "Logs" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "ip" varchar,
  "url" varchar
);

CREATE TABLE "CommentModerationActions" (
  "id" bigserial PRIMARY KEY,
  "taken_by" bigint,
  "comment_id" bigint,
  "reason" varchar,
  "taken_on" timestamptz,
  "set_hidden_to" boolean,
  "set_removed_to" boolean,
  "associated_report" bigint
);

CREATE TABLE "BanActions" (
  "id" bigserial PRIMARY KEY,
  "taken_by" bigint,
  "target_user_id" bigint,
  "reason" varchar,
  "taken_on" timestamptz,
  "domain" varchar,
  "set_banned_to" boolean
);

CREATE TABLE "Reports" (
  "id" bigserial PRIMARY KEY,
  "reporting_user_id" bigint,
  "comment" bigint,
  "reason" varchar,
  "action_taken" boolean
);

CREATE INDEX ON "Users" ("username");

CREATE INDEX ON "Users" ("email");

CREATE INDEX ON "DomainBans" ("user_id");

CREATE INDEX ON "DomainBans" ("banned_by");

CREATE INDEX ON "DomainBans" ("banned_from");

CREATE INDEX ON "Comments" ("author");

CREATE INDEX ON "Comments" ("pathid");

CREATE UNIQUE INDEX ON "Paths" ("domain", "path");

COMMENT ON COLUMN "Users"."password" IS 'Must be encrypted';

ALTER TABLE "DomainBans" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Comments" ADD FOREIGN KEY ("pathid") REFERENCES "Paths" ("id");

ALTER TABLE "Comments" ADD FOREIGN KEY ("author") REFERENCES "Users" ("id");

ALTER TABLE "VoteRecords" ADD FOREIGN KEY ("comment_id") REFERENCES "Comments" ("id");

ALTER TABLE "VoteRecords" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Paths" ADD FOREIGN KEY ("domain") REFERENCES "Domains" ("id");

ALTER TABLE "DomainModeratorAssignments" ADD FOREIGN KEY ("domain") REFERENCES "Domains" ("id");

ALTER TABLE "DomainModeratorAssignments" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "DomainModeratorAssignments" ADD FOREIGN KEY ("granted_by") REFERENCES "Users" ("id");

ALTER TABLE "GlobalModeratorAssignments" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "GlobalModeratorAssignments" ADD FOREIGN KEY ("granted_by") REFERENCES "Users" ("id");

ALTER TABLE "AdminAssignments" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Logs" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "CommentModerationActions" ADD FOREIGN KEY ("taken_by") REFERENCES "Users" ("id");

ALTER TABLE "CommentModerationActions" ADD FOREIGN KEY ("comment_id") REFERENCES "Comments" ("id");

ALTER TABLE "CommentModerationActions" ADD FOREIGN KEY ("associated_report") REFERENCES "Reports" ("id");

ALTER TABLE "BanActions" ADD FOREIGN KEY ("taken_by") REFERENCES "Users" ("id");

ALTER TABLE "BanActions" ADD FOREIGN KEY ("target_user_id") REFERENCES "Users" ("id");

ALTER TABLE "BanActions" ADD FOREIGN KEY ("domain") REFERENCES "Domains" ("id");

ALTER TABLE "Reports" ADD FOREIGN KEY ("reporting_user_id") REFERENCES "Users" ("id");

ALTER TABLE "Reports" ADD FOREIGN KEY ("comment") REFERENCES "Comments" ("id");
