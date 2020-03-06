build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

build-local:
	go build ./main.go && ./main

docker-build-local:
	docker build -t dns-convert:latest .

docker-build-push:
	docker build -t dockerhub.com/<name>/dns-convert:latest .
	docker push dockerhub.com/<name>/dns-convert:latest

create-csv:
	./main | tr -d "[:blank:]" > entries.csv