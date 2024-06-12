package node

func InitNode() (*Node, error) {
	// Cargar la configuración del nodo desde el archivo nodeConfig.yaml
	config, err := LoadNodeConfig("./configs/nodeConfig.yaml")
	if err != nil {
		return &Node{}, err
	}

	id, err := generateIDFromAddress(config.Domain)
	if err != nil {
		return &Node{}, err
	}

	node := Node{
		ID:     id,
		DHT:    InitDHT(config.NumBuckets, config.NumBuckets),
		Config: config,
	}

	return &node, nil

}
