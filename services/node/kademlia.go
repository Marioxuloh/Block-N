package node

//implementar crud sobre la tabla, para que los otros servicios puedan acceder y obtener cosas de la tabla etc etc

// inicializacion de la dht
func InitDHT(numBuckets, maxNeighborsPerBucket int) DHT {
	// En Kademlia, generalmente se usan 160 buckets, SHA-1 es de 160 bits asi se pueden colocar cualquier ID de nodo(si tiene 1 diferencia ira en el bucket n1 . . .)
	buckets := make([]map[uint64][]Neighbor, numBuckets)
	for i := range buckets {
		buckets[i] = make(map[uint64][]Neighbor, maxNeighborsPerBucket)
	}
	return DHT{
		Buckets: buckets,
	}
}

//func xorDistance(id1, id2 uint64) uint64 {
//	return id1 ^ id2
//}
