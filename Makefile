#start-docker:
#	docker desktop restart
#create-container:
#	docker run -d --name eduflex-postgres-container -e POSTGRES_PASSWORD=password -e POSTGRES_DB=schooldb -p 5432:5432 -v postgres-data:/var/lib/postgresql postgres
#start-postgres:
#	docker start eduflex-postgres-container
#run-migrations:
#	docker exec -e PGPASSWORD=password eduflex-postgres-container psql -U postgres -d postgres -c "CREATE DATABASE schooldb;"
#	docker run --rm -v "$(CURDIR)/school-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:password@host.docker.internal:5433/schooldb?sslmode=disable" up
#	docker run --rm -v "${PWD}/school-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:password@host.docker.internal:5432/schooldb?sslmode=disable" up
#see-running-containers:
#	docker ps

#run-redis:
#	docker run -d --name redis-container -p 6379:6379 -v redis-data:/data redis:latest
#	docker exec -it redis-container redis-cli

run-app:
#	docker compose up -d --build 
#	docker compose build --no-cache school-service
	docker compose up -d --build school-service

#stop-app:
#	docker compose down