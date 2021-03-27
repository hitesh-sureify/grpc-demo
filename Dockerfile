FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV GRPC_SRV_ADDR :50052
ENV DB_DRIVER mysql
ENV DB_USER hitesh
ENV DB_PASS 68c#sistEdgCD4
ENV DB_NAME MYSQLTEST
ENV DB_HOST 172.17.0.1:3306

RUN go build

EXPOSE 50052

CMD ["./grpc-demo"]