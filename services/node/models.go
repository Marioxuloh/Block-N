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
	ID      uint64
	Address string
}

type DHT struct {
	Buckets [][]Neighbor
}

// crea una nueva tabla DHT con buckets vacíos
func InitDHT(numBuckets, maxNeighborsPerBucket int) DHT {
	// Inicializar los buckets
	buckets := make([][]Neighbor, numBuckets) // En Kademlia, generalmente se usan 160 buckets, SHA-1 es de 160 bits asi se pueden colocar cualquier ID de nodo(si tiene 1 diferencia ira en el bucket n1 . . .)
	for i := range buckets {
		buckets[i] = make([]Neighbor, 0, maxNeighborsPerBucket)
	}
	return DHT{
		Buckets: buckets,
	}
}
