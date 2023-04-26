migrate_up:
	migrate -path db/migration -database "postgres://ojrgbwbb:lHIin4nHhnGhSsBi1zz4EzQT6RutIm_9@lallah.db.elephantsql.com/ojrgbwbb?sslmode=disable" -verbose up

migrate_force:
	migrate -path db/migration -database "postgres://ojrgbwbb:lHIin4nHhnGhSsBi1zz4EzQT6RutIm_9@lallah.db.elephantsql.com/ojrgbwbb?sslmode=disable" force $(version)

migrate_down:
	migrate -path db/migration -database "" -verbose down

run:
	go build e-commerce.go && ./e-commerce

gotidy:
	go mod tidy

goinit:
	go mod init

swag:
	swag init -g e-commerce.go -ot go,yaml 

migrate_init:
	migrate create -ext sql -dir db/migration -seq init_schema

launch:
	flyctl launch

docker_init:
	docker build -t everybody8/e-commerce:v$(version) .

docker_push:
	docker push everybody8/e-commerce:v$(version)

deploy:
	flyctl deploy

test_repo:
	go test ./tests/repo_test -v

test_service:
	go test ./tests/service_test -v

test_handler:
	go test ./tests/handler_test -v

add_commit:
	git add . && git commit -m "$(message)"