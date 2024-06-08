#!/bin/bash

# Función para ejecutar el servicio Node
run_node_service() {
    echo "Ejecutando el servicio Node..."
    go run services/node/main.go services/node/config.go services/node/models.go
}

# Función para ejecutar el servicio Network
run_network_service() {
    echo "Ejecutando el servicio Network..."
    go run services/network/main.go
}

# Función para ejecutar el servicio Discovery
run_discovery_service() {
    echo "Ejecutando el servicio Discovery..."
    go run services/discovery/main.go
}

# Verificar si se proporcionaron tres argumentos
if [ $# -ne 3 ]; then
    echo "Uso: $0 <node> <network> <discovery>"
    echo "Donde cada argumento es 'true' para ejecutar el servicio o cualquier otro valor para omitirlo."
    exit 1
fi

# Verificar los argumentos de entrada y ejecutar los servicios correspondientes
if [ "$1" = "true" ]; then
    run_node_service
fi

if [ "$2" = "true" ]; then
    run_network_service
fi

if [ "$3" = "true" ]; then
    run_discovery_service
fi
