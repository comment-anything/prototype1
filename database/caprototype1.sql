CREATE SCHEMA "Users";

CREATE TYPE "Users"."AccessLevels" AS ENUM (
  'Poster',
  'DomainModerator',
  'GlobalModerator',
  'Administrator'
);

CREATE TABLE "Users" (
  "ID" SERIAL PRIMARY KEY,
  "Username" varchar,
  "CreatedAt" timestamp,
  "LastLogin" timestamp,
  "Email" varchar,
  "Access" Users.AccessLevels,
  "CountryCode" int,
  "PasswordHash" varchar,
  "Session" int
);

CREATE TABLE "Sessions" (
  "ID" SERIAL PRIMARY KEY,
  "CreatedAt" timestamp,
  "Expires" timestamp
);

CREATE TABLE "DomainModerationAssignments" (
  "ModeratorID" int,
  "Domain" varchar,
  "isMasterOfDomain" boolean
);

CREATE TABLE "Domains" (
  "ID" varchar UNIQUE PRIMARY KEY NOT NULL
);

CREATE TABLE "DomainPaths" (
  "ID" varchar PRIMARY KEY,
  "Domain" varchar
);

CREATE TABLE "Comments" (
  "ID" SERIAL PRIMARY KEY,
  "PostedAt" domainPaths,
  "CreatedAt" timestamp,
  "Poster" int,
  "Content" varchar,
  "Respectables" int
);

CREATE TABLE "Countries" (
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
