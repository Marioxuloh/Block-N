package node

type NodeConfig struct {
	Name                  string `yaml:"name"`
	Address               string `yaml:"address"`
	Port                  int    `yaml:"port"`
	NumBuckets            int    `yaml:"num_buckets"`
	MaxNeighborsPerBucket int    `yaml:"max_neighbors_per_bucket"`
}

type Node struct {
	ID      key
	Address string
	DHT     DHT
}

type Neighbor struct {
	ID      key
	Address string
}

type DHT struct {
	Buckets []map[key]string
}

type key [32]byte
