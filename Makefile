lint:
	golangci-lint run

build:
	docker build .

build-dev:
	go mod vendor
	docker build --build-arg COMMIT=dev-$(shell date +%Y%m%d%H%M) -t jirwin/quad:dev .
	docker push jirwin/quad:dev
	kubectl rollout restart deployment/nostromo-bullpeen-quadlek

upgrade-quadlek:
	go get -u github.com/jirwin/quadlek@main
	go mod tidy
	go mod vendor

.PHONY: lint build-dev build
