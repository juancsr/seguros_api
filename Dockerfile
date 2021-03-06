FROM golang:1.16-alpine3.13
COPY . /app/
WORKDIR /app/
RUN go get -v -u ./...
ENTRYPOINT [ "go", "run", "main.go" ]