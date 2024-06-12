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
	for _, bucket := range n.DHT.Buckets {
		for _, value := range bucket {
			conn, err := grpc.NewClient(value+":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			}
			defer conn.Close()

			client := pb.NewBootstraperClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			req := &pb.BootstrapRequest{
				Id:      n.ID[:],
				Address: n.Address,
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
	for _, bucket := range n.DHT.Buckets {
		for _, value := range bucket {
			conn, err := grpc.NewClient(value+":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
