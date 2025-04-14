module github.com/luskits/luscsi

go 1.23.6

toolchain go1.24.0

require (
	github.com/container-storage-interface/spec v1.11.0
	github.com/kubernetes-csi/csi-lib-utils v0.21.0
	golang.org/x/net v0.37.0
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.5
	k8s.io/klog/v2 v2.130.1
	sigs.k8s.io/azurelustre-csi-driver v0.1.18
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250303144028-a0af3efb3deb // indirect
)
