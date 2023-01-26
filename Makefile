CONTAINER_NAME=mysql_user
DB_NAME=userDB
# DB_URL=mysql://localhost:3306/${DB_NAME}

docker_mysql:
	docker run --platform linux/amd64 -p 3306:3306 --env MYSQL_DATABASE=${DB_NAME} --env MYSQL_ROOT_PASSWORD=root --name ${CONTAINER_NAME} -d mysql

sqlc:
	sqlc generate
	
.PHONY: docker_mysql migrateup