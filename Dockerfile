FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o status-provider server.go
RUN chmod +x status-provider
EXPOSE 8080
CMD [ "./status-provider" ]
