/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csicommon

import (
	"fmt"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

func ParseEndpoint(ep string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(ep), "unix://") || strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("invalid endpoint: %v", ep)
}

func NewVolumeCapabilityAccessMode(mode csi.VolumeCapability_AccessMode_Mode) *csi.VolumeCapability_AccessMode {
	return &csi.VolumeCapability_AccessMode{Mode: mode}
}

func NewDefaultNodeServer(d *CSIDriver) *DefaultNodeServer {
	return &DefaultNodeServer{
		Driver: d,
	}
}

func NewDefaultIdentityServer(d *CSIDriver) *DefaultIdentityServer {
	return &DefaultIdentityServer{
		Driver: d,
	}
}

func NewDefaultControllerServer(d *CSIDriver) *DefaultControllerServer {
	return &DefaultControllerServer{
		Driver: d,
	}
}

func NewControllerServiceCapability(capability csi.ControllerServiceCapability_RPC_Type) *csi.ControllerServiceCapability {
	return &csi.ControllerServiceCapability{
		Type: &csi.ControllerServiceCapability_Rpc{
			Rpc: &csi.ControllerServiceCapability_RPC{
				Type: capability,
			},
		},
	}
}

func NewNodeServiceCapability(capability csi.NodeServiceCapability_RPC_Type) *csi.NodeServiceCapability {
	return &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{
				Type: capability,
			},
		},
	}
}

func RunNodePublishServer(endpoint string, d *CSIDriver, ns csi.NodeServer, testMode bool) {
	ids := NewDefaultIdentityServer(d)
	s := NewNonBlockingGRPCServer()
	s.Start(endpoint, ids, nil, ns, testMode)
	s.Wait()
}

func RunControllerPublishServer(endpoint string, d *CSIDriver, cs csi.ControllerServer, testMode bool) {
	ids := NewDefaultIdentityServer(d)
	s := NewNonBlockingGRPCServer()
	s.Start(endpoint, ids, cs, nil, testMode)
	s.Wait()
}

func RunControllerandNodePublishServer(endpoint string, d *CSIDriver, cs csi.ControllerServer, ns csi.NodeServer, testMode bool) {
	ids := NewDefaultIdentityServer(d)
	s := NewNonBlockingGRPCServer()
	s.Start(endpoint, ids, cs, ns, testMode)
	s.Wait()
}

func getLogLevel(method string) int32 {
	if method == "/csi.v1.Identity/Probe" ||
		method == "/csi.v1.Node/NodeGetCapabilities" ||
		method == "/csi.v1.Node/NodeGetVolumeStats" {
		return 6
	}
	return 2
}

func logGRPC(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	level := klog.Level(getLogLevel(info.FullMethod))
	klog.V(level).Infof("GRPC call: %s", info.FullMethod)
	klog.V(level).Infof("GRPC request: %s", protosanitizer.StripSecrets(req))

	resp, err := handler(ctx, req)
	if err != nil {
		klog.Errorf("GRPC error: %v", err)
	} else {
		klog.V(level).Infof("GRPC response: %s", protosanitizer.StripSecrets(resp))
	}
	return resp, err
}
