package service

import (
	"context"
	"fmt"

	"gitlab.com/freundallein/corr-id-generator/generator"
	pb "gitlab.com/freundallein/corr-id-generator/grpc/proto"
	log "gitlab.com/freundallein/gologger"
)

type CorrIDService struct {
	generator generator.Generator
}

func (service *CorrIDService) GetCorrelationId(ctx context.Context, request *pb.GetRequest) (*pb.Response, error) {
	correlationId := service.generator.GetId()
	log.Debug(fmt.Sprintf("Responde with 0x%x", correlationId))
	return &pb.Response{CorrelationId: correlationId}, nil
}
