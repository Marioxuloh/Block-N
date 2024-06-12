package gRPC

import (
	pb "Block-N/cmd/gRPC/proto"
	"Block-N/services/node"
	"context"
	"fmt"
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
			}
			defer conn.Close()

			client := pb.NewBootstraperClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			req := &pb.BootstrapRequest{
				Id:      n.ID[:],
				Address: n.Config.Address,
			}

			res, err := client.Bootstrap(ctx, req)
			if err != nil {
				return err
			}

			fmt.Println(res.GetMessage())
		}
	}
	return nil
}

func Discovery(n *node.Node, timeout int) error {
	// multidifusion a todos los nodos conocidos de discovery, a esto se le debe agregar algo para que se repita cada 10 segundos indefinidamente
	for _, bucket := range n.DHT.Buckets {
		for _, address := range bucket {
			conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			}
			defer conn.Close()

			client := pb.NewDiscovererClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			req := &pb.DiscoveryRequest{
				Id: n.ID[:],
			}

			res, err := client.Discovery(ctx, req)
			if err != nil {
				return err
			}

			fmt.Println(res.GetAddress())
		}
	}

	return nil
}
