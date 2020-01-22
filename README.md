# corr-id-generator
[![Build Status](https://travis-ci.org/freundallein/corr-id-generator.svg?branch=master)](https://travis-ci.org/freundallein/corr-id-generator)
[![Go Report Card](https://goreportcard.com/badge/github.com/freundallein/corr-id-generator)](https://goreportcard.com/report/github.com/freundallein/corr-id-generator)

Correlation ID generator

Creates unique int64 id based on time epoch and machine id.

## Configuration
Application supports configuration via environment variables:
```
export PORT=7891  # gRPC port
export MACHINE_ID=1  # unique service ID - uint8
```

## Installation
### With docker  
```
$> docker pull freundallein/corridgen
```
### With source
```
$> git clone git@github.com:freundallein/corr-id-generator.git
$> cd corr-id-generator
$> make build
```

## Usage
Docker-compose

```
version: "3.5"

networks:
  network:
    name: example-network
    driver: bridge

services:
  corridgen_1:
    image: freundallein/corridgen:latest
    container_name: corridgen_1
    restart: always
    environment: 
      - PORT=7891
      - MACHINE_ID=1
    networks: 
      - network

  corridgen_2:
    image: freundallein/corridgen:latest
    container_name: corridgen_2
    restart: always
    environment: 
      - PORT=7891
      - MACHINE_ID=2
    networks: 
      - network
```

## Protocol
```
service CorrelationIdGenerator {
    rpc GetCorrelationId(GetRequest) returns (Response) {}
}

message GetRequest {
}

message Response {
    uint64 correlationId = 1;
}
```  

Good luck.
