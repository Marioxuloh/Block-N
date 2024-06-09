package node

import (
	"math"
	"math/bits"
)

// inicializacion de la dht
func InitDHT(numBuckets, maxNeighborsPerBucket int) DHT {
	// En Kademlia, generalmente se usan 160 buckets, SHA-1 es de 160 bits asi se pueden colocar cualquier ID de nodo(si tiene 1 diferencia ira en el bucket n1 . . .)
	// en este caso yo hago un hash con sha-256 y solo cojo los primeros 20bytes(160bits)
	buckets := make([]map[[20]byte]string, numBuckets)
	for i := range buckets {
		buckets[i] = make(map[[20]byte]string, maxNeighborsPerBucket)
	}
	return DHT{
		Buckets: buckets,
	}
}

func (n *Node) Store(neighbor Neighbor) {
	i := hammingDistance(n.ID, neighbor.ID)
	n.DHT.Buckets[i][neighbor.ID] = neighbor.Address
}

func (n *Node) Delete(key [20]byte) {
	for _, bucket := range n.DHT.Buckets {
		if _, exists := bucket[key]; exists {
			delete(bucket, key)
			break
		}
	}
}

func (n *Node) Retrieve(key [20]byte) (string, bool) {
	i := hammingDistance(n.ID, key)
	value, exists := n.DHT.Buckets[i][key]
	return value, exists
}

func (n *Node) RetrieveClosestNeighbor(targetID [20]byte) Neighbor {
	closestNeighbor := Neighbor{}
	closestDistance := math.MaxInt
	for _, bucket := range n.DHT.Buckets {
		for key := range bucket {
			distance := hammingDistance(targetID, key)
			if distance < closestDistance {
				closestDistance = distance
				closestNeighbor = Neighbor{ID: key, Address: bucket[key]}
			}
		}
	}
	return closestNeighbor
}

// Returns an integer between 0 and 159 indicating the distance
func hammingDistance(id1, id2 [20]byte) int {
	distance := -1
	for i := 0; i < 20; i++ {
		distance += bits.OnesCount8(id1[i] ^ id2[i])
	}
	return distance
}
