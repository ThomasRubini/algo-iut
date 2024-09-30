FROM golang:latest AS build

WORKDIR /build

# Download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy code
COPY internal ./internal
COPY main.go .

# Compile
RUN GOOS=js GOARCH=wasm go build -o main.wasm .

RUN go env GOROOT
RUN go env GOROOT

# Raw web content
FROM scratch AS raw
COPY --from=build /build/main.wasm ./
COPY --from=build /usr/local/go/misc/wasm/wasm_exec.js .
COPY web/index.html ./

# Nginx image to access web content
FROM nginx
COPY --from=raw . /usr/share/nginx/html
