module github.com/ArkArk/dcp

go 1.13

require (
	github.com/containerd/containerd v1.3.2 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.0 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/urfave/cli/v2 v2.1.1
	google.golang.org/grpc v1.26.0 // indirect
)

// ref. https://github.com/moby/moby/issues/37683#issuecomment-415101262
replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20191113042239-ea84732a7725
