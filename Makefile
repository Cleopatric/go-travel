build:
	make stop
	docker-compose -f docker-compose.yaml up -d --build
run:
	docker-compose -f docker-compose.yaml up -d
stop:
	docker-compose -f docker-compose.yaml down
logs:
	docker-compose -f docker-compose.yaml logs -f