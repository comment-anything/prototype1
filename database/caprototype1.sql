CREATE SCHEMA IF NOT EXISTS "Users";

DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname='AccessLevels')
			THEN
				CREATE TYPE "Users"."AccessLevels" AS ENUM (
					'Poster',
					'DomainModerator',
					'GlobalModerator',
					'Administrator'
				);
			END IF;
END$$;

CREATE TABLE IF NOT EXISTS "Users"."Users"
		(
			"ID" integer NOT NULL,
			"Username" varchar(30) NOT NULL,
			"CreatedAt" timestamp,
			"LastLogin" timestamp,
			"Email" varchar(30),
			"Access" "Users"."AccessLevels",
			"CountryCode" integer,
			"PasswordHash" varchar(30) NOT NULL,
			"SessionID" integer,
			CONSTRAINT "Users_pkey" PRIMARY KEY ("ID")
		);

CREATE TABLE IF NOT EXISTS "Sessions" (
  "ID" SERIAL PRIMARY KEY,
  "CreatedAt" timestamp,
  "Expires" timestamp
);

CREATE TABLE IF NOT EXISTS "DomainModerationAssignments" (
  "ModeratorID" int,
  "Domain" varchar,
  "isMasterOfDomain" boolean
);

CREATE TABLE IF NOT EXISTS "Domains" (
  "ID" varchar UNIQUE PRIMARY KEY NOT NULL
);

CREATE TABLE IF NOT EXISTS "DomainPaths" (
  "ID" varchar PRIMARY KEY,
  "Domain" varchar
);

CREATE TABLE IF NOT EXISTS "Comments" (
  "ID" SERIAL PRIMARY KEY,
  "PostedAt" domainPaths,
  "CreatedAt" timestamp,
  "Poster" int,
  "Content" varchar,
  "Respectables" int
);

CREATE TABLE IF NOT EXISTS "Countries" (
  "Code" int PRIMARY KEY,
  "Name" varchar,
  "ContinentName" varchar
);

ALTER TABLE "Users" ADD FOREIGN KEY ("CountryCode") REFERENCES "Countries" ("Code");

ALTER TABLE "Users" ADD FOREIGN KEY ("Session") REFERENCES "Sessions" ("ID");

ALTER TABLE "DomainModerationAssignments" ADD FOREIGN KEY ("ModeratorID") REFERENCES "Users" ("ID");

ALTER TABLE "DomainModerationAssignments" ADD FOREIGN KEY ("Domain") REFERENCES "Domains" ("ID");

ALTER TABLE "DomainPaths" ADD FOREIGN KEY ("Domain") REFERENCES "Domains" ("ID");

ALTER TABLE "Comments" ADD FOREIGN KEY ("PostedAt") REFERENCES "DomainPaths" ("ID");

ALTER TABLE "Comments" ADD FOREIGN KEY ("Poster") REFERENCES "Users" ("ID");
