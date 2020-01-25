module GoMicro

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

require (
	github.com/coreos/etcd v3.3.16+incompatible // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.11.1
	github.com/micro/go-plugins v1.3.0
	github.com/micro/micro v1.11.1
	go.etcd.io/etcd v3.3.16+incompatible
	golang.org/x/net v0.0.0-20190930134127-c5a3c61f89f3
	google.golang.org/grpc v1.24.0
)
