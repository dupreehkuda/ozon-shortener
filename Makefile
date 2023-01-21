.PHONY: compose
compose:
	docker-compose up

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans

.PHONY: build
build:
	docker-compose down --remove-orphans
	docker-compose build

.PHONY: build_memory
build_memory:
	docker build -t ozon .

.PHONY: run_memory
run_memory:
	docker run -p 8080:8080 -p 9000:9000 ozon -m

.PHONY: test
test:
	go test -cover ./...