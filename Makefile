.PHONY : docker_rmi docker_build docker_cont
docker_build:
	docker build -t my_pg_img .

docker_cont:
	docker run -it --name my_pg_cont -p 5432:5432 my_pg_img

docker_connect:
	docker exec -it my_pg_cont psql -U qd -d cardadb

docker_rmi:
	docker rm  -f my_pg_cont
	docker rmi my_pg_img

rebuild: docker_rmi docker_build docker_cont
run:
	go build ./cmd/carda/main.go
	./main

test:
	go clean -testcache
	go test -v ./internal/carda/sqlquery/