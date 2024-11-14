generate_grpc_code:
    protoc \
    --go_out=go_service \
    --go_opt=paths=source_relative \
    --go-grpc_out=go_service \
    --go-grpc_opt=paths=source_relative \
    go_service.proto