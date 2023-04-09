
init:
	docker-compose up --build -d
	make exec
up:
	docker-compose up -d
	make exec
run:
	go run *.go

exec:
	docker exec -it c-b bash