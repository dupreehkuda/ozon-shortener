.PHONY: compose
compose:
	docker-compose up

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans

.PHONY: build_memory
build_memory:
	docker build -t ozon .

.PHONY: run_memory
run_memory:
	docker run -p 8080:8080 ozon -m