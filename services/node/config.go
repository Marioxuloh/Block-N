package node

import (
	"crypto/sha256"
	"fmt"
	"net"
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

// Generate a 160-bit ID from an IP address
func generateIDFromAddress(address string) ([20]byte, error) {
	var id [20]byte

	// Convert the IP address to bytes
	ip := net.ParseIP(address)
	if ip == nil {
		return id, fmt.Errorf("invalid ip address: %s", address)
	}

	// Convert IP to 16-byte format for uniformity (handling both IPv4 and IPv6)
	ip = ip.To16()
	if ip == nil {
		return id, fmt.Errorf("unable to convert IP to 16 bytes: %s", address)
	}

	// Compute the SHA-256 hash of the IP address
	hash := sha256.Sum256(ip)

	// Use only the first 20 bytes (160 bits) of the hash to form the ID
	copy(id[:], hash[:20])

	return id, nil
}
