package node

type NodeConfig struct {
	Domain                string `yaml:"domain"`
	Address               string `yaml:"address"`
	NumBuckets            int    `yaml:"num_buckets"`
	MaxNeighborsPerBucket int    `yaml:"max_neighbors_per_bucket"`
}

type Node struct {
	ID     Key
	DHT    DHT
	Config NodeConfig
}

type Neighbor struct {
	ID      Key
	Address string
}

type DHT struct {
	Buckets []map[Key]string
}

type Key [32]byte
