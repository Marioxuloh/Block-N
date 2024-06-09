package node

type NodeConfig struct {
	Address               string `yaml:"address"`
	Port                  int    `yaml:"port"`
	NumBuckets            int    `yaml:"num_buckets"`
	MaxNeighborsPerBucket int    `yaml:"max_neighbors_per_bucket"`
}

type Node struct {
	ID      uint64
	Address string
	DHT     DHT
}

type Neighbor struct {
	Address string
}

type DHT struct {
	Buckets []map[uint64][]Neighbor
}
