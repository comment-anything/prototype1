include .env

dependencies:
	docker pull ${DB_IMAGE}
	go get .

windowsdevdependencies:
 # you can look up how to install scoop through powershell. Migrate is used for db migrations.
	scoop install migrate

postgres:
	docker run --name ${DB_CONTAINER_NAME} -p ${DB_HOST_PORT}:${DB_CONTAINER_PORT} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d ${DB_IMAGE}

startdb:
	docker start ${DB_CONTAINER_NAME}

stopdb:
	docker stop ${DB_CONTAINER_NAME}

rmpostgres:
	docker stop ${DB_CONTAINER_NAME}
	docker rm ${DB_CONTAINER_NAME}

psql:
	docker exec -it ${DB_CONTAINER_NAME} psql -U ${DB_USER} ${DB_NAME}

dbshell:
	docker exec -it ${DB_CONTAINER_NAME} /bin/sh

