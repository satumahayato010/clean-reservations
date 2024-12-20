build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

test:
	docker-compose run --rm app go test -v ./...

swagger-gen:
	docker-compose run --rm app swag init
