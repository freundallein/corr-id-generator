package service

import (
	"fmt"
	"net"

	"gitlab.com/freundallein/corr-id-generator/generator"
	pb "gitlab.com/freundallein/corr-id-generator/grpc/proto"
	"gitlab.com/freundallein/corr-id-generator/settings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Starter struct {
	config  *settings.Settings
	service *CorrIDService
}

func NewService(config *settings.Settings) (*Starter, error) {
	gen := generator.NewGenerator(config.MachineId)
	service := &CorrIDService{generator: gen}
	return &Starter{config: config, service: service}, nil
}

func (st *Starter) Start() error {
	port := fmt.Sprintf(":%s", st.config.RpcPort)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCorrelationIdGeneratorServer(grpcServer, st.service)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
