FROM --platform=$BUILDPLATFORM tinygo/tinygo:0.26.0 AS build
WORKDIR /build
COPY main.go .
RUN GOOS=js GOARCH=wasm go generate main.go
RUN tinygo build -o main.wasm --target wasm --no-debug main.go

FROM scratch
COPY --from=build /build/wasm_exec.js ./wasm_exec.js
COPY --from=build /build/main.wasm ./main.wasm
ENTRYPOINT [ "main.wasm" ]
