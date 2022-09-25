
This file documents terminology used in code and documentation around this project. Additionally, it describes the contents of directory and the nature of each file in the root directory of the project source code.


# Contents

  1. [External Terms](#external-terms)
  2. [Internal Terms](#internal-terms)
  3. [Directories](#directories)
  4. [Root Files](#root-files)


# External Terms

> todo: write descriptions, paraphrase definitions into modes that apply most to our project.

External terms are technical terms for pre-existing concepts and technologies used by this project.

| Name | Description | Aliases | Reference |
| ---- | ----------- | ------- | ----- |
| Automatic Programming | "The act of generating source code based on an ontological model such as a template." | Code Generation
| Container | "An isolated user space instance executed by an Operating System which allows virtualization." Each Container runs its own operating system and accesses its own virtual drive. | Virtual Machines, Virtual Private Servers, Partitions, Virtual Kernels, Virtual Environments, Zones | [Network World](https://www.networkworld.com/article/2226996/software-containers--used-more-frequently-than-most-realize.html)
| Database ||
| Discord ||
| Docker | |
| Environment Variable ||
| Git ||
| GitHub ||
| Go | |
| HTTP | | 
| JavaScript || 
| PostgresSQL | | 
| HTML | Hypertext ....|
| Image ||
| Makefile ||
| Markdown ||
| Migration ||
| Static ||
| SqlC ||
| Server ||
| Template ||
| View ||
| Virtualization | The use of software to create an abstraction layer over computer hardware that allows the hardware of a single computer to be divided into mutliple virtual computers, called Containers. |...| [IBM](https://www.ibm.com/cloud/learn/virtualization-a-complete-guide) |

# Internal Terms

> Todo: Add terms. User, Moderator, Administrator, Comment... all the tables and such should get a term or two each.

Internal terms are terms defined exclusively for this project's lexicon. 

| Name  | Description |
| ----- |  ----------- |
| database | The go package database used by this module to communicate with Postgres. |
| Database | The running database instance active on another port that GoServer communicates with. 
| GoServer | The running go module that handles http requests, responses, database communications, and the rest of the business logic.
| Postgres | Alias for database because Postgres is the current database of choice.
| User |


## Directories

Directories are organizational units for files. Each directory contains files used by GoServer to execute some related functionality.

| Name  | Description |
| ----- |  ----------- |
| `/database` | Contains go package `database` and migration files. |
| `/database/migration` | Contains database migrations to enable iterative schema changes to the database for use with the [`golang-migrate`](https://github.com/golang-migrate) command line tool.
| `/database/query/` | SQL files describing queries that are read by [`sqlc`](https://docs.sqlc.dev/en/stable/) to generate go code.
| `/database/sqlc/` | Generated `sqlc` code providing models, methods, and validators.
| `/server/` | Contains go package `server`. |
| `/static/` | Files served statically from the server, such as images, stylsheets, and front-end Javascript, to be used by html across the application.
| `/templates/` | Files to be used with go html templates to render struct data. 
| `/views` | Contains .html files to be served to a user. |

## Root Files

Root Files live directly under the root node folder of the file tree. They document the project and configure the build and run process. 

| Name  | Description |
| ----- |  ----------- |
| `/.env` | Contains newline separated key-value pairs corresponding to environment variables used by both the database and go servers.
| `/.gitignore` | A newline separated list of paths to exclude from .git tracking.
| `/bserve.bat` | A windows script for starting the go server *and* opening the browser from the command line.
| `/contributing.md`| A description of tasks that need to be accomplished and contribution guides.
| `/glossary.md` | A file defining the project namespace domain.
| `/go.mod` | 3rd party go module dependencies that this go module depends on, added automatically by `go get`.
| `/init_schema.sql` | The initial database schema. Just for rerence; all database states are stored in `database/migration`.
| `/main.go` | The entry point for the go program.
| `/Makefile` | A file containing commands for getting dependencies, initializing, building, and running the database utilizing the environment variables and adding all appropriate options to docker. 
| `/readme.md` | A description of this project that serves as a landing point on github.
| `/sqlc.yaml` | A configuration file for the [`sqlc`](https://docs.sqlc.dev/en/stable/) code generating process.

