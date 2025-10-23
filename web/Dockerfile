FROM golang:latest AS build

WORKDIR /build

# Download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy code
COPY internal ./internal
COPY main_wasm.go .

# Compile
RUN GOOS=js GOARCH=wasm go build -o main.wasm .

# Raw web content
FROM scratch AS raw
COPY --from=build /build/main.wasm ./
COPY --from=build /usr/local/go/lib/wasm/wasm_exec.js .
COPY web/content ./

# Nginx image to access web content
FROM nginx
COPY --from=raw . /usr/share/nginx/html
