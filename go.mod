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
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	google.golang.org/grpc v1.22.1
)

// replace github.com/micro/go-micro => ../vendor/github.com/micro/go-micro
