FROM golang:1.19-alpine AS build
WORKDIR /build
COPY main.go .
RUN GOOS=js GOARCH=wasm go generate main.go
RUN GOOS=js GOARCH=wasm go build -o main.wasm main.go

FROM nginx:1.23
WORKDIR /usr/share/nginx/html/
COPY web/index.html ./index.html
COPY --from=build /build/main.wasm ./main.wasm
COPY --from=build /build/wasm_exec.js ./wasm_exec.js
