package discovery

import (
	pb "Block-N/cmd/gRPC/proto"
	"Block-N/services/node"
)

func Bootstrap(n *node.Node, req *pb.BootstrapRequest) error {
	err := n.Store(node.Neighbor{ID: node.Key(req.Id), Address: req.Address})
	if err != nil {
		return err
	}
	return nil
}
