protogen:
	@protoc \
		--go_out=proto/gen/orders \
		--go_opt=paths=source_relative \
		--go-grpc_out=proto/gen/orders \
		--go-grpc_opt=paths=source_relative \
		protos/*.proto

protogenV2:
	@protoc \
		--go_out=protos \
		--go-grpc_out=protos \
		protos/*.proto
