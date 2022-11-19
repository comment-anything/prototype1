
## Prototype1

This is a minimum proof of concept to:
 - Test feasibility
 - Test docker, make, and migrate
 - Test generating model code with sqlc
 - Test implementing middleware
 - Test auth
## Clone the Repo to the correct location in %GOPATH%

- Navigate to your %GOPATH% directory.
- By default, on Windows, it is `C:/Users/%Username%/go/`
- Already in that directory will be `bin` and `pkg` if you have installed go with default settings.
- Create the directory `src/github.com/`
- Enter that directory `cd src/github.com`
- Clone the repo `git clone github.com/comment-anything/prototype1.git`
- This module will be saved to, for example `C:/Users/karlm/go/src/github.com/comment-anything/prototype1`

Read more about [Gopath here](https://golangr.com/what-is-gopath/). You can also set up custom workspaces but if you are not going to write a lot of conflicting Go code, using the default go path is probably the way to go.

## Set up environment variables

Create a file in the root called `.env`. This will hold private information used across the program. Do not upload this file to github.

Example .env file for default database:

```
# DB_IMAGE is the docker image that the database will be built from.
DB_IMAGE=postgres:14.5-alpine
# DB_CONTAINER_NAME is the name of the container in docker.
DB_CONTAINER_NAME=923postgres
# DB_CONTAINER_PORT is the port the container will be listening on in its environment. It will be mapped to the value of DB_HOST_PORT on the host device. These values can be the same but if you already have postgres installed on your computer like I do, you may want to map to a different port.
DB_CONTAINER_PORT=5432

DB_HOST=localhost
# The port that Go will listen to database with(for port mapping)
DB_HOST_PORT=5433
DB_USER=root
DB_PASSWORD=dbsuperuser991
DB_DATABASE_NAME=comm-anything

# SERVER_PORT is the port that Go will be served on.
SERVER_PORT=3000
SERVER_LOG_ALL=true
```

## Install additional tools

The easiest way to deploy is with [Docker](https://www.docker.com/products/docker-desktop/). 

Follow the instructions on docker.com to install Docker on your machine. You may have to configure your BIOS to enable virtual machines.

Docker allows the deployment of isolated container virtual machines. Both the Postgres database server and the Go back end server will run in Docker containers.

If you are on Windows you will also need *make*. You can get make by installing [choclatey](https://chocolatey.org/install) and running `choco install make`.

## Building, Running

Use the `make` commands to get other dependencies and run the servers.

```shell
make dependencies
make postgres
```

You can then use make commands to access the database server
`make psql` will open the psql cli on the postgres docker container
`make dbshell` will open up the linux shell on the postgres docker container

And you can build the database tables.

```shell
make createdb
make migrateup
```

Finally, run the server. It will listen on the port configured in the .env file. If you used the example .Env file posted above, that's port 3000. After running the server, you can navigate to the url `localhost:3000` to access the test pages.

```

```



+ ChangeUserPassword(ctx context.Context, arg ChangeUserPasswordParams)
+ CreateComment(ctx context.Context, arg CreateCommentParams)
+ CreateDomain(ctx context.Context, id string)
+ CreatePath(ctx context.Context, arg CreatePathParams)
+ CreateUser(ctx context.Context, arg CreateUserParams) User
+ DeleteUser(ctx context.Context, id int64)
+ GetDomain(ctx context.Context, id string) string)
+ GetPath(ctx context.Context, arg GetPathParams) Path
+ GetUserByEmail(ctx context.Context, email string) User
+ GetUserByUserId(ctx context.Context, id int64) User
+ GetUserByUserName(ctx context.Context, username string) User
+ ListUsers(ctx context.Context)