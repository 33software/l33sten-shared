#auth service proto gen
protoc -I protos/auth/proto --go_out=protos/auth/gen --go_opt=paths=source_relative --go-grpc_out=protos/auth/gen --go-grpc_opt=paths=source_relative auth.proto

#user service proto gen 
protoc -I protos/user/proto --go_out=protos/user/gen --go_opt=paths=source_relative --go-grpc_out=protos/user/gen --go-grpc_opt=paths=source_relative user.proto

#playlist service proto gen 
protoc -I protos/playlist/proto --go_out=protos/playlist/gen --go_opt=paths=source_relative --go-grpc_out=protos/playlist/gen --go-grpc_opt=paths=source_relative playlist.proto

#track service proto gen
protoc -I protos/track/proto --go_out=protos/track/gen --go_opt=paths=source_relative --go-grpc_out=protos/track/gen --go-grpc_opt=paths=source_relative track.proto

#streaming service (?) proto gen
protoc -I protos/streaming/proto --go_out=protos/streaming/gen --go_opt=paths=source_relative --go-grpc_out=protos/streaming/gen --go-grpc_opt=paths=source_relative streaming.proto


