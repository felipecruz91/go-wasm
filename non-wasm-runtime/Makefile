build:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

docker-build:
	docker buildx build -t felipecruz/go-wasm .

docker-run:
	docker run --rm -it -p 8080:80 felipecruz/go-wasm

.PHONY: build
