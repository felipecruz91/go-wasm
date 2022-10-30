FROM --platform=$BUILDPLATFORM tinygo/tinygo:0.26.0 AS build
WORKDIR /build
COPY main.go .
RUN tinygo build -o main.wasm -target wasi -wasm-abi=generic --no-debug main.go

FROM scratch
COPY --from=build /build/main.wasm ./main.wasm
ENTRYPOINT [ "main.wasm" ]
