package service

import (
	"context"
	"log"
	"fmt"
	"net"

	
	pb "github.com/freundallein/corr-id-generator/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/freundallein/corr-id-generator/generator"
	"github.com/freundallein/corr-id-generator/settings"
)

type CorrIDService struct {
	config  *settings.Settings
	generator generator.Generator
}

func (service *CorrIDService) GetCorrelationId(ctx context.Context, request *pb.GetRequest) (*pb.Response, error) {
	correlationId := service.generator.GetId()
	log.Printf("[service] responde with 0x%x\n", correlationId)
	return &pb.Response{CorrelationId: correlationId}, nil
}

func New(config *settings.Settings) (*CorrIDService, error) {
	gen := generator.New(config.MachineId)
	return &CorrIDService{config: config, generator: gen}, nil
}

func (serv *CorrIDService) Start() error {
	port := fmt.Sprintf(":%s", serv.config.RpcPort)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCorrelationIdGeneratorServer(grpcServer, serv)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
