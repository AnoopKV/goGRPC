# Note
Make sure you register grpcServer instance to reflection as given below:
reflection.Register(grpcServer)
otherwise neither postman nor grpcurl detects service/method

## Installation
to run with grpcurl, install this using go install
```bash 
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

## Testing
### Via grpcurl api:
```bash 
grpcurl --plaintext -d '{"id":"1", "status":"Pending", "orderItems":[{"code":"","name":"","unitPrice":0,"quantity":0}]}' localhost:5001 OrderService.CreateOrder
```

### Via Postman
Create New gRPC request, give server name localhost:5001, service name automatically it fetch if your grpc server is registered with reflection.
