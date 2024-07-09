package gRPC

import (
	"context"
	"log"
	"net"

	pb "Block-N/cmd/gRPC/proto"
	"Block-N/services/discovery"
	"Block-N/services/node"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBootstraperServer
	pb.UnimplementedDiscovererServer
	*node.Node
}

func InitgRPCServer(n *node.Node) error {

	lis, err := net.Listen("tcp", n.Config.Address)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	pb.RegisterBootstraperServer(s, &server{Node: n})
	pb.RegisterDiscovererServer(s, &server{Node: n})

	log.Printf("server gRPC listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *server) Bootstrap(ctx context.Context, req *pb.BootstrapRequest) (*pb.BootstrapResponse, error) {
	// hara store del nodo que ha enviado el mensaje demomento no se verifica si hay nodos duplicados etc etc
	err := discovery.Bootstrap(s.Node, req)
	if err != nil {
		return &pb.BootstrapResponse{Message: "Bootstrap not successful"}, err
	}

	return &pb.BootstrapResponse{Message: "Bootstrap successful"}, nil

}

func (s *server) Discovery(ctx context.Context, req *pb.DiscoveryRequest) (*pb.DiscoveryResponse, error) {
	// devolvera un grupo de vecinos cercanos
	neighbor_list, err := discovery.Discovery(s.Node, req)
	if err != nil {
		return &pb.DiscoveryResponse{}, err
	}
	return neighbor_list, nil
}

// func Retrieve(n *node.Node, domain string, jumps int) (node.Neighbor, error) {
func (s *server) Retrieve(ctx context.Context, req *pb.RetrieveRequest) (*pb.RetrieveResponse, error) {
	// devolvera un grupo de vecinos cercanos
	neighbor, err := Retrieve(s.Node, req.GetDomain(), req.GetJumps()+1)
	if err != nil {
		return &pb.RetrieveResponse{}, err
	}
	return neighbor, nil
}
