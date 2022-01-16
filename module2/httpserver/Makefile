
build:
	mkdir -p bin/ && go build -x -v  -o ./bin/httpserver ./...


docker-build:
	docker build -t ducknight/httpserver:$(TAG) .

docker-push:
	docker push ducknight/httpserver:$(TAG)