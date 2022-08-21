## up: 执行 docker-compose
up:
	cd ./resource && docker-compose up -d

## down: 停止 docker-compose
down:
	cd ./resource && docker-compose down

## 进入mongo
mongo:
	docker exec -it mongo-restaurant bash

## admin: 启动 manage api
admin:
	go run ./cmd/manage/*.go