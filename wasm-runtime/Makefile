docker-build:
	docker buildx build \
	--platform=wasi/wasm32 \
	-t felipecruz/go-wasm .

docker-run:
	docker run \
	--rm \
	--runtime=io.containerd.wasmedge.v1 \
	--platform=wasi/wasm32 \
	felipecruz/go-wasm
