#auth service proto gen
protoc -I proto proto/auth/auth.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

#user service proto gen 
protoc -I proto proto/user/user.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

#playlist service proto gen 
protoc -I protos/playlist/proto --go_out=protos/playlist/gen --go_opt=paths=source_relative --go-grpc_out=protos/playlist/gen --go-grpc_opt=paths=source_relative playlist.proto

#track service proto gen
protoc -I protos/track/proto --go_out=protos/track/gen --go_opt=paths=source_relative --go-grpc_out=protos/track/gen --go-grpc_opt=paths=source_relative track.proto

#streaming service (?) proto gen
protoc -I protos/streaming/proto --go_out=protos/streaming/gen --go_opt=paths=source_relative --go-grpc_out=protos/streaming/gen --go-grpc_opt=paths=source_relative streaming.proto


