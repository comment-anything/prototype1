include .env

dependencies:
	docker pull ${DB_IMAGE}
	docker pull kjconroy/sqlc
	go get .
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

sqlc:
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

serve:
	go run .
	
test_server:
	go test ./server -cover 

test_generated:
	go test ./database/generated -cover

doc:
	golds -gen -s -nouses -wdpkgs-listing=promoted -source-code-reading=rich -dir=docs -render-doclinks ./main.go
	golds -dir=docs


