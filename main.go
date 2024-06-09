package main

import (
	"Block-N/services/node"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	node, err := node.InitNode()
	if err != nil {
		log.Fatal(err)
	}
	println()
	fmt.Println("----------------Node initialized---------------")
	fmt.Println("ID:", hex.EncodeToString(node.ID[:]))
	fmt.Println("Name:", node.Config.Name)
	fmt.Println("Address:", node.Address)
	fmt.Println("NodePort:", node.Config.Port)
	fmt.Println("NumBuckets:", node.Config.NumBuckets)
	fmt.Println("MaxNeighborsPerBucket:", node.Config.MaxNeighborsPerBucket)
	fmt.Println("-----------------------------------------------")
	println()
}
