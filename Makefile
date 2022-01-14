up:
	docker-compose up -d
down:
	docker-compose down
ps:
	docker-compose ps
reboot: down up;

rebuild:
	docker-compose build --no-cache
logs:
	docker-compose logs
exec:
	docker-compose exec ${name} sh
genmock:
	genmock -source=${source} destination=${mock} -package=${pkg}
