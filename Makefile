#start-docker:
#	docker desktop restart

#run-migrations:

#	docker compose up -d school-db
#	docker compose up -d auth-db
#	docker compose up -d student-db
#	docker compose up -d financing-db
#	docker compose up -d loan-db
#	docker compose up -d ledger-db

#	docker run --rm -v "$(CURDIR)/school-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:password@host.docker.internal:5433/schooldb?sslmode=disable" up
#	docker run --rm -v "$(CURDIR)/student-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:postgres@host.docker.internal:5434/studentdb?sslmode=disable" up
#	docker run --rm -v "$(CURDIR)/auth-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:password@host.docker.internal:5435/authdb?sslmode=disable" up
#	docker run --rm -v "$(CURDIR)/financing-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:postgres@host.docker.internal:5436/financingdb?sslmode=disable" up
#	docker run --rm -v "$(CURDIR)/loan-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:postgres@host.docker.internal:5437/loandb?sslmode=disable" up
#	docker run --rm -v "$(CURDIR)/ledger-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:postgres@host.docker.internal:5438/ledgerdb?sslmode=disable" up


#	docker exec -e PGPASSWORD=password eduflex-schooldb-container psql -U postgres -d postgres -c "CREATE DATABASE schooldb;"



#	
#	docker run --rm -v "${PWD}/school-service/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "postgres://postgres:password@host.docker.internal:5432/schooldb?sslmode=disable" up


#see-running-containers:
#	docker ps

#run-redis:
#	docker run -d --name redis-container -p 6379:6379 -v redis-data:/data redis:latest
#	docker exec -it redis-container redis-cli

run-app:
#	docker compose up -d --build 
#	docker compose build --no-cache school-service 
#	docker compose up -d --build school-service 

	docker compose up -d --build api-gateway 
	docker compose up -d --build auth-service
	docker compose up -d --build school-service
	docker compose up -d --build student-service
	docker compose up -d --build financing-service
	docker compose up -d --build loan-service
	docker compose up -d --build ledger-service

#stop-app:
#	docker compose down

# Push project updates to repository

# Push with a default message (optional)
push-quick:
	git add -A
	git commit -m "Quick update: $(shell date '+%Y-%m-%d %H:%M:%S')"
	git push





#create-schooldb-container:
#	docker run -d --name eduflex-schooldb-container -e POSTGRES_PASSWORD=password -e POSTGRES_DB=schooldb -p 5420:5432 -v postgres-data:/var/lib/postgresql postgres

#create-authdb-container:
#	docker run -d --name eduflex-authdb-container -e POSTGRES_PASSWORD=password -e POSTGRES_DB=authdb -p 5421:5432 -v postgres-data:/var/lib/postgresql postgres

#create-studentdb-container:
#	docker run -d --name eduflex-studentdb-container -e POSTGRES_PASSWORD=password -e POSTGRES_DB=studentdb -p 5422:5432 -v postgres-data:/var/lib/postgresql postgres

#start-postgres:
#	docker start eduflex-studentdb-container
#	docker start eduflex-schooldb-container
#	docker start eduflex-authdb-container