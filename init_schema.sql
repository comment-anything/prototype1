CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_login" timestamptz NOT NULL DEFAULT (now()),
  "access_level" int8 NOT NULL DEFAULT 0
);

CREATE TABLE "DomainBans" (
  "user" bigint,
  "banned_from" varchar,
  "banned_by" bigint,
  "banned_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Comments" (
  "id" bigserial PRIMARY KEY,
  "pathid" bigint NOT NULL,
  "author" bigint NOT NULL,
  "content" varchar NOT NULL,
  "response_type" int8 NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "parent" bigint,
  "hidden" boolean DEFAULT false,
  "hidden_by" bigint,
  "hidden_at" timestamptz
);

CREATE TABLE "Domains" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL
);

CREATE TABLE "Paths" (
  "id" bigserial PRIMARY KEY,
  "domain" varchar,
  "path" varchar
);

CREATE TABLE "Moderators" (
  "id" bigserial PRIMARY KEY,
  "domain" varchar NOT NULL,
  "user" bigint NOT NULL,
  "level" int8 NOT NULL DEFAULT 1,
  "granted_at" timestamptz NOT NULL DEFAULT (now()),
  "granted_by" bigint NOT NULL
);

CREATE TABLE "GlobalModerators" (
  "id" bigserial PRIMARY KEY,
  "user" bigint,
  "granted_at" timestamptz NOT NULL DEFAULT (now()),
  "granted_by" bigint NOT NULL
);

CREATE TABLE "Admins" (
  "id" bigserial PRIMARY KEY,
  "user" bigint
);

CREATE TABLE "RemovedComments" (
  "id" bigserial PRIMARY KEY,
  "pathid" bigint NOT NULL,
  "author" bigint NOT NULL,
  "content" varchar NOT NULL,
  "response_type" int8 NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL,
  "parent" bigint,
  "removed_at" timestamptz NOT NULL DEFAULT (now()),
  "removed_by" bigint NOT NULL
);

CREATE INDEX ON "Users" ("username");

CREATE INDEX ON "Users" ("email");

CREATE INDEX ON "DomainBans" ("user");

CREATE INDEX ON "DomainBans" ("banned_by");

CREATE INDEX ON "DomainBans" ("banned_from");

CREATE INDEX ON "Comments" ("author");

CREATE INDEX ON "Comments" ("pathid");

CREATE UNIQUE INDEX ON "Paths" ("domain", "path");

CREATE INDEX ON "RemovedComments" ("author");

COMMENT ON COLUMN "Users"."password" IS 'Must be encrypted';

COMMENT ON COLUMN "Comments"."response_type" IS 'e.g, joke, fact, diatribe, opinion';

COMMENT ON COLUMN "RemovedComments"."response_type" IS 'e.g, joke, fact, diatribe, opinion';

ALTER TABLE "DomainBans" ADD FOREIGN KEY ("user") REFERENCES "Users" ("id");

ALTER TABLE "DomainBans" ADD FOREIGN KEY ("banned_from") REFERENCES "Domains" ("id");

ALTER TABLE "DomainBans" ADD FOREIGN KEY ("banned_by") REFERENCES "Users" ("id");

ALTER TABLE "Comments" ADD FOREIGN KEY ("pathid") REFERENCES "Paths" ("id");

ALTER TABLE "Comments" ADD FOREIGN KEY ("author") REFERENCES "Users" ("id");

ALTER TABLE "Comments" ADD FOREIGN KEY ("parent") REFERENCES "Comments" ("id");

ALTER TABLE "Comments" ADD FOREIGN KEY ("hidden_by") REFERENCES "Users" ("id");

ALTER TABLE "Paths" ADD FOREIGN KEY ("domain") REFERENCES "Domains" ("id");

ALTER TABLE "Moderators" ADD FOREIGN KEY ("domain") REFERENCES "Domains" ("id");

ALTER TABLE "Moderators" ADD FOREIGN KEY ("user") REFERENCES "Users" ("id");

ALTER TABLE "Moderators" ADD FOREIGN KEY ("granted_by") REFERENCES "Users" ("id");

ALTER TABLE "GlobalModerators" ADD FOREIGN KEY ("user") REFERENCES "Users" ("id");

ALTER TABLE "GlobalModerators" ADD FOREIGN KEY ("granted_by") REFERENCES "Users" ("id");

ALTER TABLE "Admins" ADD FOREIGN KEY ("user") REFERENCES "Users" ("id");

ALTER TABLE "RemovedComments" ADD FOREIGN KEY ("pathid") REFERENCES "Paths" ("id");

ALTER TABLE "RemovedComments" ADD FOREIGN KEY ("author") REFERENCES "Users" ("id");

ALTER TABLE "RemovedComments" ADD FOREIGN KEY ("parent") REFERENCES "Comments" ("id");

ALTER TABLE "RemovedComments" ADD FOREIGN KEY ("removed_by") REFERENCES "Users" ("id");
