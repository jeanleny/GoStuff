all :
	docker compose -f ./API/docker-compose.yml up --build


down :
	docker compose -f ./API/docker-compose.yml down

clean :
	docker compose -f ./API/docker-compose.yml down -v

fclean : clean
	docker system prune --all --force
