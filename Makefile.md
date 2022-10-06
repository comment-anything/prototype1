# Makefile documentation

This document is intended to describe the usage of the makefile.

Check the [commands table.](#commands-table)

The syntax is `make <command>` from a terminal in the root directory of this project.

E.g, `make test_server` will run the server tests.


# Commands Table

| make command   | purpose               | description                                                                               |
|----------------|-----------------------|-------------------------------------------------------------------------------------------|
| dependencies   | preparing environment | Uses docker to pull the alpine postgres image and gets go dependencies.                   |
| test_generated | testing               | Tests sqlc generated code.                                                                |
| test_server    | testing               | Tests the server.                                                                         |
| stopdb         | database              | Stops the running postgres db.                                                            |
| startdb        | database              | Starts the postgres container.                                                            |
| sqlc           | code generation       | Generates code from queries in database/queries                                           |
| serve          | server                | Starts the go HTTP Server                                                                 |
| rmpostgres     | database              | Stops the postgres container and removes it. Could be dangerous.                          |
| psql           | database              | Accesses the postgres shell in the running container.                                     |
| postgres       | database              | Starts the postgres container, passing in all necessary configuration from .env           |
| migrateup      | database              | Migrates the database up a level, advancing the schema.                                   |
| migratedown    | database              | Migrates the database down a level, regressing the schema and possibly causing data loss. |
| initmigrate    | database              | Creates an initial migrate file. (Probably deprecated)                                    |
| dropdb         | database              | Drops the db specified by the .env file from the running postgres server.                 |
| dbshell        | database              | Accesses the _linux_ shell on the postgres container.                                     |
| createdb       | database              | Creates a db in the running postgres docker according to the environment settings.        |