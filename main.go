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
	fmt.Println("------Nodo inicializado en la red Block-N------")
	fmt.Println("ID:", hex.EncodeToString(node.ID[:]))
	fmt.Println("Address:", node.Address)
	fmt.Println("-----------------------------------------------")
	println()
}
