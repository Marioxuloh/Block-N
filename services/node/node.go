package node

func InitNode() (*Node, *NodeConfig, error) {
	// Cargar la configuración del nodo desde el archivo nodeConfig.yaml
	config, err := LoadNodeConfig("./configs/nodeConfig.yaml")
	if err != nil {
		return &Node{}, &NodeConfig{}, err
	}

	id, err := generateIDFromAddress(config.Address)
	if err != nil {
		return &Node{}, &NodeConfig{}, err
	}

	node := Node{
		ID:      id,
		Address: config.Address,
		DHT:     InitDHT(config.NumBuckets, config.NumBuckets),
	}

	/*
		id2 := [20]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xAA, 0xAF, 0xFF, 0xFF, 0xFF, 0xFF, 0xAF, 0xFF, 0xFE}
		node.Store(Neighbor{ID: id2, Address: "192.168.1.122"})
		for i, bucket := range node.DHT.Buckets {
			fmt.Printf("Bucket %d: %v\n", i, bucket)
		}
		fmt.Println(hammingDistance(id, id2))
	*/
	return &node, &config, nil

}
