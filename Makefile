migrateup:
	migrate -path db/migration -database "mysql://sql8597858:qneWhfChgE@tcp(sql8.freemysqlhosting.net:3306)/sql8597858" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://sql8597858:qneWhfChgE@tcp(sql8.freemysqlhosting.net:3306)/sql8597858" -verbose down

run:
	go build e-commerce.go && ./e-commerce

gotidy:
	go mod tidy

goinit:
	go mod init

swag:
	swag init

migrate_init:
	migrate create -ext sql -dir db/migration -seq init_schema

launch:
	flyctl launch

docker_init:
	docker build -t everybody8/benny-foods:v$(version) .

docker_push:
	docker push everybody8/benny-foods:v$(version)

deploy:
	flyctl deploy