FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV GRPC_SRV_ADDR :50052
ENV DB_USER postgres
ENV DB_PASS postgres
ENV DB_NAME testdb
ENV DB_HOST localhost
ENV DB_PORT 5432

RUN go build

EXPOSE 50052

CMD ["./grpc-demo"]