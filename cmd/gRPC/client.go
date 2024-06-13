package gRPC

import (
	pb "Block-N/cmd/gRPC/proto"
	"Block-N/services/node"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Bootstrap(n *node.Node) error {
	// multidifusion a todos los nodos conocidos de bootstrap
	for _, bucket := range n.DHT.Buckets {
		for _, address := range bucket {
			conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			} else {
				defer conn.Close()

				client := pb.NewBootstraperClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				req := &pb.BootstrapRequest{
					Id:      n.ID[:],
					Address: n.Config.Address,
				}

				_, err := client.Bootstrap(ctx, req)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func Discovery(n *node.Node, timeout int) error {
	// multidifusion a todos los nodos conocidos de discovery, a esto se le debe agregar algo para que se repita cada 10 segundos indefinidamente
	aux := 0

outerLoop:
	for _, bucket := range n.DHT.Buckets {
		for _, address := range bucket {
			conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatal(err)
			} else {
				defer conn.Close()

				client := pb.NewDiscovererClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				req := &pb.DiscoveryRequest{
					Id: n.ID[:],
				}

				res, err := client.Discovery(ctx, req)
				if err != nil {
					log.Fatal(err)
				}

				neighbor_list := res.GetNeighborList()

				// storear los nuevos vecinos en la dht
				if len(neighbor_list) != 0 {
					for _, neighbor := range neighbor_list {
						n.Store(node.Neighbor{ID: node.Key(neighbor.Id), Address: neighbor.Address})
					}
				}
				aux++
				if aux >= n.Config.MaxNeighborsPerBucket {
					break outerLoop
				}
			}
		}
	}

	return nil
}

// llamada recursiva hasta encontrar el nodo, si devuelve vacio todo es que unreachable, jumps inicialmente debe ser 0
func Retrieve(n *node.Node, domain string, jumps int32) (*pb.RetrieveResponse, error) {

	id := node.GenerateIDFromAddress(domain)
	neighborAddress, exists := n.Retrieve(id)
	if exists {
		return &pb.RetrieveResponse{Id: id[:], Address: neighborAddress}, nil
	}

	if jumps >= int32(n.Config.MaxNeighborsPerBucket) {
		return &pb.RetrieveResponse{Id: id[:], Address: "peer unreachable"}, nil
	}

	conn, err := grpc.NewClient(n.RetrieveClosestNeighbor(id).Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return &pb.RetrieveResponse{}, err
	}

	defer conn.Close()

	client := pb.NewRetrieverClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.RetrieveRequest{
		Domain: domain,
		Jumps:  jumps,
	}

	res, err := client.Retrieve(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.RetrieveResponse{Id: res.GetId(), Address: res.GetAddress()}, nil

}
