package discovery

import (
	pb "Block-N/cmd/gRPC/proto"
	"Block-N/services/node"
)

// esta funcion se invocara cuando llegue un mensaje discovery a este nodo
func Discovery(n *node.Node, req *pb.DiscoveryRequest) error {
	return nil
}
