# Obviously don't use this in production
.PHONY: postgresInitDev postgres createDevDb dropDevDb

postgresInitDev:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine

postgres:
	docker exec -it postgres16 psql

createDevDb:
	docker exec -it postgres16 createdb --username=root --ownder=root stranger-danger

dropDevDb:
	docker exec -it postgres16 dropdb stranger-danger