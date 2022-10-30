FROM --platform=$BUILDPLATFORM golang:1.19-alpine AS build
WORKDIR /build
COPY main.go .
RUN GOOS=js GOARCH=wasm go build -o main.wasm main.go

FROM scratch
COPY --from=build /build/main.wasm ./main.wasm
ENTRYPOINT [ "main.wasm" ]
