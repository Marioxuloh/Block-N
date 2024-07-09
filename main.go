package main

import (
	"Block-N/cmd/api"
	"Block-N/cmd/gRPC"
	"Block-N/services/node"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	n, err := node.InitNode()
	if err != nil {
		log.Fatal(err)
	}
	println()
	fmt.Println("----------------Node initialized---------------")
	fmt.Println("ID:", hex.EncodeToString(n.ID[:]))
	fmt.Println("Domain:", n.Config.Domain)
	fmt.Println("Address:", n.Config.Address)
	fmt.Println("NumBuckets:", n.Config.NumBuckets)
	fmt.Println("MaxNeighborsPerBucket:", n.Config.MaxNeighborsPerBucket)
	fmt.Println("-----------------------------------------------")
	println()

	//n.Store(node.Neighbor{ID: node.GenerateIDFromAddress("SnyderNode1"), Address: "192.168.1.144:5000"})

	go gRPC.InitgRPCServer(n)
	if err != nil {
		log.Fatal(err)
	}

	flag := true
	go func() {
		for flag {
			err = gRPC.Bootstrap(n)
			if err != nil {
				log.Fatal(err)
				time.Sleep(time.Second * 2)
			} else {
				flag = false
			}
		}
	}()

	go func() {
		for {
			/*
				fmt.Println("----------------------DHT----------------------")
				for _, bucket := range n.DHT.Buckets {
					for key, address := range bucket {
						println(key[:], address)
					}
				}
				fmt.Println("-----------------------------------------------")
				println()
			*/
			err = gRPC.Discovery(n, 10000)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 10)
		}
	}()

	go api.InitRESTServer(n)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	wg.Wait()
}
