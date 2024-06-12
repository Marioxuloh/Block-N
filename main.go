package main

import (
	"Block-N/cmd/gRPC"
	"Block-N/services/node"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	n, err := node.InitNode()
	if err != nil {
		log.Fatal(err)
	}
	println()
	fmt.Println("----------------Node initialized---------------")
	fmt.Println("ID:", hex.EncodeToString(n.ID[:]))
	fmt.Println("Domain:", n.Config.Domain)
	fmt.Println("Address:", n.Address)
	fmt.Println("NodePort:", n.Config.Port)
	fmt.Println("NumBuckets:", n.Config.NumBuckets)
	fmt.Println("MaxNeighborsPerBucket:", n.Config.MaxNeighborsPerBucket)
	fmt.Println("-----------------------------------------------")
	println()

	go gRPC.InitServer(n)
	if err != nil {
		log.Fatal(err)
	}

	err = gRPC.Bootstrap(n)
	if err != nil {
		log.Fatal(err)
	}

	err = gRPC.Discovery(n, 10000)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
