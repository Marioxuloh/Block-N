# Proyecto de Red Routing Overlay Descentralizada con gRPC y Kademlia

## Introducción

### Routing Overlay
Un **routing overlay** es una red lógica superpuesta sobre una red física. Los nodos en un routing overlay pueden comunicarse entre sí directamente a través de enlaces virtuales, independientemente de su ubicación física en la red subyacente. Esto permite la creación de redes eficientes y escalables, facilitando la transmisión de datos y la realización de tareas distribuidas.

![Routing Overlay](images/overlay.jpg)

### DHT Kademlia
Kademlia es una **Distributed Hash Table (DHT)** que permite la localización eficiente de nodos y datos en una red P2P. Utiliza un sistema de enrutamiento basado en XOR y una tabla de enrutamiento estructurada en buckets, lo que facilita la escalabilidad y la tolerancia a fallos. Kademlia es ampliamente utilizada en redes descentralizadas debido a su eficiencia y robustez.

![Kademlia](images/kademlia.jpg)

## Arquitectura del Proyecto

El proyecto está organizado en una arquitectura de microservicios, donde cada servicio cumple una función específica y se comunica con los demás a través de gRPC. A continuación se describen los componentes y su estructura:

### Descripción de los Componentes

#### **Servicios**

1. **node-service**: Maneja la lógica de los nodos individuales en la red Kademlia.
   - `main.go`: Punto de entrada del servicio.
   - `node_service.go`: Implementación de la lógica del nodo y Kademlia.
   - `node_service_test.go`: Pruebas unitarias para el servicio.
   - `models.go`: Definición de los modelos de datos.
   - `kademlia.go`: Implementación del protocolo Kademlia.
   - `proto/node.proto`: Definición de los servicios y mensajes gRPC para el nodo.
   - `proto/node.pb.go`: Código generado de Protocol Buffers para gRPC.

2. **network-service**: Maneja la comunicación entre nodos y la lógica de red.
   - `main.go`: Punto de entrada del servicio.
   - `network_service.go`: Implementación de la lógica de red.
   - `network_service_test.go`: Pruebas unitarias para el servicio.
   - `models.go`: Definición de los modelos de datos.
   - `proto/network.proto`: Definición de los servicios y mensajes gRPC para la red.
   - `proto/network.pb.go`: Código generado de Protocol Buffers para gRPC.

3. **discovery-service**: Maneja el descubrimiento de nodos en la red.
   - `main.go`: Punto de entrada del servicio.
   - `discovery_service.go`: Implementación de la lógica de descubrimiento de nodos.
   - `discovery_service_test.go`: Pruebas unitarias para el servicio.
   - `models.go`: Definición de los modelos de datos.
   - `proto/discovery.proto`: Definición de los servicios y mensajes gRPC para descubrimiento.
   - `proto/discovery.pb.go`: Código generado de Protocol Buffers para gRPC.

#### **Paquetes Compartidos**

1. **kademlia**: Implementación del protocolo Kademlia.
   - `kademlia.go`: Implementación principal del protocolo.
   - `routing_table.go`: Lógica de la tabla de enrutamiento.
   - `contact.go`: Definición y manejo de contactos.
   - `bucket.go`: Implementación de los buckets.

2. **utils**: Utilidades compartidas entre servicios.
   - `utils.go`: Funciones utilitarias.
   - `logger.go`: Lógica de logging.

#### **Configuraciones**

- `node-config.yaml`: Configuración del servicio de nodos.
- `network-config.yaml`: Configuración del servicio de red.
- `discovery-config.yaml`: Configuración del servicio de descubrimiento.

#### **Scripts**

- `start.sh`: Script para iniciar todos los servicios.
- `deploy.sh`: Script para desplegar todos los servicios.

#### **Datos**

- `node-service/node.db`: Base de datos para el servicio de nodos.
- `network-service/network.db`: Base de datos para el servicio de red.
- `discovery-service/discovery.db`: Base de datos para el servicio de descubrimiento.

### Comunicación con gRPC y Protocol Buffers

Cada servicio define sus propios métodos y mensajes gRPC utilizando Protocol Buffers. Los archivos `.proto` se encuentran en el subdirectorio `proto` de cada servicio, y el código generado por `protoc` se utiliza para facilitar la comunicación eficiente entre los servicios.

### Ejemplo de Definición de Proto File (node.proto)

```protobuf
syntax = "proto3";

package node;

service NodeService {
  rpc JoinNetwork (JoinRequest) returns (JoinResponse);
  rpc StoreData (StoreRequest) returns (StoreResponse);
  rpc LookupData (LookupRequest) returns (LookupResponse);
}

message JoinRequest {
  string nodeId = 1;
  string address = 2;
}

message JoinResponse {
  bool success = 1;
  string message = 2;
}

message StoreRequest {
  string key = 1;
  bytes value = 2;
}

message StoreResponse {
  bool success = 1;
  string message = 2;
}

message LookupRequest {
  string key = 1;
}

message LookupResponse {
  bool success = 1;
  bytes value = 2;
  string message = 3;
}
