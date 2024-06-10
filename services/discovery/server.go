package discovery

import (
	"context"
	"log"
	"net"

	pb "Block-N/services/discovery/proto" // Reemplaza con la ruta correcta a tus archivos generados

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBootstraperServer
	pb.UnimplementedDiscovererServer
}

func (s *server) Bootstrap(ctx context.Context, req *pb.BootstrapRequest) (*pb.BootstrapResponse, error) {
	// Implementa la lógica de tu servicio aquí
	log.Printf("Received: %v", req)
	return &pb.BootstrapResponse{Message: "Bootstrap successful"}, nil
}

func (s *server) Discovery(ctx context.Context, req *pb.DiscoveryRequest) (*pb.DiscoveryResponse, error) {
	// Implementa la lógica de tu servicio aquí
	log.Printf("Received: %v", req)
	return &pb.DiscoveryResponse{Address: "Bootstrap successful"}, nil
}

func InitServer(address string) error {

	lis, err := net.Listen("tcp", address+":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	pb.RegisterBootstraperServer(s, &server{})
	pb.RegisterDiscovererServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
