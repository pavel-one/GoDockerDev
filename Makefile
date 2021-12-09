up:
	docker-compose up -d
down:
	docker-compose down
exec:
	docker-compose exec app bash
exec.root:
	docker-compose exec -u root app bash
log:
	docker-compose logs -f app