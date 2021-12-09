up:
	docker-compose up -d
down:
	docker-compose down
exec:
	docker-compose exec app bash
log:
	docker-compose logs -f app