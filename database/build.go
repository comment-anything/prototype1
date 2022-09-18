package database

import (
	"database/sql"
	"fmt"
	"log"
)

func BuildFromSQL() {
	postgres := DBConnector.Connect()
	buildUsers(postgres)
	postgres.Close()
}

func buildUsers(postgres *sql.DB) {
	_, err := postgres.Exec(`CREATE SCHEMA IF NOT EXISTS "Users";`)
	checkBuildError(err, postgres, "Users Schema")
	/**
		I think Users.AccessLevels, a Postgres Eum may be wrong. It probably makes more sense to just store the raw enum value in the database as a byte and only enum it in the Go code? I may experiment more with the postgres enum data type.

		Also, Creating If Not Exists for an enum is a weird query.

		-klm127
	**/
	_, err = postgres.Exec(`
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
		END$$;`)
	checkBuildError(err, postgres, "Users.AccessLevels Enum")
	_, err = postgres.Exec(`
		CREATE TABLE IF NOT EXISTS "Users"."Users"
		(
			"Username" character varying(30) COLLATE pg_catalog."default" NOT NULL,
			"Email" character varying(30) COLLATE pg_catalog."default",
			"Access" "Users"."AccessLevels",
			"CountryCode" integer,
			"PasswordHash" character varying(30) COLLATE pg_catalog."default" NOT NULL,
			"SessionID" integer,
			"CreatedAt" integer,
			"LastLogin" integer,
			"ID" integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
			CONSTRAINT "Users_pkey" PRIMARY KEY ("ID")
		);
	`)
	checkBuildError(err, postgres, "Users Table")
}

func checkBuildError(err error, db *sql.DB, error_on string) {
	if err != nil {
		db.Close()
		fmt.Printf("Problem in the the build process on %s.\n", error_on)
		log.Fatalf(err.Error())
	}
}
