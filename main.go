package main

import (
	"Block-N/services/node"
	"log"
)

func main() {
	node, err := node.InitNode()
	if err != nil {
		log.Fatal(err)
	}
	println(node.ID)
}
