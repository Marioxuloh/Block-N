package main

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

// crea una nueva tabla DHT con buckets vacíos
func InitDHT(numBuckets, maxNeighborsPerBucket int) DHT {
	// Inicializar los buckets
	buckets := make([]map[uint64][]Neighbor, numBuckets) // En Kademlia, generalmente se usan 160 buckets, SHA-1 es de 160 bits asi se pueden colocar cualquier ID de nodo(si tiene 1 diferencia ira en el bucket n1 . . .)
	for i := range buckets {
		buckets[i] = make(map[uint64][]Neighbor, maxNeighborsPerBucket)
	}
	return DHT{
		Buckets: buckets,
	}
}
