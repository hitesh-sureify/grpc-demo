FROM golang

WORKDIR /api

COPY . .

ENV GRPC_SRV_ADDR emp-srv:50052

RUN go get .

ENTRYPOINT go run client.go

EXPOSE 8000