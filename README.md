# grpc-template
### Pre-Requisite
- protoc must be installed
### Step 1 : Define communication protocol
- Define message types and services in a **<your-project>.proto** file. Example format : [proto-file](https://github.com/hitesh-sureify/grpc-template/blob/main/proto/employee.proto)
- Compile the **<your-project>.proto** file using a proto compiler **protoc** using command format : protoc --proto_path=<proto-file-path> --go_out=<output-dir>.
- After successfull complilation, a **<your-project>.pb.go** file will get created.

### Step 2 : Implementation of service interface
- In '.go' file where service needs to be implemented, import the auto-generated **<your-project>.pg.go** file. Please see : [service-implementation](https://github.com/hitesh-sureify/grpc-template/blob/main/server.go).
- Implement <your-project>ServiceServer interface here. 

### Step 3 : Create a client and implement client interface
- In '.go' file where client needs to be implemented, import the auto-generated **<your-project>.pg.go** file.
  Please see : 
  1. [rest-client-implementation](https://github.com/hitesh-sureify/grpc-template/blob/main/client/client.go).
  2. [cmd-client-implementation](https://github.com/hitesh-sureify/grpc-template/blob/main/client/cmd/client.go)
- Implement <your-project>ServiceClient interface here. 