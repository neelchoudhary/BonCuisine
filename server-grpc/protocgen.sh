protoc recipe.proto --proto_path=api/proto/v1/recipe --go_out=plugins=grpc:pkg/v1/recipe/api
protoc account.proto --proto_path=api/proto/v1/account --go_out=plugins=grpc:pkg/v1/account/api