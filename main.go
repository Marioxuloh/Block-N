package main

import (
	"Block-N/services/node"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	node, nodeConfig, err := node.InitNode()
	if err != nil {
		log.Fatal(err)
	}
	println()
	fmt.Println("---------Node initialized on Block-Net---------")
	fmt.Println("ID:", hex.EncodeToString(node.ID[:]))
	fmt.Println("Address:", node.Address)
	fmt.Println("NodePort:", nodeConfig.Port)
	fmt.Println("NumBuckets:", nodeConfig.NumBuckets)
	fmt.Println("MaxNeighborsPerBucket:", nodeConfig.MaxNeighborsPerBucket)
	fmt.Println("-----------------------------------------------")
	println()
}
