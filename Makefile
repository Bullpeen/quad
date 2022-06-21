lint:
	golangci-lint run

build:
	docker build .

build-dev:
	go mod vendor
	docker build --build-arg COMMIT=dev-$(shell date +%Y%m%d%H%M) -t jirwin/quad:dev .
	docker push jirwin/quad:dev
	kubectl rollout restart deployment/nostromo-bullpeen-quadlek

.PHONY: lint build-dev build
