.SITE: run_db, migrations

run_db:
		docker run --name=apps -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=expenditure -p 5432:5432 -d --rm postgres 
go_migrate:
		migrate -path ./migrations -database postgres://postgres:qwerty@127.0.0.1:5432/expenditure?sslmode=disable up