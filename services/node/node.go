package node

import (
	"fmt"
)

func InitNode() (Node, error) {
	// Cargar la configuración del nodo desde el archivo nodeConfig.yaml
	config, err := LoadNodeConfig("./configs/nodeConfig.yaml")
	if err != nil {
		return Node{}, err
	}

	id, err := generateIDFromAddress(config.Address)
	if err != nil {
		return Node{}, err
	}

	node := Node{
		ID:      id,
		Address: config.Address,
		DHT:     InitDHT(config.NumBuckets, config.NumBuckets),
	}

	// Agregar algunos vecinos ficticios a la tabla DHT
	for i := 0; i < len(node.DHT.Buckets); i++ {
		// Agregar vecinos a cada bucket (por ejemplo, solo uno para este caso)
		node.DHT.Buckets[i][2] = append(node.DHT.Buckets[i][2], Neighbor{Address: "192.168.1.102"})
		node.DHT.Buckets[i][3] = append(node.DHT.Buckets[i][3], Neighbor{Address: "192.168.1.103"})
	}
	// Iterar sobre los buckets e imprimir cada uno
	for i, bucket := range node.DHT.Buckets {
		fmt.Printf("Bucket %d: %v\n", i, bucket)
	}

	println(node.ID, node.Address, node.DHT.Buckets)

	return node, nil

}
