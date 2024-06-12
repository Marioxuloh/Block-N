package node

import (
	"crypto/sha256"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadNodeConfig(filename string) (NodeConfig, error) {

	var config NodeConfig

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// Generate a 160-bit ID from name of node
func generateIDFromAddress(address string) (Key, error) {
	var id Key

	// Compute the SHA-256 hash of the IP address
	hash := sha256.Sum256([]byte(address))

	// Use only the first 20 bytes (160 bits) of the hash to form the ID
	copy(id[:], hash[:])

	return id, nil
}
