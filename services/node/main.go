package main

import (
	"fmt"
	"log"
)

func main() {
	// Cargar la configuración del nodo desde el archivo nodeConfig.yaml
	config, err := LoadNodeConfig("./configs/nodeConfig.yaml")
	if err != nil {
		log.Fatal("Error loading node config:", err)
	}

	id, err := generateIDFromAddress(config.Address)
	if err != nil {
		log.Fatal(err)
	}

	node := Node{
		ID:      id,
		Address: config.Address,
		DHT:     InitDHT(config.NumBuckets, config.NumBuckets),
	}

	// Agregar algunos vecinos ficticios a la tabla DHT
	for i := 0; i < len(node.DHT.Buckets); i++ {
		// Agregar vecinos a cada bucket (por ejemplo, solo uno para este caso)
		node.DHT.Buckets[i] = append(node.DHT.Buckets[i], Neighbor{ID: 2, Address: "192.168.1.1"})
		node.DHT.Buckets[i] = append(node.DHT.Buckets[i], Neighbor{ID: 2, Address: "192.168.1.1"})
	}
	// Iterar sobre los buckets e imprimir cada uno
	for i, bucket := range node.DHT.Buckets {
		fmt.Printf("Bucket %d: %v\n", i, bucket)
	}

	println(node.ID, node.Address, node.DHT.Buckets)

}
