userProto:
	protoc --go_out=. --go-grpc_out=. ./protos/user.proto

roomProto:
	protoc --go_out=. --go-grpc_out=. ./protos/room_service.proto
