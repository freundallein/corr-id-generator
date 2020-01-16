package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/freundallein/corr-id-generator/service"
	"github.com/freundallein/corr-id-generator/settings"
)

const (
	timeFormat   = "02.01.2006 15:04:05"
	machineIdKey = "MACHINE_ID"
	portKey      = "PORT"
)

type logWriter struct {
}

// Write - custom logger formatting
func (writer logWriter) Write(bytes []byte) (int, error) {
	msg := fmt.Sprintf("%s | [corridgen] %s", time.Now().UTC().Format(timeFormat), string(bytes))
	return fmt.Print(msg)
}

func getEnv(key string, fallback string) (string, error) {
	if value := os.Getenv(key); value != "" {
		return value, nil
	}
	return fallback, nil
}

func getIntEnv(key string, fallback int) (int, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fallback, err
		}
		return int(i), nil
	}
	return fallback, nil
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	port, err := getEnv(portKey, "7891")
	if err != nil {
		log.Fatalf("[config] %s", err.Error())
	}
	machineId, err := getIntEnv(machineIdKey, 1)
	if err != nil {
		log.Fatalf("[config] %s", err.Error())
	}
	config := &settings.Settings{
		RpcPort: port,
		MachineId: uint8(machineId),
	}
	serv, err := service.NewService(config)
	if err != nil {
		log.Fatalf("Can't create gRPC service: %s", err)
	}
	log.Println("Starting service...")
	err = serv.Start()
	if err != nil {
		log.Fatal(fmt.Sprintf("With gRPC service: %s", err))
	}
}
