package main

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

// generateIDFromAddress genera un ID único a partir de la dirección IP del nodo
func generateIDFromAddress(address string) (uint64, error) {
	// Convertir la dirección IP a bytes
	ip := net.ParseIP(address)
	if ip == nil {
		return 0, fmt.Errorf("Invalid IP Address: %s", address)
	}

	// Calcular el hash SHA256 de la dirección IP
	hash := sha256.Sum256(ip)

	// Convertir el hash a uint64
	var id uint64
	for _, b := range hash[:8] { // Utilizar los primeros 8 bytes del hash
		id = id<<8 | uint64(b)
	}

	return id, nil
}
