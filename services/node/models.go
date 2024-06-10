package node

type NodeConfig struct {
	Domain                string `yaml:"domain"`
	Address               string `yaml:"address"`
	Port                  int    `yaml:"port"`
	NumBuckets            int    `yaml:"num_buckets"`
	MaxNeighborsPerBucket int    `yaml:"max_neighbors_per_bucket"`
}

type Node struct {
	ID      key
	Address string
	DHT     DHT
	Config  NodeConfig
}

type Neighbor struct {
	ID      key
	Address string
}

type DHT struct {
	Buckets []map[key]string
}

type key [32]byte
