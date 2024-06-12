package discovery

import (
	pb "Block-N/cmd/gRPC/proto"
	"Block-N/services/node"
)

// esta funcion se invocara cuando llegue un mensaje discovery a este nodo
func Discovery(n *node.Node, req *pb.DiscoveryRequest) (*pb.DiscoveryResponse, error) {
	var neighbor_list []*pb.Neighbor

	// etiquetamops el bucle y salimos cuando ya hemos cogido los n.Config.MaxNeighborsPerBucket=20 vecinos mas cercanos
outerLoop:
	for _, bucket := range n.DHT.Buckets {
		for key, address := range bucket {
			neighbor_list = append(neighbor_list, &pb.Neighbor{Id: key[:], Address: address})
			if len(neighbor_list) >= n.Config.MaxNeighborsPerBucket {
				break outerLoop
			}
		}
	}
	return &pb.DiscoveryResponse{NeighborList: neighbor_list}, nil
}
