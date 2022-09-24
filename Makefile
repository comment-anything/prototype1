include .env

dependencies:
	docker pull ${DB_IMAGE}
	go get .
	go install github.com/kyleconroy/sqlc/cmd/sqlc@1.15.0
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	

windowsdevdependencies:
	docker pull kjconroy/sqlc

sqlcgenerate:
	docker run --rm -v "$(CURDIR):/src" -w /src kjconroy/sqlc generate

postgres:
	docker run --name ${DB_CONTAINER_NAME} -p ${DB_HOST_PORT}:${DB_CONTAINER_PORT} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d ${DB_IMAGE}

createdb:
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_DATABASE_NAME}

dropdb:
	docker exec -it ${DB_CONTAINER_NAME} dropdb ${DB_DATABASE_NAME}

startdb:
	docker start ${DB_CONTAINER_NAME}

stopdb:
	docker stop ${DB_CONTAINER_NAME}

rmpostgres:
	docker stop ${DB_CONTAINER_NAME}
	docker rm ${DB_CONTAINER_NAME}

initmigrate:
	migrate create -ext sql -dir database/migration -seq init_schema

migrateup:
	migrate -path database/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_HOST_PORT}/${DB_DATABASE_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_HOST_PORT}/${DB_DATABASE_NAME}?sslmode=disable" -verbose down

psql:
	docker exec -it ${DB_CONTAINER_NAME} psql -U ${DB_USER} ${DB_DATABASE_NAME}

dbshell:
	docker exec -it ${DB_CONTAINER_NAME} /bin/sh

