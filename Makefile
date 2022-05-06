up:
	docker-compose up -d
down:
	docker-compose down
reload:
	docker-compose restart
remove:
	docker-compose down --rmi all --volumes --remove-orphans
logs:
	docker-compose logs
shell:
	docker-compose exec api bash
reup:
	make remove
	make up