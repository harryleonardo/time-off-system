BINARY=engine

.PHONY: clean run test init domain fmt

fmt:
	gofmt -s -w .

# init project using hygen
init:
	hygen init new && gofmt -s -w .

# create new domain using hygen
domain:
	hygen domain new && gofmt -s -w .

# run golang service 
start:
	go run app/main.go

# run dev docker-compose
run:
	docker-compose up -d --build

# stop docker compose
stop:
	docker-compose down

# build docker image
docker:
	docker build -t $(basename $PWD):latest .

# run golang test
test:
	go test ./...

# clean binary build of this go service
clean: 
	if [ -f ${BINARY} ]; then rm ${BINARY} ; fi

vendor:
	go mod vendor

engine:
	go build -o engine app/main.go

# check swagger
check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

# create swagger docs
swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./swagger.json --scan-models

# serve swagger
serve-swagger: check-swagger
	swagger serve -F=swagger swagger.json

.PHONY: start run stop docker test vendor engine 

db_create:
	cd ~/go/src/github.com/time-off-system/database/migrations/ && goose create ${name} sql

db_up:
	cd ~/go/src/github.com/time-off-system/database/migrations/ && goose mysql "root:@/time-off-dev?parseTime=true" up

db_seeder:
	cd ~/go/src/github.com/time-off-system/database/seeder/ && goose create ${name} sql

db_seeder_up:
	cd ~/go/src/github.com/time-off-system/database/seeder/ && goose mysql "root:@/time-off-dev?parseTime=true" up
