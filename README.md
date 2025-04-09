# envoy-grpc
This is a sample envoy grpc service compatible to use with envoy filters.

References to test sample envoy grpc service
https://github.com/salrashid123/envoy_grpc_decode
https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md

Commands:
1. make // This generates a folder named demo and creates contents of protobufs in it.
2. go run main.go // Listens on port 9000 for requests. Optional:
3. sh build.sh // creates a binary of this sample service
4. docker tag name image-name
5. docker push image-name
