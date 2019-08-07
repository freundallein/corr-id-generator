package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/freundallein/corr-id-generator/service"
	"gitlab.com/freundallein/corr-id-generator/settings"

	log "gitlab.com/freundallein/gologger"
	configuration "gitlab.com/freundallein/gonfig"
)

var configType = flag.String("config-type", "default", "Choose default, local or external")

// Init - init logger and config
func Init(serviceName string) *settings.Settings {
	log.InitLogger(serviceName)
	log.Debug("Logger initiated")
	flag.Parse()
	configurator := configuration.New(*configType)
	log.Debug(fmt.Sprintf("%s configurator created", configurator.GetName()))
	config := &settings.Settings{Name: serviceName}
	err := configurator.SetConfigStruct(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Can't set config struct: %s", err))
	}
	err = configurator.ReadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Can't read config: %s", err))
	}
	log.Debug("Read config")
	return config
}

func main() {
	config := Init(os.Getenv("SERVICE_NAME"))
	log.Debug("Create service")
	service, err := service.NewService(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Can't create gRPC service: %s", err))
	}
	log.Debug("Start service")
	err = service.Start()
	if err != nil {
		log.Fatal(fmt.Sprintf("With gRPC service: %s", err))
	}
}
