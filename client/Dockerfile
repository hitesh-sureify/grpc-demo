FROM golang

WORKDIR /api

COPY . .

ENV GRPC_SRV_ADDR employee-service:50052

RUN go get .

ENTRYPOINT go run client.go

EXPOSE 8000